package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
    SERVERS_CONFIG_PATH = "./servers.json"
    SCRIPTS_CONFIG_PATH = "./scripts.json"
    TOKEN_ENV_VAR = "DISCORD_TOKEN"
)

type (
    Server struct {
	Name string `json:"name"`
	IpAddr string `json:"ipAddr"`
	User string `json:"user"`
	KeyPath string `json:"keyPath"`
    }

    Script struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ScriptPath string `json:"scriptPath"`
	Servers []string `json:"servers"`
    }

    Configuration struct {
	Servers []Server
	Scripts []Script
	Token string
    }
)

func InitConfig() (Configuration, error) {
    token, ok := os.LookupEnv(TOKEN_ENV_VAR)
    if !ok {
	return Configuration{}, errors.New(fmt.Sprint("No environment variable found for", TOKEN_ENV_VAR))
    }

    data, err := ioutil.ReadFile(SERVERS_CONFIG_PATH)
    if err != nil {
	return Configuration{}, errors.New(fmt.Sprint("Error reading", SERVERS_CONFIG_PATH, err))
    }
    var servers []Server
    err = json.Unmarshal(data, &servers)
    if err != nil {
	return Configuration{}, errors.New(fmt.Sprint("Error unmarshalling json from", SERVERS_CONFIG_PATH, err))
    }

    data, err = ioutil.ReadFile(SCRIPTS_CONFIG_PATH)
    if err != nil {
	return Configuration{}, errors.New(fmt.Sprint("Error reading", SCRIPTS_CONFIG_PATH, err))
    }
    var scripts []Script
    err = json.Unmarshal(data, &scripts)
    if err != nil {
	return Configuration{}, errors.New(fmt.Sprint("Error unmarshalling json from", SERVERS_CONFIG_PATH, err))
    }

    return Configuration{
    	Servers: servers,
    	Scripts: scripts,
    	Token:   token,
    }, nil
}
