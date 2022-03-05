package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
)

type HashValue struct {
	Hash string
	Value string
}

type Hashs []HashValue

const settingsFilename = "test.json"

func HashAdd() {
	var hash Hashs

	str := []byte("aaaaa")
	for i := byte('a'); i<='z'; i++{
		str[0] = i
		for j := byte('a');j <= 'z'; j++{
			str[1] = j
			for k := byte('a');k<='z'; k++{
				str[2] = k
				for x := byte('a');x<='z'; x++{
					str[3] = x
					for y := byte('a');y<='z'; y++{
						str[4] = y
						Value := string(str)
						H := sha256.Sum256(str)
						Hash := string(H[:])
						test := fmt.Sprintf("%x", Hash)
						m := HashValue{test, Value}
						hash = append(hash, m)
					}
				}
			}
		}
	}



	file, _ := os.ReadFile(settingsFilename)

	if(len(file)!=0){
		err := json.Unmarshal(file, &hash)
		if err != nil{
			fmt.Println("Hi")
		}
	}

	value, _ := json.MarshalIndent(hash, "", " ")


	err := os.WriteFile(settingsFilename, value, 0)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println("Готово")
}
