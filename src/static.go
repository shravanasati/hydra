/*
This file contains static data like variables (gitignores and licenses) which are to be
embed and boilerplates for various programming languages.

Author: Shravan Asati
Originally Written: 8 May 2021
Last edited: 4 June 2021
*/

package main

import _ "embed"


// * all licenses 

//go:embed licenses/APACHE
var APACHE string
//go:embed licenses/BSD
var BSD string
//go:embed licenses/EPL
var EPL string
//go:embed licenses/GPL
var GPL string
//go:embed licenses/MIT
var MIT string
//go:embed licenses/MPL
var MPL string
//go:embed licenses/UNI
var UNI string


// * all gitignores
//go:embed gitignores/go.gitignore
var goGitignore string
//go:embed gitignores/python.gitignore
var pythonGitignore string
//go:embed gitignores/c.gitignore
var cGitignore string
//go:embed gitignores/cpp.gitignore
var cppGitignore string
//go:embed gitignores/ruby.gitignore
var rubyGitignore string



// * all boilerplates

//go:embed boilerplates/html
var HTMLBoilerplate string

//go:embed boilerplates/cssReset
var cssReset string

//go:embed boilerplates/flask
var flaskBoilerplate string 

//go:embed boilerplates/gemspec
var gemspecContent string

//go:embed boilerplates/setupContent
var setupContent string