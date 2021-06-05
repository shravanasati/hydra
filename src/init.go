/*
The following code is responsible for the init command.

Author: Shravan Asati
Originally Written: 28 March 2021
Last edited: 8 May 2021
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"strconv"
)

// year returns the current year, used for editing LICENSE file.
func year() string {
	return strconv.Itoa(time.Now().Year())
}

// handleException handles the exception by printing it and exiting the program.
func handleException(err error) {
	if err != nil {
		fmt.Println("FATAL ERROR: Project initialisation failed! This should never happen. You may want to file an issue at the hydra repository: https://github.com/Shravan-1908/hydra/issues/new?assignees=&labels=&template=bug_report.md&title=")
		fmt.Println(err)
		os.Exit(-1)
	}
}

type Initialiser struct {
	projectName string
	license string
	lang string
}

// makeFile creates a file with the provided content.
func makeFile(filename, content string) {
	f, e := os.Create(filename)
	handleException(e)
	_, er := f.WriteString(content)
	handleException(er)
	f.Close()
	cwd, _ := os.Getwd()
	fmt.Printf("\n - Created file '%v' at %v.", filename, cwd)
}

// makeDir creates a directory.
func makeDir(dirname string) {
	os.Mkdir(dirname, os.ModePerm)
	cwd, _ := os.Getwd()
	fmt.Printf("\n - Created directory '%v' at %v.", dirname, cwd)
}

// execute executes a command in the shell.
func execute(base string, command ...string) error {
	cmd := exec.Command(base, command...)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

// getGitignore returns the gitignore variable from static.go, corresponding to the provided language. 
func getGitignore(language string) string {
	switch language {
	case "python":
		return pythonGitignore
	case "go":
		return goGitignore
	case "c":
		return cGitignore
	case "c++":
		return cppGitignore
	case "ruby":
		return rubyGitignore
	default:
		return fmt.Sprintf("Unknown language: %v.", language)
	}
}

// manipulateLicense replaces the `:NAME:` and `:YEAR:` values of the license with actual values.
func manipulateLicense(license string) string {
	licenseText := strings.Replace(license, ":YEAR:", year(), 1)
	licenseText = strings.Replace(licenseText, ":NAME:", getConfig("fullName"), 1)

	return licenseText
}

// getGitignore returns the license variable from static.go, corresponding to the provided license. 
func getLicense(license string) string {
	switch license {
	case "MIT":
		return manipulateLicense(MIT)
	case "BSD":
		return manipulateLicense(BSD)
	case "APACHE":
		return manipulateLicense(APACHE)
	case "EPL":
		return manipulateLicense(EPL)
	case "MPL":
		return manipulateLicense(MPL)
	case "GPL":
		return manipulateLicense(GPL)
	case "UNI":
		return manipulateLicense(UNI)
	default:
		return fmt.Sprintf("Undefined license: %v.", license)
	}
}

// basicInit makes the README, LICENSE and gitignore files.
func (init *Initialiser) basicInit() string {
	fmt.Printf("Initialising project: '%v' in %v.\n", init.projectName, init.lang)

	makeDir(init.projectName)
	os.Chdir(fmt.Sprintf("./%v", init.projectName))

	gwd, _ := os.Getwd()
	makeFile("LICENSE", getLicense(init.license))
	makeFile("README.md", fmt.Sprintf("# %v", init.projectName))
	makeFile(".gitignore", getGitignore(init.lang))
	return gwd
}

// pythonInit is the python project initialisation function.
func (init *Initialiser) pythonInit() {
	gwd := init.basicInit()

	setupContent = strings.Replace(setupContent, ":PROJECT_NAME:", init.projectName, 2)
	setupContent = strings.Replace(setupContent, ":LICENSE:", init.license, 1)
	setupContent = strings.Replace(setupContent, ":GITHUB:", getConfig("githubUsername"), 1)
	setupContent = strings.Replace(setupContent, ":AUTHOR_NAME:", getConfig("fullName"), 1)
	makeFile("setup.py", setupContent)

	makeDir(init.projectName)
	os.Chdir(fmt.Sprintf("./%v", init.projectName))
	makeFile("__init__.py", "")
	os.Chdir(gwd)

	makeDir("tests")
	os.Chdir("./tests")
	makeFile("__init__.py", "")
	makeFile(fmt.Sprintf("test_%v.py", init.projectName), "")
	os.Chdir(gwd)

	e := execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a git repository.")
	} else {
		fmt.Println("\n - Intialised a Git repository for your project.")
	}
}


// goInit is the go project initialisation function.
func (init *Initialiser) goInit() {
	gwd := init.basicInit()

	makeDir("src")
	os.Chdir("./src")
	makeFile("main.go", "package main")
	os.Chdir(gwd)

	makeDir("tests")
	os.Chdir("./tests")
	makeFile(fmt.Sprintf("%v_test.go", init.projectName), "package main")
	os.Chdir(gwd)

	makeDir("bin")
	makeDir("pkg")

	e := execute("go", "mod", "init", fmt.Sprintf("github.com/%v/%v", getConfig("githubUsername"), init.projectName))
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


// webInit is the web-frontend project initialisation function.
func (init *Initialiser) webInit() {
	gwd := init.basicInit()
	
	indexContent := strings.Replace(HTMLBoilerplate, ":PROJECT_NAME:", init.projectName, 2)
	indexContent = strings.Replace(indexContent, ":CSS_LINK:", `<link rel="stylesheet" href="./css/style.css">`, 1)
	indexContent = strings.Replace(indexContent, ":SCRIPT_LINK:", `<script src="./js/script.js"> </script>`, 1)
	makeFile("index.html", indexContent)
	makeFile("README.md", fmt.Sprintf("# %v", init.projectName))

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


// flaskInit is the python-flask project initialisation function.
func (init *Initialiser) flaskInit()  {
	gwd := init.basicInit()

	makeFile("app.py", flaskBoilerplate)

	// * making the static dir which contains images, styles and scripts dir
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


	// * making the templates dir
	makeDir("templates")
	os.Chdir("./templates")
	indexContent := strings.Replace(HTMLBoilerplate, ":PROJECT_NAME:", init.projectName, 2)
	indexContent = strings.Replace(indexContent, ":CSS_LINK:", `<link rel="stylesheet" href="{{ url_for('static', filename='styles/style.css') }}">`, 1)
	indexContent = strings.Replace(indexContent, ":SCRIPT_LINK:", `<script src=" {{ url_for('static', filename='scripts/script.js') }} "> </script>`, 1)
	makeFile("index.html", indexContent)
	os.Chdir(gwd)

	// * initialising git repository
	e := execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a git repository.")
	} else {
		fmt.Println("\n - Intialised a Git repository for your project.")
	}
}


// cInit is the C project initialisation function.
func (init *Initialiser) cInit() {
	gwd := init.basicInit()

	makeFile("Makefile.am", "")

	makeDir("src")
	os.Chdir("./src")
	makeFile("Makefile.am", "")
	makeFile("main.c", "")
	makeFile("main.h", "")
	os.Chdir(gwd)

	makeDir("tests")
	os.Chdir("./tests")
	makeFile("Makefile.am", "")
	makeFile(fmt.Sprintf("%v_test.c", init.projectName), "")
	os.Chdir(gwd)

	makeDir("libs")
	os.Chdir("../libs")
	makeFile("Makefile.am", "")

	e := execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a repository.")
	} else {
		fmt.Println(" - Intialised a Git repository for your project.")
	}
}

// cppInit is the C++ project initialisation function.
func (init *Initialiser) cppInit() {
	gwd := init.basicInit()
	makeFile("CMakeLists.txt", "")

	makeDir("src")
	os.Chdir("./src")
	makeFile("main.cpp", "")
	makeFile("main.h", "")
	os.Chdir(gwd)

	makeDir("include")
	os.Chdir("./include")
	makeDir(init.projectName)
	os.Chdir(fmt.Sprintf("./%v", init.projectName))
	makeFile("header.h", "")
	os.Chdir(gwd)

	makeDir("tests")
	os.Chdir("./tests")
	makeFile(fmt.Sprintf("%v_test.cpp", init.projectName), "")
	os.Chdir(gwd)

	makeDir("libs")

	e := execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a repository.")
	} else {
		fmt.Println(" - Intialised a Git repository for your project.")
	}
}


// rubyInit is the ruby project initialisation function.
func (init *Initialiser) rubyInit() {
	gwd := init.basicInit()

	makeFile("Gemfile", "")
	makeFile("Rakefile", "")

	gemspecContent = strings.Replace(gemspecContent, ":PROJECT_NAME:", init.projectName, 4)
	gemspecContent = strings.Replace(gemspecContent, ":LICENSE:", init.license, 1)
	gemspecContent = strings.Replace(gemspecContent, ":GITHUB:", getConfig("githubUsername"), 1)
	gemspecContent = strings.Replace(gemspecContent, ":AUTHOR_NAME:", getConfig("fullName"), 1)
	makeFile(fmt.Sprintf("%v.gemspec", init.projectName), gemspecContent)

	makeDir("bin")

	makeDir("lib")
	os.Chdir("./lib")
	makeFile(fmt.Sprintf("%v.rb", init.projectName), "")
	os.Chdir(gwd)

	makeDir("tests")
	os.Chdir("./tests")
	makeFile(fmt.Sprintf("test_%v.rb", init.projectName), "")
	os.Chdir(gwd)

	e := execute("git", "init")
	if e != nil {
		fmt.Println("\n ** Git isn't installed on your system. Cannot initiate a git repository.")
	} else {
		fmt.Println("\n - Intialised a Git repository for your project.")
	}
}
