/*
The following code contains unittests for hydra continous integration.

Author: Shravan Asati
Originally Written: 21 April 2021
Last Edited: 21 April 2021
*/

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

func generateRandom(value string) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())

	if value == "name" {
		// * generating random values for configurations
	
		var rname string
		// * generating random name
		for len(rname) <= 8 {
			rname += string(letters[rand.Intn(len(letters))])
		}
		return rname
	
	} else if value == "githubUsername" {
		// * generating random github usename 
		var rgithub string
		for len(rgithub) <= 10 {
			rgithub += string(letters[rand.Intn(len(letters))])
		}
		return rgithub
	
	} else if value == "lang" {
		// * random choice of languages
		rlang := supportedLangs[rand.Intn(len(supportedLangs))]
		return rlang
	
	} else if value == "license" {
		// * random supported license
		licenses := []string{}
		for k := range supportedLicenses {
			licenses = append(licenses, k)
		}
		rlicense := licenses[rand.Intn(len(licenses))]
		return rlicense
	
	} else {
		panic(fmt.Sprintf("Invalid value for random generation: %v.", value))
	}
}

func TestConfig(t *testing.T) {
	// * the below line is to ensure the `config.json` file exists
	config("default", "default", "default", "default")

	// * getting all initial values so that after the tests, the cleanup function can restore the configuration
	initialName := getConfig("fullName")
	initialGithubUsername := getConfig("githubUsername")
	initialLang := getConfig("defaultLang")
	initialLicense := getConfig("defaultLicense")

	// * storing random values
	rname := generateRandom("name")
	rgithub := generateRandom("githubUsername")
	rlang := generateRandom("lang")
	rlicense := generateRandom("license")


	// * setting configuration
	config(rname, rgithub, rlang, rlicense)

	// * checking fullname
	gotName := getConfig("fullName")
	if gotName != rname {
		t.Errorf("Got %v, expected %v", gotName, rname)
	}

	// * checking github username
	gotGithub := getConfig("githubUsername")
	if gotGithub != rgithub {
		t.Errorf("Got %v, expected %v", gotGithub, rgithub)
	}

	// * checking default language
	gotLang := getConfig("defaultLang")
	if gotLang != rlang {
		t.Errorf("Got %v, expected %v", gotLang, rlang)
	}

	// * checking default license
	gotLicense := getConfig("defaultLicense")
	if gotLicense != rlicense {
		t.Errorf("Got %v, expected %v", gotLicense, rlicense)
	}

	// * cleaning up
	t.Cleanup(func(){
		t.Log("\nCleaning up...\n")
		config(initialName, initialGithubUsername, initialLang, initialLicense)
	})
}

func TestPythonInit(t *testing.T) {
	gwd, e := os.Getwd()
	handleException(e)

	// * generating random project name and license for initialisation
	rlicense := generateRandom("license")
	rprojectName := generateRandom("name")

	init := Initialiser{
		projectName: rprojectName,
		license: rlicense,
		lang: "python",
	}
	init.pythonInit()

	// * getting all files present in the directory
	files, e := ioutil.ReadDir("./")
	handleException(e)

	// * converting into filenames
	filenames := []string{}
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	// * checking for presence
	if !stringInSlice(rprojectName, filenames) {
		t.Errorf("project %v not in the directory", rprojectName)
	}


	// * getting contents of the project initialised
	projectFiles, er := ioutil.ReadDir("./")
	handleException(er)
	projectFileNames := []string{}
	for _, f := range projectFiles {
		projectFileNames = append(projectFileNames, f.Name())
	}

	// * checking for various files
	if !stringInSlice("LICENSE", projectFileNames) {t.Errorf("LICENSE file not present.")}
	if !stringInSlice("README.md", projectFileNames) {t.Errorf("README.md file not present.")}
	if !stringInSlice(".gitignore", projectFileNames) {t.Errorf(".gitignore file not present.")}
	if !stringInSlice("setup.py", projectFileNames) {t.Errorf("setup.py file not present.")}
	if !stringInSlice(rprojectName, projectFileNames) {t.Errorf("%v dir not present.", rprojectName)}
	if !stringInSlice("tests", projectFileNames) {t.Errorf("tests dir not present.")}

	t.Cleanup(func() {
		t.Log("Cleaning up...")
		os.Chdir(gwd)
		os.RemoveAll(rprojectName)
	})
}

func TestGoInit(t *testing.T) {
	gwd, e := os.Getwd()
	handleException(e)

	// * generating random project name and license for initialisation
	rlicense := generateRandom("license")
	rprojectName := generateRandom("name")
	init := Initialiser{
		projectName: rprojectName,
		license: rlicense,
		lang: "python",
	}
	init.goInit()

	// * getting contents of the project initialised
	projectFiles, er := ioutil.ReadDir("./")
	handleException(er)
	projectFileNames := []string{}
	for _, f := range projectFiles {
		projectFileNames = append(projectFileNames, f.Name())
	}

	// * checking for various files
	if !stringInSlice("LICENSE", projectFileNames) {t.Errorf("LICENSE file not present.")}
	if !stringInSlice("README.md", projectFileNames) {t.Errorf("README.md file not present.")}
	if !stringInSlice(".gitignore", projectFileNames) {t.Errorf(".gitignore file not present.")}
	if !stringInSlice("src", projectFileNames) {t.Errorf("src dir not present.")}
	if !stringInSlice("bin", projectFileNames) {t.Errorf("bin dir not present.")}
	if !stringInSlice("pkg", projectFileNames) {t.Errorf("pkg dir not present.")}
	if !stringInSlice("tests", projectFileNames) {t.Errorf("tests dir not present.")}

	t.Cleanup(func() {
		t.Log("Cleaning up...")
		os.Chdir(gwd)
		os.RemoveAll(rprojectName)
	})
}