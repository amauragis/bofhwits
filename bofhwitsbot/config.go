package bofhwitsbot

import (
    "gopkg.in/yaml.v1"
    "log"
    "io/ioutil"
    //     "fmt"
)

type Config struct {
    Address string
    Username string
    Nick string
    Channel string
    Twitter struct {
        Appapi string
        Appsecret string
        Accountapi string
        Accountsecret string
    }
}

// populate a config struct from a yaml file.
func LoadConfig(path string) Config{    
    source, err := ioutil.ReadFile(path)
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

