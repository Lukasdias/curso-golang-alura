package main

import (
	"fmt"
)


func main() {
	var command int 

	websiteDict := make(map[string]Website)
	for  {
		printMenu()

		fmt.Println("Escolha uma opção: ")
		command = getCommandValue()
		fmt.Println("O comando escolhido foi", command)
		
		switch command {
			case 1: monitoring(websiteDict)
			case 2: logs(websiteDict)
			case 0: exit()
			default: fmt.Println("Comando não reconhecido")
		}
	}
}



