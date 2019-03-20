package main

import (
	"time"
)

type Country struct {
	ID int64 `yaml:"column(id);pk;auto"`

	Number                         string `yaml:"column(number);size(255)"`
	Alpha2                         string `yaml:"column(alpha2);size(2)"`
	Alpha3                         string `yaml:"column(alpha3);size(3)"`
	Currency                       string `yaml:"column(currency);size(255)"`
	Name                           string `yaml:"column(name);size(255)"`
	UnofficialNames                []string `yaml:"column(unofficial_names)"`
	Continent                      string `yaml:"column(continent);size(255)"`
	Region                         string `yaml:"column(region);size(255)"`
	Subregion                      string `yaml:"column(subregion);size(255)"`
	//Geo                            string `yaml:"column(geo);size(255)"`
	WorldRegion                    string `yaml:"column(world_region);size(255)"`
	CountryCode                    string `yaml:"column(country_code);size(255)"`
	//NationalDestinationCodeLengths string `yaml:"column(national_destination_code_lengths);size(255)"`
	//NationalNumberLengths          string `yaml:"column(national_number_lengths);size(255)"`
	InternationalPrefix            string `yaml:"column(international_prefix);size(3)"`
	NationalPrefix                 string `yaml:"column(national_prefix);size(255)"`
	Ioc                            string `yaml:"column(ioc);size(3)"`
	Gec                            string `yaml:"column(gec);size(255)"`
	UnLocode                       string `yaml:"column(un_locode);size(255)"`
	LanguagesOfficial              string `yaml:"column(languages_official);size(255)"`
	LanguagesSpoken                string `yaml:"column(languages_spoken);size(255)"`
	Nationality                    string `yaml:"column(nationality);size(255)"`
	AddressFormat                  string `yaml:"column(address_format);size(255)"`
	DissolvedOn                    string `yaml:"column(dissolved_on);size(255)"`
	EuMember                       string `yaml:"column(eu_member);size(255)"`
	AltCurrency                    string `yaml:"column(alt_currency);size(255)"`
	VatRates                       string `yaml:"column(vat_rates);size(255)"`
	PostalCode                     string `yaml:"column(postal_code);size(255)"`
	CurrencyCode                   string `yaml:"column(currency_code);size(255)"`

	InsertedAt time.Time `yaml:"auto_now_add;type(datetime);column(inserted_at)"`
	UpdatedAt  time.Time `yaml:"auto_now;type(datetime);column(updated_at)"`
}

func (p *Country) TableName() string {
	//todo
	return "country"
}
