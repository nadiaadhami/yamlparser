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
	us := "./database/data/subdivisions/US.yaml"
	ad := "./database/data/subdivisions/AD.yaml"

	parse1()
	parse2()
	parse3()
	parseAD()
	database.ParseCountry(f)
	m := database.LoadCountries(d)
	fmt.Println("Number of countries =", len(m))
	fmt.Println("AD UnofficialNames =", m["AD"].UnofficialNames)

	// parse subdivisions
	usSub := database.ParseSubdivision(us)
	//fmt.Println("US subdivisions: ", sub)
	fmt.Println("Number of US subdivisions: ", len(usSub))
	adSub := database.ParseSubdivision(ad)
	fmt.Println("\nNumber of AD subdivisions: ", len(adSub))
	fmt.Println("\ncountry = US")
	subdivisionName := "AA"
	fmt.Println("subdivision:", subdivisionName)
	fmt.Println(" subdivision Name =", usSub[subdivisionName].Name)
	fmt.Println(" Translations =", usSub[subdivisionName].Translations)
	fmt.Println(" Translations size =", len(usSub[subdivisionName].Translations))
	fmt.Println(" UnofficialNames =", usSub[subdivisionName].UnofficialNames)

	subdivisionName = "AK"
	fmt.Println("subdivision:", subdivisionName)

	fmt.Println(" subdivision Name =", usSub[subdivisionName].Name)
	fmt.Println(" Translations =", usSub[subdivisionName].Translations)
	fmt.Println(" Translations size =", len(usSub[subdivisionName].Translations))
	fmt.Println(" UnofficialNames =", usSub[subdivisionName].UnofficialNames)

	fmt.Println("country = BE")
	be := "./database/data/subdivisions/BE.yaml"
	beSubs := database.ParseSubdivision(be)
	subdivisionName = "BRU"
	fmt.Println("subdivision:", subdivisionName)
	fmt.Println("subdivision Name =", beSubs[subdivisionName].Name)
	fmt.Println("subdivision Translations =", beSubs[subdivisionName].Translations)
	fmt.Println("subdivision Translations size =", len(beSubs[subdivisionName].Translations))
	fmt.Println("subdivision UnofficialNames =", beSubs[subdivisionName].UnofficialNames)

	subs := database.LoadSubdivisions(d)
	fmt.Println("\nTotal subdivision =", len(subs))

	stateMap := database.GetStates("US", "./database/data/subdivisions/US.yaml")
	database.PrintStates(stateMap)
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
