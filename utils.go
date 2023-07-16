package main

import (
	"fmt"
	"net/http"
	"os/user"
	"path/filepath"
	"runtime"
	"time"
)


type Website struct {
	status bool
	lastVisited string
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

func generateWebsiteArray(length int) []string {
	if length == 0 {
		return []string{
			"https://random-status-code.herokuapp.com",
			"https://www.google.com",
			"https://www.wikipedia.org",
			"https://www.youtube.com",
		}
	}
	websites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.google.com",
		"https://www.wikipedia.org",
		"https://www.youtube.com",
		"https://www.github.com",
		"https://www.amazon.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
		"https://www.nytimes.com",
		"https://www.stackoverflow.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
		"https://www.netflix.com",
		"https://www.instagram.com",
		"https://www.spotify.com",
		"https://www.dropbox.com",
		"https://www.ubuntu.com",
		"https://www.oracle.com",
		"https://www.ibm.com",
		"https://www.nike.com",
		"https://www.adidas.com",
		"https://www.puma.com",
		"https://www.bbc.co.uk",
		"https://www.cnn.com",
		"https://www.yahoo.com",
		"https://www.twitch.tv",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.snapchat.com",
		"https://www.pinterest.com",
	}

	return websites
}
