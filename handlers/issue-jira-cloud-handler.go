package handlers

import (
	jira "Atlassian/jira"
)

func UpdateCustomFieldOriginKey(jiraClient *jira.Client, allIssues []jira.Issue) {
	for _, i := range allIssues {
		j := &jira.Issue{
			Key: i.Key,
			Fields: &jira.IssueFields{
				// Added the line 141 in "issue.go" file from "Atlassian/jira"
				// OriginKey string `json:"customfield_13231,omitempty" structs:"customfield_13231,omitempty"`
				OriginKey: i.Key,
			},
		}

		opts := &jira.UpdateQueryOptions{NotifyUsers: false}
		_, _, err := jiraClient.Issue.UpdateWithOptions(j, opts)
		println(i.Key, "updated with success.")
		if err != nil {
			println(err.Error())
		}
	}
}
