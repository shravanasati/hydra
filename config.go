/*
The following code is responsible for the config command.

Author: Shravan Asati
Originally Written: 30 March 2021
Last edited: 13 April 2021
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)


type Configuration struct {
	FullName string `json:"FullName"`
	GithubUsername string `json:"GithubUsername"`
	DefaultLang string `json:"DefaultLang"`
	DefaultLicense string `json:"DefaultLicense"`
}

func jsonify(config *Configuration) (string) {
	byteArray, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	return string(byteArray)
}

func readJson(jsonString string) *Configuration {
	var result *Configuration
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		handleException(err)
	}
	return result
}

func getConfig(value string) string {
	// to make sure the config file exists
	config("default", "default", "default", "default")

	usr, _ := user.Current()
	configFile := (filepath.Join(usr.HomeDir, ".hydra/config.json"))
	file, ferr := os.Open(configFile)
	handleException(ferr)
	wholeText := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wholeText = wholeText + line
	}
	file.Close()

	config := readJson(wholeText)
	switch value {
	case "fullName":
		return config.FullName
	case "githubUsername":
		return config.GithubUsername
	case "defaultLang":
		return config.DefaultLang
	case "defaultLicense":
		return config.DefaultLicense
	default:
		return fmt.Sprintf("Undefined value: %v.", value)
	}
} 

func checkForCorrectConfig() bool {
	// to make sure the config file exists
	config("default", "default", "default", "default")

	usr, _ := user.Current()
	configFile := (filepath.Join(usr.HomeDir, ".hydra/config.json"))
	file, ferr := os.Open(configFile)
	handleException(ferr)
	wholeText := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wholeText = wholeText + line
	}
	file.Close()

	config := readJson(wholeText)
	
	if config.FullName == "" || config.GithubUsername == "" {
		return false
	} else {
		return true
	}
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func config(fullName, githubUsername, defaultLang, defaultLicense string) {
	// * defining path of hydra config file
	usr, _ := user.Current()
	hydraDir := filepath.Join(usr.HomeDir, ".hydra")
	
	if pathOk, _ := exists(hydraDir); !pathOk { os.Mkdir(hydraDir, os.ModePerm) }

	configFile := filepath.Join(hydraDir,"config.json")
	
	// * creating a file in case it doesnt exists
	if configOk, _ := exists(configFile); !configOk {
		f, err := os.Create(configFile)
		handleException(err)
		defaultConfig := Configuration{FullName: "", GithubUsername: "", DefaultLang: "", DefaultLicense: "MIT"}
		_, er := f.WriteString(jsonify(&defaultConfig))
		handleException(er)
		f.Close()
	}

	// * reading data from the file
	file, ferr := os.Open(configFile)
	handleException(ferr)
	wholeText := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wholeText = wholeText + line
	}
	file.Close()
	

	// * writing new config to the file by first deleting it
	configStruct := readJson(wholeText)
	if fullName != "default" {
		configStruct.FullName = fullName
		fmt.Printf("Successfully configured the full name to '%v'. \n", fullName)
	}

	if githubUsername != "default" {
		configStruct.GithubUsername = githubUsername
		fmt.Printf("Successfully configured the GitHub username to '%v'. \n", githubUsername)
	}

	if defaultLang != "default" {
		configStruct.DefaultLang = defaultLang
		fmt.Printf("Successfully configured the default language to '%v'. \n", defaultLang)
	}

	if defaultLicense != "default" {
		configStruct.DefaultLicense = defaultLicense
		fmt.Printf("Successfully configured the default license to '%v'. \n", defaultLicense)
	}

	os.Remove(configFile)
	f, err := os.Create(configFile)
	handleException(err)
	_, er := f.WriteString(jsonify(configStruct))
	handleException(er)
	f.Close()
}