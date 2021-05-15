package main

import (
	c "Atlassian/config"
	hd "Atlassian/handlers"
	m "Atlassian/models"
	"fmt"
	"strings"

	jira "Atlassian/jira"
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
	allIssues, _ := hd.GetIssuesByJql(*jiraClient, jql)
	println(len(allIssues))
	usersMap := make(map[string]m.User)
	groups := strings.Split(c.JiraConfig.JiraProjectGroups, ",")
	var allUsers []m.User
	for _, g := range groups {
		jiraUsers, _ := hd.GetMembersFromGroup(jiraClient, g)
		if false {
			var user m.User
			var users []m.User
			page := hd.StartDriver()
			for _, u := range *jiraUsers {
				emailAddress := hd.ScrapEmailAddress(page, u.AccountID)
				user.AccountID = u.AccountID
				user.DisplayName = u.DisplayName
				user.EmailAddress = emailAddress
				users = append(users, user)
				usersMap[u.AccountID] = user
				println(usersMap[u.AccountID].String())
				allUsers = append(allUsers, user)
			}
			hd.StopDriver()
			hd.ReportUsers(users)
		}
	}
	hd.ReportIssues(allIssues)
	hd.SaveUsers(allUsers)
	// as rollback test -> clear all issues of the destination project.
	hd.RemoveIssues(*jiraClient)
	// create issues if not already exists another in the destination project with the same title / subject.
	// Comments: "Author, CreatedAt"
	// Issue: "Reporter, Assignee, Created, Status (From-To), Link"
	// hd.CreateIssues(*jiraClient, issues, usersMap)
	// create organization if not already exists with the same name.
	// add users to organization if not already exists another in the destination organization.
	// hd.UpdateCustomFieldOriginKey(jiraClient, allIssues)
	hd.AddUserToOrganization(*jiraClient, allUsers)
}
