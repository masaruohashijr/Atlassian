package handlers

import (
	m "Atlassian/models"
	"fmt"

	jira "Atlassian/jiracloud"
)

func GetUserByAccountId(jiraClient *jira.Client, accountId string) (*jira.User, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/user?accountId=%s", accountId)
	req, err := jiraClient.NewRequest("GET", apiEndpoint, nil)
	user := new(jira.User)
	_, err = jiraClient.Do(req, user)
	if err != nil {
		print(err.Error())
	}
	return user, err
}

func SaveUsers(allUsers []m.User) {

}
