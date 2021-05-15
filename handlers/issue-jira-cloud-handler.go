package handlers

import (
	jira "Atlassian/jiracloud"
)

func UpdateTitleAndCustomField(jiraClient *jira.Client, allIssues []jira.Issue) {
	for _, i := range allIssues {
		j := &jira.Issue{
			Key: i.Key,
			Fields: &jira.IssueFields{
				// Added the line 141 in "issue.go" file from "Atlassian/jiracloud"
				// OriginKey string `json:"customfield_13231,omitempty" structs:"customfield_13231,omitempty"`
				OriginKey: i.Key,
			},
		}
		_, _, err := jiraClient.Issue.Update(j)
		println(i.Key, "updated with success.")
		if err != nil {
			println(err.Error())
		}
	}
}
