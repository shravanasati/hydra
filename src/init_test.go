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
	"os"
	"testing"
)


func getFiles(dir string) []string {
	projectFiles, er := ioutil.ReadDir(dir)
	handleException(er)
	projectFileNames := []string{}
	for _, f := range projectFiles {
		projectFileNames = append(projectFileNames, f.Name())
	}
	return projectFileNames
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
		lang: "go",
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
	if !stringInSlice("tests", projectFileNames) {t.Errorf("tests dir not present.")}

	t.Cleanup(func() {
		t.Log("Cleaning up...")
		os.Chdir(gwd)
		os.RemoveAll(rprojectName)
	})
}

func TestWebInit(t *testing.T) {
	gwd, e := os.Getwd()
	handleException(e)

	// * generating random project name and license for initialisation
	rlicense := generateRandom("license")
	rprojectName := generateRandom("name")
	init := Initialiser{
		projectName: rprojectName,
		license: rlicense,
		lang: "web",
	}
	init.webInit()

	// * getting contents of the project initialised
	projectFileNames := getFiles("./")

	// * checking for various files
	fmt.Println(len(projectFileNames), projectFileNames)
	if len(projectFileNames) != 8 {t.Errorf("proper structure not made")}

	if !stringInSlice("LICENSE", projectFileNames) {t.Errorf("LICENSE file not present.")}
	if !stringInSlice("README.md", projectFileNames) {t.Errorf("README.md file not present.")}
	if !stringInSlice(".gitignore", projectFileNames) {t.Errorf(".gitignore file not present.")}
	if !stringInSlice("index.html", projectFileNames) {t.Errorf("index.html file not present.")}
	if !stringInSlice("css", projectFileNames) {t.Errorf("css dir not present.")}
	if !stringInSlice("js", projectFileNames) {t.Errorf("js dir not present.")}
	if !stringInSlice("img", projectFileNames) {t.Errorf("img dir not present.")}

	cssFiles := getFiles("./css")
	if !(stringInSlice("style.css", cssFiles)) {t.Errorf("style.css not present")}

	jsFiles := getFiles("./js")
	if !(stringInSlice("script.js", jsFiles)) {t.Errorf("script.js not present")}


	t.Cleanup(func() {
		t.Log("Cleaning up...")
		os.Chdir(gwd)
		os.RemoveAll(rprojectName)
	})
}