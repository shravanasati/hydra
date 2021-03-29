package main

import (
	"fmt"
	"strconv"
	"time"
	"os"
	"os/exec"
)

func year() string {
	return strconv.Itoa(time.Now().Year())
}

func getGitignore(language string) string {
	return ""
}

func getLicense(license string) string {
	return ""
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

func pythonInit(projectName, license string) {
	fmt.Printf("Initialising project: '%v'.\n", projectName)
	gwd, _ := os.Getwd()
	makeFile("LICENSE", license)
	makeFile("README.md", fmt.Sprintf("# %v", projectName))
	makeFile(".gitignore", "pygitignore")
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
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a repository.")
	} else {
		fmt.Println("\n - Intialised a Git repository for your project.")
	}
}

func goInit(projectName, license string) {
	fmt.Printf("Initialising project: '%v'\n.", projectName)
	gwd, _ := os.Getwd()
	makeFile("LICENSE", license)
	makeFile("README.md", fmt.Sprintf("# %v", projectName))
	makeFile(".gitignore", "gogitignore")

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

	e := execute("go", "mod", "init", projectName)
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