package handlers

import (
	c "Atlassian/config"

	"gopkg.in/andygrunwald/go-jira.v1"
)

func RemoveIssues(jiraClient jira.Client, issues []jira.Issue) {
	last := 0
	opt := &jira.SearchOptions{
		MaxResults: 100,
		StartAt:    last,
	}
	search := "project = \"Demo service desk\" AND created > startOfDay(0)"
	resultIssues, _, _ := jiraClient.Issue.Search(search, opt)
	for _, r := range resultIssues {
		jiraClient.Issue.Delete(r.ID)
	}
}
func CreateIssues(jiraClient jira.Client, issues []jira.Issue) {
	newIssue := new(jira.Issue)
	fields := new(jira.IssueFields)
	newIssue.Fields = fields
	for _, i := range issues {
		newIssue.Fields.Summary = i.Fields.Summary
		//println(i.Fields.Summary)
		//println(i.Fields.Description)
		println(i.Fields.Reporter.AccountID)
		println(i.Fields.Reporter.DisplayName)
		//println(i.Fields.Created)
		newIssue.Fields.Description = i.Fields.Description
		//newIssue.Fields.Comments = i.Fields.Comments
		newIssue.Fields.Project.Key = c.SDeskConfig.SDeskProjectKey
		newIssue.Fields.Type.Name = c.SDeskConfig.SDeskIssueTypeName
		newIssue, _, _ := jiraClient.Issue.Create(newIssue)
	}
}
