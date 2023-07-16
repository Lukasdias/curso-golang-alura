package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

const MONITORING_COUNT = 1
const MONITORING_SLEEP = 1

func monitoring(websites map[string]Website) {
	pool := getAllWebsitesUrls()

	for i := 0; i < MONITORING_COUNT; i++ {
		fmt.Println("Monitorando pela", color.New(color.BlinkRapid, color.FgHiRed).Sprint(i+1), "vez")
		for _, url := range pool {
			statusCode := float64(getStatusCode(url))
			websites[url] = Website{statusCode, getCurrentTime()}
			fmt.Println(url, "está com status code", getStatusCodeColorString(int(statusCode)))
		}
		fmt.Println("")
		time.Sleep(MONITORING_SLEEP * time.Second)
	}

	saveLocallyTests(websites)
}

func logs(websites map[string]Website) {
	if(len(websites) == 0) {
		fmt.Println(color.New(color.FgHiRed).Sprint("Nenhum site foi monitorado ainda"))
	}
	for key, value := range websites {
		fmt.Println(key, "está com status code", getStatusCodeColorString(int(value.status)))
	}
}

func exit() {
	fmt.Println("Saindo...")
	os.Exit(0)
}