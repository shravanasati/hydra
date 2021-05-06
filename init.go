/*
The following code is responsible for the init command.

Author: Shravan Asati
Originally Written: 28 March 2021
Last edited: 6 May 2021
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"strconv"
	"embed"
)

// just to use the embed package
var _ embed.FS


// * all licenses 
//go:embed .\licenses\APACHE
var APACHE string
//go:embed .\licenses\BSD
var BSD string
//go:embed .\licenses\EPL
var EPL string
//go:embed .\licenses\GPL
var GPL string
//go:embed .\licenses\MIT
var MIT string
//go:embed .\licenses\MPL
var MPL string

// * all gitignores
//go:embed .\gitignores\go.gitignore
var goGitignore string
//go:embed .\gitignores\python.gitignore
var pythonGitignore string


func year() string {
	return strconv.Itoa(time.Now().Year())
}

func handleException(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func makeFile(filename, content string) {
	f, e := os.Create(filename)
	handleException(e)
	_, er := f.WriteString(content)
	handleException(er)
	f.Close()
	cwd, _ := os.Getwd()
	fmt.Printf("\n - Created file '%v' at %v.", filename, cwd)
}

func makeDir(dirname string) {
	os.Mkdir(dirname, os.ModePerm)
	cwd, _ := os.Getwd()
	fmt.Printf("\n - Created directory '%v' at %v.", dirname, cwd)
}

func execute(base string, command ...string) error {
	cmd := exec.Command(base, command...)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func getGitignore(language string) string {
	switch language {
	case "python":
		return pythonGitignore
	case "go":
		return goGitignore
	default:
		return fmt.Sprintf("Unknown language: %v.", language)
	}
}

func getLicense(license string) string {
	switch license {
	case "MIT":
		licenseText := MIT
		licenseText = strings.Replace(licenseText, ":YEAR:", year(), 1)
		licenseText = strings.Replace(licenseText, ":NAME:", getConfig("fullName"), 1)
		return licenseText
	case "BSD":
		licenseText := BSD
		licenseText = strings.Replace(licenseText, ":YEAR:", year(), 1)
		licenseText = strings.Replace(licenseText, ":NAME:", getConfig("fullName"), 1)
		return licenseText
	case "APACHE":
		licenseText := APACHE
		licenseText = strings.Replace(licenseText, ":YEAR:", year(), 1)
		licenseText = strings.Replace(licenseText, ":NAME:", getConfig("fullName"), 1)
		return licenseText
	case "EPL":
		licenseText := EPL
		licenseText = strings.Replace(licenseText, ":YEAR:", year(), 1)
		licenseText = strings.Replace(licenseText, ":NAME:", getConfig("fullName"), 1)
		return licenseText
	case "MPL":
		licenseText := MPL
		licenseText = strings.Replace(licenseText, ":YEAR:", year(), 1)
		licenseText = strings.Replace(licenseText, ":NAME:", getConfig("fullName"), 1)
		return licenseText
	case "GPL":
		licenseText := GPL
		licenseText = strings.Replace(licenseText, ":YEAR:", year(), 1)
		licenseText = strings.Replace(licenseText, ":NAME:", getConfig("fullName"), 1)
		return licenseText
	default:
		return fmt.Sprintf("Undefined license: %v.", license)
	}
	
}

func pythonInit(projectName, license string) {
	fmt.Printf("Initialising project: '%v'.\n", projectName)

	makeDir(projectName)
	os.Chdir(fmt.Sprintf("./%v", projectName))

	gwd, _ := os.Getwd()
	makeFile("LICENSE", getLicense(license))
	makeFile("README.md", fmt.Sprintf("# %v", projectName))
	makeFile(".gitignore", getGitignore("python"))
	makeFile("setup.py", "from setuptools import setup \n\nsetup()")

	makeDir(projectName)
	os.Chdir(fmt.Sprintf("./%v", projectName))
	makeFile("__init__.py", "")
	os.Chdir(gwd)

	makeDir("tests")
	os.Chdir("./tests")
	makeFile("__init__.py", "")
	makeFile(fmt.Sprintf("test_%v.py", projectName), "")
	os.Chdir(gwd)

	e := execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a git repository.")
	} else {
		fmt.Println("\n - Intialised a Git repository for your project.")
	}
}

func goInit(projectName, license string) {
	fmt.Printf("Initialising project: '%v'\n.", projectName)

	makeDir(projectName)
	os.Chdir(fmt.Sprintf("./%v", projectName))

	gwd, _ := os.Getwd()
	makeFile("LICENSE", getLicense(license))
	makeFile("README.md", fmt.Sprintf("# %v", projectName))
	makeFile(".gitignore", getGitignore("go"))

	makeDir("src")
	os.Chdir("./src")
	makeFile("main.go", "package main")
	os.Chdir(gwd)

	makeDir("tests")
	os.Chdir("./tests")
	makeFile(fmt.Sprintf("%v_test.go", projectName), "package main")
	os.Chdir(gwd)

	makeDir("bin")
	makeDir("pkg")

	e := execute("go", "mod", "init", fmt.Sprintf("github.com/%v/%v", getConfig("githubUsername"), projectName))
	if e != nil {
		fmt.Println("\n ** Go isn't installed on your system. Cannot enable dependency tracking.")
	} else {
		fmt.Println("\n - Enabled dependency tracking for your Go application.")
	}
	e = execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a repository.")
	} else {
		fmt.Println(" - Intialised a Git repository for your project.")
	}
}

var HTMLBoilerplate string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>:PROJECT_NAME:</title>
</head>
<body>
    <h1>:PROJECT_NAME:</h1>
</body>
</html>
`

var cssReset string = `
* {
	margin: 0px;
	padding: 0px;
	box-sizing: border-box;
	border: 0;
	font-size: 100%;
}
`

func webInit(projectName, license string) {
	fmt.Printf("Initialising project: '%v'.\n", projectName)

	makeDir(projectName)
	os.Chdir(fmt.Sprintf("./%v", projectName))

	gwd, _ := os.Getwd()
	makeFile("LICENSE", getLicense(license))
	makeFile("index.html", strings.Replace(HTMLBoilerplate, ":PROJECT_NAME:", projectName, 2))
	makeFile("README.md", fmt.Sprintf("# %v", projectName))

	makeDir("img")
	
	makeDir("css")
	os.Chdir("./css")
	makeFile("style.css", cssReset)
	os.Chdir(gwd)

	makeDir("js")
	os.Chdir("./js")
	makeFile("script.js", "")
	os.Chdir(gwd)

	e := execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a repository.")
	} else {
		fmt.Println(" - Intialised a Git repository for your project.")
	}
}

func flaskInit(projectName, license string)  {
	fmt.Printf("Initialising project: '%v'\n.", projectName)

	makeDir(projectName)
	os.Chdir(fmt.Sprintf("./%v", projectName))

	gwd, _ := os.Getwd()
	makeFile("LICENSE", getLicense(license))
	makeFile("README.md", fmt.Sprintf("# %v", projectName))
	makeFile(".gitignore", getGitignore("python"))
	makeFile("app.py", "from flask import Flask\n\napp = Flask(__name__)")


	makeDir("static")
	os.Chdir("./static")

	makeDir("images")

	makeDir("scripts")
	os.Chdir("./scripts")
	makeFile("script.js", "")
	os.Chdir("..")

	makeDir("styles")
	os.Chdir("./styles")
	
	makeFile("style.css", cssReset)
	os.Chdir(gwd)


	makeDir("templates")
	os.Chdir("./templates")
	makeFile("index.html", strings.Replace(HTMLBoilerplate, ":PROJECT_NAME:", projectName, 2))
	os.Chdir(gwd)

	e := execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a git repository.")
	} else {
		fmt.Println("\n - Intialised a Git repository for your project.")
	}
}