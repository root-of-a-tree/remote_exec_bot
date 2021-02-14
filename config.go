package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
    SERVERS_CONFIG_PATH = "./servers.json"
    SCRIPTS_CONFIG_PATH = "./scripts.json"
    TOKEN_ENV_VAR = "DISCORD_TOKEN"
)

var (
    Config Configuration
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
	Servers []Server `json:"servers"`
	Scripts []Script `json:"scripts"`
	Token string `json:"token"`
    }
)

func init() {
    token, ok := os.LookupEnv(TOKEN_ENV_VAR)
    if !ok {
	fmt.Println("No environment variable found for", TOKEN_ENV_VAR)
	return
    }

    data, err := ioutil.ReadFile(SERVERS_CONFIG_PATH)
    if err != nil {
	fmt.Println("Error reading", SERVERS_CONFIG_PATH, err)
	return
    }
    var servers []Server
    err = json.Unmarshal(data, &servers)
    if err != nil {
	fmt.Println("Error unmarshalling json from", SERVERS_CONFIG_PATH, err)
	return
    }

    data, err = ioutil.ReadFile(SCRIPTS_CONFIG_PATH)
    if err != nil {
	fmt.Println("Error reading", SCRIPTS_CONFIG_PATH, err)
	return
    }
    var scripts []Script
    err = json.Unmarshal(data, &scripts)
    if err != nil {
	fmt.Println("Error unmarshalling json from", SERVERS_CONFIG_PATH, err)
	return
    }

    Config = Configuration{
    	Servers: servers,
    	Scripts: scripts,
    	Token:   token,
    }
    return
}
