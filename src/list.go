/*
This file contains code which is responsible for the `list` command.

Author: Shravan Asati
Originally Written: 15 April 2021
Last edited: 15 April 2021
*/

package main

import "fmt"

func list(item string) string {
	if item == "langs" {
		fmt.Printf("\n%v %v supports following languages for project initialisation: \n", NAME, VERSION)
		content := ""
		for i, v := range supportedLangs {
			content += fmt.Sprintf("%v. %v \n", i+1, v)
		}
		return content

	} else if item == "licenses" {
		fmt.Printf("\n%v %v supports following licenses for project initialisation: \n", NAME, VERSION)
		i := 1
		content := ""
		for k, v := range supportedLicenses {
			content += fmt.Sprintf("%v. %v - %v \n", i, k, v)
			i++
		}
		return content

	} else if item == "configs" {
		config("default", "default", "default", "default")
		content := "\nThe hydra user configurations are: \n"
		content += fmt.Sprintf("Full Name: %v \n", getConfig("fullName"))
		content += fmt.Sprintf("GitHub Username: %v \n", getConfig("githubUsername"))
		content += fmt.Sprintf("Default Language: %v \n", getConfig("defaultLang"))
		content += fmt.Sprintf("Default License: %v \n", getConfig("defaultLicense"))
		content += fmt.Sprintln("\nTo know how to set the configuration, type in `hydra config -h`.")
		return content

	} else {
		return fmt.Sprintf("Invalid value for the 'item' argument: '%v'.\nSee `hydra list -h` for help.", item)
	}
}
