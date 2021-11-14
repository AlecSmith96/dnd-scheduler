package main

import (
	"log"
	"os"

	"github.com/AlecSmith96/dnd-scheduler/adapters"
	"github.com/AlecSmith96/dnd-scheduler/entities"
	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("Unable to open config file: %v", err)
	}
	defer f.Close()

	var config entities.Config
	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Unable to decode config file: %v", err)
	}

	db, err := adapters.GetConn(&config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	adapters.TearDownDB(db)
	adapters.PopulateDB(db)

	adapters.Router()
}
