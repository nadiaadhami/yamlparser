package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"yamlparser/database"
)

type Config struct {
	Firewall_network_rules map[string]Options
}
type Config2 struct {
	Firewall_network_rules Options
}
type Config3 struct {
	Src string `yaml:"src"`
	Dst string `yaml:"dst"`
}
type Options struct {
	Src string `yaml:"src"`
	Dst string `yaml:"dst"`
}
type CountryMock struct {
	Name   string `yaml:"name"`
	Alpha2 string `yaml:"alpha2"`
}

// reference : https://stackoverflow.com/questions/28682439/go-parse-yaml-file

func main() {
	f := "./database/data/countries/AD.yaml"
	d := "./database/data/"

	parse1()
	parse2()
	parse3()
	parseAD()
	database.ParseCountry(f)
	m := database.LoadCountries(d)
	fmt.Println("Number of countries =", len(m))

}

func parse1() {
	filename, _ := filepath.Abs("./database/data/test/test1.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config.Firewall_network_rules)
}
func parse2() {
	filename, _ := filepath.Abs("./database/data/test/test2.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var config Config2
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config.Firewall_network_rules)
}
func parse3() {
	filename, _ := filepath.Abs("./database/data/test/test3.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var config Config3
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config)
}
func parseAD() {
	filename, _ := filepath.Abs("./database/data/countries/AD.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var config CountryMock
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config)
}
