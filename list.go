/*
This file contains code which is responsible for the `list` command.

Author: Shravan Asati
Originally Written: 15 April 2021
Last edited: 15 April 2021
*/

package main

import "fmt"

func list(item string) {
	if item == "langs" {
		fmt.Printf("\n%v %v supports following languages for project initialisation: \n", NAME, VERSION)
		for i, v := range supportedLangs {
			fmt.Printf("%v. %v \n", i+1, v)
		}
	
	} else if item == "licenses" {
		fmt.Printf("\n%v %v supports following licenses for project initialisation: \n", NAME, VERSION)
		i := 1
		for k, v := range supportedLicenses {
			fmt.Printf("%v. %v - %v \n", i, k, v)
			i++
		}
	
	} else if item == "configs" {
		config("default", "default", "default", "default")
		fmt.Println("\nThe hydra user configurations are:")
		fmt.Println("Full Name:", getConfig("fullName"))
		fmt.Println("GitHub Username:", getConfig("githubUsername"))
		fmt.Println("Default Language:", getConfig("defaultLang"))
		fmt.Println("Default License:", getConfig("defaultLicense"))
		fmt.Println("\nTo know how to set the configuration, type in `hydra config -h`.")
	
	} else {
		fmt.Printf("Invalid value for the 'item' argument: '%v'.\nSee `hydra list -h` for help.", item)
	} 
}