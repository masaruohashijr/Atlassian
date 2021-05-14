package main

import (
	c "Atlassian/config"
	hd "Atlassian/handlers"
	"fmt"

	"gopkg.in/andygrunwald/go-jira.v1"
)

func main() {

	c.InitConfig()

	tp := jira.BasicAuthTransport{
		Username: c.JiraConfig.JiraUsername,
		Password: c.JiraConfig.JiraApiToken,
	}

	jiraClient, err := jira.NewClient(tp.Client(), c.JiraConfig.JiraAddress)
	if err != nil {
		print(err.Error())
	}

	jql := c.JiraConfig.JiraUseCaseJQL
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)
	issues, _ := hd.GetIssuesByJql(*jiraClient, jql)

	jiraUsers, _ := hd.GetMembersFromGroup(jiraClient, c.JiraConfig.JiraUserGroup)
	var user m.User
	var users []m.User
	page := hd.StartDriver()
	for _, u := range *jiraUsers {
		emailAddress := hd.ScrapEmailAddress(page, u.AccountID)
		user.AccountID = u.AccountID
		user.DisplayName = u.DisplayName
		user.EmailAddress = emailAddress
		users = append(users, user)
	}
	hd.StopDriver()
	hd.ReportIssues(issues)
	hd.ReportUsers(users)
	// as rollback test -> clear all issues of the destination project.
	hd.RemoveIssues(*jiraClient, issues)
	// create issues if not already exists another in the destination project with the same title / subject.
	hd.CreateIssues(*jiraClient, issues)
	// create organization if not already exists with the same name.
	// add users to organization if not already exists another in the destination organization.
}
