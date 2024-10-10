package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Data struct {
    A   int `json:"a"`    
    B   int `json:"b"` 
}

func calculateSum(data []Data, start, end int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	sum := 0
	for i := start; i < end; i++ {
		sum += data[i].A + data[i].B
	}
	resultChan <- sum
}

func main() {
    if len(os.Args) < 3 {
        log.Fatalf("Формат применения: %s <путь_к_json> <количество_goroutine>\n", os.Args[0])
    }

    numGoroutine, err := strconv.Atoi(os.Args[2])
    if err != nil || numGoroutine <= 0 {
        log.Fatalf("Ошибка при получении чисел goroutine: %v\n", err)
    }

    jsonPath := os.Args[1] 
    jsonData, err := os.ReadFile(jsonPath)
    if err != nil {
        log.Fatalf("Ошибка при чтении json файла: %v\n", err)
    }

    var data []Data
    if err := json.Unmarshal(jsonData, &data); err != nil {
        log.Fatalf("Ошибка при парсинге json файла: %v\n", err)
    }

    dataLen := len(data)
    resultChan := make(chan int, numGoroutine)

    var wg sync.WaitGroup

    chunkSize := (dataLen + numGoroutine - 1) / numGoroutine

    for i := 0; i < numGoroutine; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > dataLen {
			end = dataLen
		}
		wg.Add(1)
		go calculateSum(data, start, end, &wg, resultChan)
	}

    go func() {
		wg.Wait()
		close(resultChan)
	}()

    totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

    fmt.Printf("Total sum: %d\n", totalSum)
}
