package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Доступно ядер:", runtime.NumCPU())
	fmt.Println("Выберите сколько хотите использовать")
	numCPU := bufio.NewReader(os.Stdin)
	strNum, _ := numCPU.ReadString('\n')
	strNum = strings.TrimRight(strNum, "\n")
	Num, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println("Неправильно введено значение")
		main()
	}
	if Num >= 1 && Num <= 8 {
		fmt.Println("Выберите способ ввода хэш значения:\n1. Ввод хэша с клавиатуры.\n2. Чтение хэша из файла\n3. Занести хэш таблицу.")
		rd := bufio.NewReader(os.Stdin)
		timeStr, _ := rd.ReadString('\n')
		timeStr = strings.TrimRight(timeStr, "\n")

		switch timeStr {
		case "1":
			User_Input(Num)
		case "2":
			FileInput(Num)
		case "3":
			HashAdd()
		default:
			fmt.Println("Неправильно введено значение")
		}
	} else {
		fmt.Println("Вы ввели неправильное значение")
		main()
	}
}
