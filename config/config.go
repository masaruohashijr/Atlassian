package config

import (
	m "Atlassian/models"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var GmailConfig m.Config
var JiraConfig m.Config
var SDeskConfig m.Config

func InitConfig() {
	JiraConfig = ReadConfig(m.JIRA)
	GmailConfig = ReadConfig(m.GMAIL)
	SDeskConfig = ReadConfig(m.SDESK)
}

func ReadConfig(t m.ConfigType) m.Config {
	var configfile = "config/"
	if t == m.JIRA {
		configfile += "jira.cfg"
	} else if t == m.GMAIL {
		configfile += "gmail.cfg"
	} else if t == m.SDESK {
		configfile += "serviceDesk.cfg"
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
