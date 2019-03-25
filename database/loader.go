package database

import (
	"client-integration-go/ecard-prototype-go/logger"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"yamlparser/models"
)

func LoadCountries(dataDir string) map[string]models.Country {
	countriesFile := dataDir + "countries.yaml"
	logger.LogInfo("LoadCountries", countriesFile)
	cArray := ParseCountriesFile(countriesFile)
	m:= LoadCountriesList(dataDir, cArray)
	return m
}
func LoadCountriesList(dataDir string, arr []string) map[string]models.Country {
	var m map[string]models.Country
	m = make(map[string]models.Country)

	for _, v := range arr {
		//fmt.Println(i, v)
		pathToFile := dataDir+"countries/" + v + ".yaml"
		c := ParseCountry(pathToFile)
		//fmt.Println( i, c.Name, c.Alpha2)
		m[c.Alpha2] = c
	}
	fmt.Println("Number of countries loaded:", len(m))
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
	//fmt.Printf("Value: %#v\n", config)
	return config
}
func ParseCountriesFile(f string) [] string {
	filename, _ := filepath.Abs(f)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var sa models.Countries
	//type StringArray []string
	err = yaml.Unmarshal(yamlFile, &sa)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", sa)
	return sa
}
