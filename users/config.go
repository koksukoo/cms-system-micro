package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database confDB
	Hashids  confHash
	Server   confServer
	Redis    confRedis
}

type confDB struct {
	Server   string
	User     string
	Password string
	Database string
	Port     string
}

type confHash struct {
	Salt string
}

type confServer struct {
	Port string
}

type confRedis struct {
	Server string
	Port   string
}

var conf Config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatal(err)
	}
}
