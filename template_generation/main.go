package main

import (
	"fmt"

	"os"
	"text/template"
	"www.github.com/friday963/gonetworkautomation/template_generation/structures"
	"www.github.com/friday963/gonetworkautomation/template_generation/templatefiles"
	"gopkg.in/yaml.v3"
)



func main() {
	data, err := os.ReadFile("input.yml")
	if err != nil {
		fmt.Println(err)
	}
	var config structures.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
	}	

	t, err := template.New("config").Parse(templatefiles.CeosTemplate)
	err = t.Execute(os.Stdout, config)

}
