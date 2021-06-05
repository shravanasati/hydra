package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)


func update() {
	fmt.Println("Updating hydra...")

	fmt.Println("Downloading the hydra executable...")
	// * determining the os-specific url
	url := ""
	switch runtime.GOOS {
	case "windows":
		url = "https://github.com/Shravan-1908/hydra/releases/latest/download/hydra-windows-amd64.exe"
	case "linux":
		url = "https://github.com/Shravan-1908/hydra/releases/latest/download/hydra-linux-amd64"
	case "darwin":
		url = "https://github.com/Shravan-1908/hydra/releases/latest/download/hydra-darwin-amd64"
	default:
		fmt.Println("Your OS isnt supported by hydra.")
		return
	}

	// * sending a request
	res, err := http.Get(url)
	
	if err != nil {
		fmt.Println("Error: Unable to download the executable. Check your internet connection.")
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	// * determining the executable path
	downloadPath, e := os.UserHomeDir()
	if e != nil {
		fmt.Println("Error: Unable to retrieve hydra path.")
		fmt.Println(e)
		return
	}
	downloadPath += "/.hydra/hydra"
	if runtime.GOOS == "windows" {downloadPath += ".exe"}

	os.Rename(downloadPath, downloadPath + "-old")

	exe, er := os.Create(downloadPath)
	if er != nil {
		fmt.Println("Error: Unable to access file permissions.")
		fmt.Println(er)
		return
	}
	defer exe.Close()

	// * writing the recieved content to the hydra executable
	_, errr := io.Copy(exe, res.Body)
	if errr != nil {
		fmt.Println("Error: Unable to write the executable.")
		fmt.Println(errr)
		return
	}

	// * performing an additional `chmod` utility for linux and mac 
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		execute("chmod", "u+x", downloadPath)
	}

	fmt.Println("Update completed!")
}

func deletePreviousInstallation() {
	hydraDir, _ := os.UserHomeDir()
	
	hydraDir += "/.hydra"

	files, _ := ioutil.ReadDir(hydraDir)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), "-old") {
			// fmt.Println("found existsing installation")
			os.Remove(hydraDir + "/" + f.Name())
		}
		// fmt.Println(f.Name())
	}
}
