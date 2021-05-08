/*
This file contains static data like variables (gitignores and licenses) which are to be
embed and boilerplates for various programming languages.

Author: Shravan Asati
Originally Written: 8 May 2021
Last edited: 8 May 2021
*/

package main

import _ "embed"


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
//go:embed .\licenses\UNI
var UNI string


// * all gitignores

//go:embed .\gitignores\go.gitignore
var goGitignore string
//go:embed .\gitignores\python.gitignore
var pythonGitignore string
//go:embed .\gitignores\c.gitignore
var cGitignore string
//go:embed .\gitignores\c++.gitignore
var cppGitignore string
//go:embed .\gitignores\ruby.gitignore
var rubyGitignore string


// * all boilerplates

var HTMLBoilerplate string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>:PROJECT_NAME:</title>
	:CSS_LINK:
</head>

<body>
    <h1>:PROJECT_NAME:</h1>

	:SCRIPT_LINK:
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

var flaskBoilerplate string = `
from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def home():
	return render_template('index.html')

app.run(debug=True)
`

var gemspecContent string = `
Gem::Specification.new do |s|
	s.name        = ":PROJECT_NAME:"
	s.version     = '1.0.0'
	s.license     = ":LICENSE:"
	s.summary     = "Project summary here"
	s.description = "Much longer explanation of the project."
	s.authors     = [":AUTHOR_NAME:"]
	s.email       = 'Your email here.'
	s.files       = ["lib/:PROJECT_NAME:.rb"]
	s.homepage    = 'https://rubygems.org/gems/:PROJECT_NAME:'
	s.metadata    = { "source_code_uri" => "https://github.com/:GITHUB:/:PROJECT_NAME:" }
end
`

var setupContent string = `
from setuptools import find_packages, setup

VERSION = '1.0.0'
with open("README.md") as f:
    README = f.read()

setup(
    name = ":PROJECT_NAME:",
    version = VERSION,
    description = "Project summary here",
    long_description_content_type = "text/markdown",
    long_description = README,
    url = "https://github.com/:GITHUB:/:PROJECT_NAME:",
    author = ":AUTHOR_NAME:",
    author_email = "Your email here",
    packages = find_packages(),
    install_requires = [],
    license = ':LICENSE:',
    keywords = [],
    classifiers = []
)
`