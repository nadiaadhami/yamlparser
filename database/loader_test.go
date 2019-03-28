package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func TestLoadCountries(t *testing.T) {
	d := "../database/data/"
	m := LoadCountries(d)
	assert.Equal(t, len(m), 250)
	assert.Equal(t, len(m["AD"].UnofficialNames), 3)
	assert.Equal(t, m["AD"].Name, "Andorra")
}

func TestParseSubdivisionsUS(t *testing.T) {
	us := "./data/subdivisions/US.yaml"
	usSubs := ParseSubdivision(us)
	assert.Equal(t, len(usSubs), 60)

	subdivisionAA := "AA"
	aa := usSubs[subdivisionAA]

	assert.Equal(t, aa.Name, "Armed Forces Americas")
	assert.Equal(t, len(aa.Translations), 1)
	assert.NotNil(t, aa.Geo)
	assert.Equal(t, aa.Geo.Latitude, float32(33.893497))
	assert.NotNil(t, aa.UnofficialNames)

	subdivisionAK := "AK"
	ak := usSubs[subdivisionAK]

	assert.Equal(t, ak.Name, "Alaska")
	assert.Equal(t, len(ak.Translations), 62)
	assert.NotNil(t, ak.Geo)
	assert.Equal(t, ak.Geo.Latitude, float32(64.20084))
	assert.NotNil(t, ak.UnofficialNames)
}

func TestParseSubdivisionsBE(t *testing.T) {
	f := "./data/subdivisions/BE.yaml"
	beSubs := ParseSubdivision(f)
	assert.Equal(t, len(beSubs), 13)
	subdivisionName := "BRU"
	s := beSubs[subdivisionName]

	assert.Equal(t,s.Name, "Brussels")
	assert.Equal(t,len(s.Translations), 61)
	assert.NotNil(t,s.Geo)
	assert.Equal(t, s.Geo.Latitude, float32(50.850338))
	assert.NotNil(t, s.UnofficialNames)

}

