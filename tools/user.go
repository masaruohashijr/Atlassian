package tools

import (
	"fmt"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func GetUserByAccountId(jiraClient *jira.Client, accountId string) (*jira.User, error) {
	// "557058:2664a93a-3fc4-4407-a4a2-2984447d2e81"
	apiEndpoint := fmt.Sprintf("/rest/api/2/user?accountId=%s", accountId)
	req, err := jiraClient.NewRequest("GET", apiEndpoint, nil)
	user := new(jira.User)
	_, err = jiraClient.Do(req, user)
	if err != nil {
		print(err.Error())
	}
	//print(user.DisplayName + "\n\n")
	return user, err
}
