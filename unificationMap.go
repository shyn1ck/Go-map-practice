package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map2 := make(map[string]string)

	fmt.Println("Введите данные для map1 (ключ:значение), введите 'exit' для завершения:")
	for {
		var key, value string
		fmt.Scanln(&key, &value)
		if key == "exit" {
			break
		}
		map1[key] = value
	}

	fmt.Println("Введите данные для map2 (ключ:значение), введите 'exit' для завершения:")
	for {
		var key, value string
		fmt.Scanln(&key, &value)
		if key == "exit" {
			break
		}
		map2[key] = value
	}

	unificationMap := make(map[string]string)
	for key, value := range map1 {
		unificationMap[key] = value
	}
	for key, value := range map2 {
		unificationMap[key] = value
	}

	fmt.Println("Объединенный map:", unificationMap)
}
