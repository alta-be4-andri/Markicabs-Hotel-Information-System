package plugins

import (
	"os"

	"github.com/kelvins/geocoder"
)

func Geocode(city string) (float64, float64, error) {
	api_key := os.Getenv("API_KEY")
	geocoder.ApiKey = api_key

	var lng, lat float64
	var address geocoder.Address
	address.City = city

	location, err := geocoder.Geocoding(address)
	if err != nil {
		return 0, 0, err
	} else {
		lng = location.Latitude
		lat = location.Longitude
	}
	return lng, lat, nil
}
