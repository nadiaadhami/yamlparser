package models

type Subdivision struct {
	Name            string         `yaml:"name"`
	UnofficialNames string         `yaml:"unofficial_names"`
	Translations    TranslationMap `yaml:"translations"`
	Geo             Geo            `yaml:"geo"`
}

type Subdivisions map[string]Subdivision

type TranslationMap map[string]string

