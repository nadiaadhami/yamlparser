package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
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
	parse1()
	parse2()
	parse3()
	parseAD()
	parseCountry()
}
func parse1() {
	filename, _ := filepath.Abs("./test1.yml")
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
	filename, _ := filepath.Abs("./test2.yml")
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
	filename, _ := filepath.Abs("./test3.yml")
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
	filename, _ := filepath.Abs("./countries/AD.yaml")
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
func parseCountry() {
	filename, _ := filepath.Abs("./countries/AD.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var config Country
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config)
	//fmt.Println("Geo", config.Geo)
	fmt.Println("UnofficialNames", config.UnofficialNames)

}
