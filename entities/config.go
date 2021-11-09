package entities

import (

)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname string `yaml:"dbname"`
	} `yaml:"database"`

}