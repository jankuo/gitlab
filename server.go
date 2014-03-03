package main

import (
	"os"
	"strings"
	"text/template"
)

// A simple template definition to test our function.
// We print the input text several ways.
//  - the original
//  - title-cased
//  - title-cased and then printed whith %q
//  - printed with %q and then title-cased.
const templateText = `
Input: {{ printf "%q" . }}
Output 0 : {{ title . }}
Output 1 : {{ title . | printf "%q " }}
Output 2 : {{ printf "%q" . | title }}
`

func main() {

	//First we create a FuncMap with to register the function.
	funcMap := template.FuncMap{
		//The name "Title" is what function will be called in the template
		"title": strings.Title,
	}

	t, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, "the go programing language")
	if err != nil {
		panic(err)
	}
}
