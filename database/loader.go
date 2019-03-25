package database

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"yamlparser/models"
)

func LoadCountriesList(arr []string) map[string]models.Country {
	var m map[string]models.Country
	m = make(map[string]models.Country)

	for i, v := range arr {
		//fmt.Println(i, v)
		pathToFile := "./database/data/countries/" + v + ".yaml"
		c := ParseCountry(pathToFile)
		fmt.Println("======", i, c.Name, c.Alpha2)
		m[c.Alpha2] = c
	}
	return m
}

func ParseCountry(f string) models.Country {
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
func ParseCountriesFile(f string) [] string {
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
