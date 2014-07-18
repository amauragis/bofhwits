package main

import (
    "gopkg.in/yaml.v1"
//     "fmt"
    "log"
    "io/ioutil"
)

type Config struct {
    Address string
    Username string
    Nick string
    Channel string
}

func LoadConfig() Config{
    
    
    source, err := ioutil.ReadFile("config/bofhwits.yaml")
     if err != nil {
        log.Fatal(err)
    }
    
    var configs Config
    
    err = yaml.Unmarshal(source, &configs)
    if err != nil {
        log.Fatal(err)
    }
    
    return configs
    
//     fmt.Printf("address: %#v\n", configs.Address)
//     fmt.Printf("user: %#v\n", configs.Username)
//     fmt.Printf("nick: %#v\n", configs.Nick)
//     fmt.Printf("chan: %#v\n", configs.Channel)
}

