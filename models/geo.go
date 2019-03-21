package models

type Geo struct {
	Latitude     float32 `yaml:"latitude"`
	LatitudeDec  string  `yaml:"latitude_dec"`
	Longitude    float32 `yaml:"longitude"`
	LongitudeDec string  `yaml:"longitude_dec"`
	MaxLatitude  float32 `yaml:"max_latitude"`
	MaxLongitude float32 `yaml:"max_longitude"`
	MinLatitude  float32 `yaml:"min_latitude"`
	MinLongitude float32 `yaml:"min_longitude"`
	Bounds       Bounds  `yaml:"bounds"`
}

type LatLon struct {
	Lat float32 `yaml:"lat"`
	Lng float32 `yaml:"lng"`
}
type Bounds struct {
	Northeast LatLon `yaml:"northeast"`
	Southwest LatLon `yaml:"southwest"`
}
