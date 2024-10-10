package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Data struct {
    A   int `json:"a"`    
    B   int `json:"b"` 
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

    fmt.Printf("Чило go routine: %d\n", numGoroutine)
}
