package utils

import (
	"-invoice_manager/internal/models"
	"encoding/json"
	"log"
	"os"
)

func DecodeConf() (conf models.Config) {

	configfile, err := os.Open("../../config.json")
	if err != nil {
		log.Println(err)
	}
	defer configfile.Close()

	decoder := json.NewDecoder(configfile)
	if err := decoder.Decode(&conf); err != nil {
		log.Println(err)
	}
	return conf
}
