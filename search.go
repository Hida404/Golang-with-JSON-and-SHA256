package main

import (
	"fmt"
	"sync"
)

func search(hash string, start int, end int, hashs Hashs, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i<end; i++{
		if hashs[i].Hash == hash {
			fmt.Println(hashs[i].Hash, "=", hashs[i].Value )
			break
		}
	}
}