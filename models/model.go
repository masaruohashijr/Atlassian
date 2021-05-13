package models

import (
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

type MembersFromGroup struct {
	Values     []jira.User `json:"values" structs:"values"`
	StartAt    int         `json:"startAt" structs:"startAt"`
	MaxResults int         `json:"maxResults" structs:"maxResults"`
	Total      int         `json:"total" structs:"total"`
}

type ConfigType int

const (
	JIRA ConfigType = iota
	GMAIL
)

type Config struct {
	JiraUsername    string
	JiraApiToken    string
	JiraAddress     string
	JiraUserGroup   string
	JiraProfilePage string
	JiraNextButton  string
	JiraUseCaseJQL  string
	DomainParts     string
	AdminEmail      string
	AdminPassword   string
}
