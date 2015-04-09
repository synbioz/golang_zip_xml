package extractor

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Car struct {
	XMLName      xml.Name `xml:"car"`
	Manufacturer string   `xml:"manufacturer"`
	Model        string   `xml:"model"`
	Year         int      `xml:"year"`
}

type Cars struct {
	XMLName xml.Name `xml:"cars"`
	Cars    []Car    `xml:"car"`
}

func ExtractCarsFromXml(fp string) ([]Car, error) {
	file, err := os.Open(fp)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var cars Cars

	xml.Unmarshal(content, &cars)

	return cars.Cars, nil
}
