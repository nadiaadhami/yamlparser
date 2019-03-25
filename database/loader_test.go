package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}
func TestLoader(t *testing.T) {
	d := "../database/data/"
	m := LoadCountries(d)
	assert.Equal(t, len(m), 250)
}
