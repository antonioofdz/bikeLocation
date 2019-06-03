package handlers

import (
	"errors"
	"fmt"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
)

func GetAddress(lat float64, lon float64) (string, error) {
	return getAddress(google.Geocoder("AIzaSyB09lpKFkFLeWZjin_csgssU4FywZ0z_UA"), lat, lon)
}

func getAddress(geocoder geo.Geocoder, lat float64, lon float64) (string, error) {
	address, _ := geocoder.ReverseGeocode(lat, lon)
	if address != nil {
		fmt.Printf("Address of (%.6f,%.6f) is %s\n", lat, lon, address.FormattedAddress)
		fmt.Printf("Detailed address: %#v\n", address)

		return address.FormattedAddress, nil
	}
	fmt.Println("got <nil> address")
	return "", errors.New("got <nil> address")
}
