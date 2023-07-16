package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func monitoring(websites map[string]Website) {
	fmt.Println("Monitoring...")
	
	pool := generateWebsiteArray(0)

	for _, url := range pool {
		isUp := checkIfSiteIsUp(url)
		websites[url] = Website{checkIfSiteIsUp(url), getCurrentTime()}
		if(isUp) {
			fmt.Println(url, "está", color.New(color.FgHiGreen).Sprint("online"))
		} else {
			fmt.Println(url, "está", color.New(color.FgHiRed).Sprint("offline"))
		}
	}
}

func logs(websites map[string]Website) {
	if(len(websites) == 0) {
		fmt.Println(color.New(color.FgHiRed).Sprint("Nenhum site foi monitorado ainda"))
	}
	for key, value := range websites {
		if(value.status) {
			fmt.Println(key, "estava online em:", value.lastVisited, "")
		} else {
			fmt.Println(key, "estava offline em:", value.lastVisited, "")
		}
	}
}

func exit() {
	fmt.Println("Saindo...")
	os.Exit(0)
}