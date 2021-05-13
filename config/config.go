package config

import (
	m "Atlassian/models"
	"Atlassian/vendor/github.com/BurntSushi/toml"
	"log"
	"os"
)

func ReadConfig(t m.ConfigType) m.Config {
	var configfile = "config/"
	if t == m.JIRA {
		configfile += "jira.cfg"
	} else if t == m.GMAIL {
		configfile += "gmail.cfg"
	}
	log.Println(configfile)
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("File configuration "+configfile+" missing: ", configfile)
	}
	var config m.Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
