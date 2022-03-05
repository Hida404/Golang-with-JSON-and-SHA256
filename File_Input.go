package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func FileInput(n int) {
	var str []string
	file, err := os.Open("Test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var hash string
		if len(scanner.Text()) != 64 {
			continue
		} else {
			run := []rune(scanner.Text())
			for i := range run {
				if (string(run[i]) >= "0" && string(run[i]) <= "9") || (string(run[i])) >= "a" && string(run[i]) <= "z" {
					hash += string(run[i])
				} else {
					fmt.Println("Найден неправильное хэш значение")
					hash = ""
					break

				}
			}
		}
		str = append(str, hash)
	}
	var arrHash []string
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
	for i := range str {
		if str[i] != "" {
			arrHash = append(arrHash, str[i])
		}
	}
	var wg sync.WaitGroup

	var hashs Hashs
	startTime := time.Now()
	const settingsFilename = "test.json"
	fileJson, _ := os.ReadFile(settingsFilename)
	err = json.Unmarshal(fileJson, &hashs)
	if err != nil {
		fmt.Println("Error")
	}
	val := len(hashs) / n
	end := val

	for i := range arrHash {
		start := 0
		for j := 0; j < n; j++ {
			wg.Add(1)
			go search(arrHash[i], start, end, hashs, &wg)
			start += val
			end += val
		}
		end = val
		wg.Wait()
	}
	elapsed := time.Since(startTime)
	log.Printf("Elapsed time %s", elapsed)

}
