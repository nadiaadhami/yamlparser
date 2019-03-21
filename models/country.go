package models

type Country struct {
	ID int64 `yaml:"column(id);pk;auto"`

	Number                         string   `yaml:"number"`
	Alpha2                         string   `yaml:"alpha2"`
	Alpha3                         string   `yaml:"alpha3"`
	Currency                       string   `yaml:"currency"`
	Name                           string   `yaml:"name"`
	UnofficialNames                []string `yaml:"unofficial_names"`
	Continent                      string   `yaml:"continent"`
	Region                         string   `yaml:"region"`
	Subregion                      string   `yaml:"subregion"`
	Geo                            Geo      `yaml:"geo"`
	WorldRegion                    string   `yaml:"world_region"`
	CountryCode                    string   `yaml:"country_code"`
	NationalDestinationCodeLengths []int    `yaml:"national_destination_code_lengths"`
	NationalNumberLengths          []int    `yaml:"national_number_lengths"`
	InternationalPrefix            string   `yaml:"international_prefix"`
	NationalPrefix                 string   `yaml:"national_prefix"`
	Ioc                            string   `yaml:"ioc"`
	Gec                            string   `yaml:"gec"`
	UnLocode                       string   `yaml:"un_locode"`
	LanguagesOfficial              []string `yaml:"languages_official"`
	LanguagesSpoken                []string `yaml:"languages_spoken"`
	Nationality                    string   `yaml:"nationality"`
	AddressFormat                  string   `yaml:"address_format"`
	DissolvedOn                    string   `yaml:"dissolved_on"`
	EuMember                       string   `yaml:"eu_member"`
	AltCurrency                    string   `yaml:"alt_currency"`
	VatRates                       string   `yaml:"vat_rates"`
	PostalCode                     string   `yaml:"postal_code"`
	CurrencyCode                   string   `yaml:"currency_code"`
	StartOfWeek                    string   `yaml:"start_of_week"`
}

func (p *Country) TableName() string {
	//todo
	return "country"
}
