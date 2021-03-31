package main

import (
	"fmt"
	"github.com/thatisuday/commando"
)

const (
	VERSION string = "1.0.0"
	NAME string = "hydra"
)

func main() {
	fmt.Println(NAME, VERSION)

	// * basic configuration
	commando.
		SetExecutableName(NAME).
		SetVersion(VERSION).
		SetDescription("hydra is command line utility to generate language-specific project structure. \nFor more detailed information and documentation, visit https://github.com/Shravan-1908/hydra . \n")

	// * the init command
	commando.
		Register("init").
		SetDescription("Intialises the project structure.\n\nUsage: \n name : project name \n lang : programming language in which the project is being built.").
		SetShortDescription("Intialises the project structure.").
		SetDescription("The `init` command initialises the project structure.").
		AddArgument("name", "Name of the project", "").
		AddArgument("lang", "Language/framework of the project.", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

			projectName := args["name"].Value
			projectLang := args["lang"].Value
			if projectLang == "python" {
				pythonInit(projectName)
			} else if projectLang == "go" {
				goInit(projectName)
			} else {
				fmt.Printf("Unsupported language type: %v. Cannot initiate the project.", projectLang)
			}
			
		})


	commando.Parse(nil)
}