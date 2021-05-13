package main

import (
	"Atlassian/config"
	"Atlassian/models"
	"Atlassian/tools"
	"fmt"
	"strconv"

	"gopkg.in/andygrunwald/go-jira.v1"
)

func main() {

	jiraConfig := config.ReadConfig(models.JIRA)

	tp := jira.BasicAuthTransport{
		Username: jiraConfig.JiraUsername,
		Password: jiraConfig.JiraApiToken,
	}

	jiraClient, err := jira.NewClient(tp.Client(), jiraConfig.JiraAddress)
	if err != nil {
		print(err.Error())
	}

	print("****************************************************\n")
	print("Unresolved issues for the SIAFE Alagoas project\n")
	jql := jiraConfig.JiraUseCaseJQL
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)
	issues, _ := tools.GetIssuesByJql(*jiraClient, jql)
	for i := 0; i < len(issues); i++ {
		print(issues[i].Key + " - " + issues[i].Fields.Summary + "\n")
	}

	print("****************************************************\n")
	print("Members of the SIAFE Group of the State of Alagoas (Brazil)\n")
	users, _ := tools.GetMembersFromGroup(jiraClient, jiraConfig.JiraUserGroup)
	for i, u := range *users {
		println(strconv.Itoa(i+1), u.DisplayName, u.AccountID)
		emailAddress := tools.ScrapEmailAddress(u.AccountID)
		println("***************************")
		println(emailAddress)
		println("***************************")
	}
}
