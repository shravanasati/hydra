package main

import (
	"fmt"
	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("hydra").
		SetVersion("1.0.0").
		SetDescription("hydra is command line utility to generate language-specific project structure. \nFor more detailed information and documentation, visit https://github.com/Shravan-1908/hydra. \nOr try the command `hydra --help`.")

	// * the config command
	commando.
		Register("config").
		SetShortDescription("Alter or set the configurations for hydra.").
		SetDescription("The config command is used to configure settings for hydra, which are used when a project is intialised using hydra. \nIt consists of following flags:...").
		AddFlag("name", "The full name of the user. It is used to create licenses.", commando.String, "name").
		AddFlag("github-username", "The GitHub username of the user. It is used to create src files and intialise the project.", commando.String, "username").
		AddFlag("default-lang", "The default language for project initialisation. Once set, if no `lang` argument is provided with the `init` command, hydra falls back to the default language project structure.", commando.String, "go").
		AddFlag("default-license", "The default license for project initialisation. Options are:\n 1. MIT \n 2. Apache \n 3. MPL \n 4. GNU GPL v3", commando.String, "MIT").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			config((flags["name"].Value).(string), flags["github-username"].Value.(string), flags["default-lang"].Value.(string), flags["default-license"].Value.(string))
		})

	commando.
		Register("init").
		SetDescription("Intialises the project structure.\n\nUsage: name : project name \n lang : programming language in which the project is being built.").
		SetShortDescription("Intialises the project structure.").
		SetDescription("Long description here").
		AddArgument("name", "Name of the project", "").
		AddArgument("lang", "Language of the project you're working on.", "default").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
			// init()
		})


	commando.Parse(nil)
}
