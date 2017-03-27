package main

import (
	"text/template"
	"os"
	"log"
)

type CarDetails struct{
	Name string
	Img string
}

var details *template.Template //define a global pointer to a template

func init() { //parse everything with .html extension
	details = template.Must(template.ParseGlob("templates/*.html")) //gives you back a pointer to a template and checks for errors
}

func main() {

	mustang := CarDetails{
		Name: "Ford Mustang",
		Img:  "src='http://placehold.it/450x250'",
	}

/*
	err := details.ExecuteTemplate(os.Stdout,"photoGal.html",mustang)
	if err != nil{
		log.Fatalln(err)
	}
*/

	err := details.ExecuteTemplate(os.Stdout,"details.html",mustang)
	if err != nil{
		log.Fatalln(err)
	}
}