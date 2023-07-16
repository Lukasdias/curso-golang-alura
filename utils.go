package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fatih/color"
)


type Website struct {
	status float64 `json:"status"`
	lastVisited string 	`json:"lastVisited"`
}

func getCommandValue() int {
	var command int
	fmt.Scan(&command)
	return command
}

func getOsUser() string {
	currentUser, err :=  user.Current()
	if err != nil {
		fmt.Println("Erro ao recuperar usuário logado")
		return ""
	}
	return currentUser.Username
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(0)
    sourceDir := filepath.Dir(filename)
    return  sourceDir
}

func getCurrentTime() string {
	currentTime := time.Now()
	formatTime := currentTime.Format("02/01/2006 15:04:05")
	return formatTime
}

func checkIfSiteIsUp(url string) bool {
	response, _ := http.Get(url)
	return response.StatusCode == 200
}

func getStatusCode(url string) int {
	response, _ := http.Get(url)
	return response.StatusCode
}

func getStatusCodeColorString(statusCode int) string {
	switch statusCode {
	case 200:
		return color.New(color.FgHiGreen).Sprint(statusCode)
	case 404:
		return color.New(color.FgHiRed).Sprint(statusCode)
	default:
		return color.New(color.FgHiYellow).Sprint(statusCode)
	}
}

func getAllWebsitesUrls() [] string{
	var urls []string
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls
}

func getLocallySavedTests() map[string]Website{
	filedData, err := ioutil.ReadFile("dictionary.json")
	if err != nil {
		fmt.Println("Não foi possível recuperar os dados salvos localmente")
		return make(map[string]Website)
	}else {
		color.New(color.FgHiGreen).Println("Dados recuperados com sucesso")
		var storedDictionary map[string]map[string]interface{}
		err = json.Unmarshal(filedData, &storedDictionary)
		if err != nil {
			log.Fatal(err)
		}
		return convertToWebsiteDict(storedDictionary)
	}
}

func convertToWebsiteDict(storedDictionary map[string]map[string]interface{}) map[string]Website {
	websiteDict := make(map[string]Website)
	for key, value := range storedDictionary {
		websiteDict[key] = Website{
			status: value["status"].(float64),
			lastVisited: value["lastVisited"].(string),
		}
	}
	return websiteDict
}

func saveLocallyTests(websites map[string]Website) bool {
	// golang não permite que structs sejam serializadas diretamente
	// logo é preciso converter para um formato que possa ser serializado
	serializableDict := make(map[string]map[string]interface{})
	for key, value := range websites {
		serializableDict[key] = map[string]interface{}{
			"status":      value.status,
			"lastVisited": value.lastVisited,
		}
	}

	jsonData, err := json.Marshal(serializableDict)
	if err != nil {
		return false
	}

	err = ioutil.WriteFile("dictionary.json", jsonData, 0666)
	if err != nil {
		return false
	}

	return true
}
