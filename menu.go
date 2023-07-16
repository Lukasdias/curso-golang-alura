package main

import (
	"fmt"

	"github.com/fatih/color"
)

func printMenu() {
	fmt.Println("")
	color.New(color.FgYellow, color.Bold, color.Italic).Println("*** Menu ***")
	
	currentUser := getOsUser()
	currentPath := getCurrentPath()

	fmt.Println("Bem vindo", color.New(color.FgHiMagenta).Sprint(currentUser))
	fmt.Println("Você está no diretório", color.New(color.FgHiMagenta).Sprint(currentPath))

	fmt.Println("1.", color.New(color.FgHiGreen).Sprint("Monitorar sites"))
	fmt.Println("2.", color.New(color.FgHiGreen).Sprint("Ver logs de sites monitorados"))
	fmt.Println("0.", color.New(color.FgHiGreen).Sprint("Sair"))
	
}