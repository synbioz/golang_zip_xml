package main

import (
	"fmt"
	"github.com/synbioz/golang_zip_xml/extractor"
	"github.com/synbioz/golang_zip_xml/unzip"
	"io/ioutil"
	"log"
	"path"
)

func main() {
	log.Println("unzip file ...")

	err := unzip.Unzip("cars.zip")

	if err != nil {
		return
	}

	log.Println("begin parsing...")

	files, _ := ioutil.ReadDir(unzip.TMP_DIRECTORY)

	for _, f := range files {
		log.Printf("extract from %s....", f.Name())

		path := path.Join(unzip.TMP_DIRECTORY, f.Name())

		if cars, err := extractor.ExtractCarsFromXml(path); err == nil {
			log.Printf("show cars from %s", f.Name())
			ShowEachCar(cars)
		}

	}
}

func ShowEachCar(cars []extractor.Car) {
	for _, car := range cars {
		fmt.Printf("Voiture de marque: %s", car.Manufacturer)
		fmt.Printf(" de modèle: %s de l'année %d", car.Model, car.Year)
		fmt.Println()
	}
}
