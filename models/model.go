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
