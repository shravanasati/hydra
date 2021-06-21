package main

import (
	"fmt"
	"math/rand"
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
	t.Cleanup(func() {
		t.Log("\nCleaning up...\n")
		config(initialName, initialGithubUsername, initialLang, initialLicense)
	})
}
