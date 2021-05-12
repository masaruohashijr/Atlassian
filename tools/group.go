package tools

import (
	"Atlassian/models"
	"fmt"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func GetMembersFromGroup(jiraClient *jira.Client, groupname string) (*[]jira.User, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/3/group/member?groupname=%s", groupname)
	req, err := jiraClient.NewRequest("GET", apiEndpoint, nil)
	v := new(models.MembersFromGroup)
	_, err = jiraClient.Do(req, v)
	if err != nil {
		panic(err)
	}
	return &v.Values, err
}
