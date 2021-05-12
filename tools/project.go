package tools

import (
	"fmt"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func GetProjects(jiraClient *jira.Client, accountId string) (*[]jira.Project, error) {
	req, _ := jiraClient.NewRequest("GET", "rest/api/2/project", nil)
	projects := new([]jira.Project)
	_, err := jiraClient.Do(req, projects)
	if err != nil {
		panic(err)
	}
	for _, project := range *projects {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}
	return projects, err
}
