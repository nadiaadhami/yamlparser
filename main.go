package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"yamlparser/models"
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
// todo parse Geo

func main() {
	countryFile := "./data/countries/AD.yaml"
	countriesFile := "./data/countries.yaml"
	parse1()
	parse2()
	parse3()
	parseAD()
	parseCountry(countryFile)
	cArray := parseCountries(countriesFile)
	loadCountries(cArray)
}
func parse1() {
	filename, _ := filepath.Abs("./data/test/test1.yml")
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
	filename, _ := filepath.Abs("./data/test/test2.yml")
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
	filename, _ := filepath.Abs("./data/test/test3.yml")
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
	filename, _ := filepath.Abs("./data/countries/AD.yaml")
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
func parseCountry(f string) models.Country {
	filename, _ := filepath.Abs(f)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var config models.Country
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config)
	//fmt.Println("Geo", config.Geo)
	//fmt.Println("UnofficialNames", config.UnofficialNames)
	return config
}
func parseCountries(f string) [] string {
	filename, _ := filepath.Abs(f)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var cArray models.Countries
	//type StringArray []string
	//var sa StringArray
	err = yaml.Unmarshal(yamlFile, &cArray)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", cArray)
	return cArray
}
func loadCountries(arr []string) {
	for i, v := range arr {
		fmt.Println(i, v)
		pathToFile := "./data/countries/"+v+".yaml"
		c := parseCountry(pathToFile)
		fmt.Println("======", i, c.Name, c.Alpha2)
	}
}