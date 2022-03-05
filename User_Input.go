package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)


func User_Input(numCPU int){
	fmt.Println("Введите хэш значение")
	rd := bufio.NewReader(os.Stdin)
	timeStr, _ := rd.ReadString('\n')
	timeStr = strings.TrimRight(timeStr, "\n")
	timeStr = strings.TrimSpace(timeStr)
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(timeStr, " ")
	ar:= strings.Split(replace, " ")
	count := 0
	for i:= range ar{
		if(len(ar[i]) != 64){
			fmt.Println("Неправильное значение хэша")
			main()
		} else {
			ru := []rune(ar[i])
			for i := range ru{
				if((string(ru[i]) >= "0" && string(ru[i])<="9") || (string(ru[i]))>= "a" && string(ru[i]) <= "z"){
					count++
				} else{
					fmt.Println("Неправильно введено значение. Повторите попытку")
					main()
				}
			}
		}
	}
	var wg sync.WaitGroup
	if count % 64 == 0 {
		var hashs Hashs
		startTime := time.Now()
		const settingsFilename = "test.json"
		file, _ := os.ReadFile(settingsFilename)
		err := json.Unmarshal(file, &hashs)
		if err != nil{
			fmt.Println("Error")
		}
		val := len(hashs)/numCPU
		end := val

		for i:= range ar{
			start := 0
			for j := 0; j < numCPU; j++{
				wg.Add(1)
				go search(ar[i], start, end, hashs, &wg)
				start+=val
				end += val
			}
			end = val
			wg.Wait()
		}
		elapsed := time.Since(startTime)
		log.Printf("Elapsed time %s", elapsed)
	}
}



