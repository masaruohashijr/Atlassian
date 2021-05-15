package handlers

import (
	"Atlassian/models"
	"fmt"

	jira "Atlassian/jira"
)

func GetMembersFromGroup(jiraClient *jira.Client, groupname string) (*[]jira.User, error) {
	if !ExistsSerialData(groupname) {
		apiEndpoint := fmt.Sprintf("/rest/api/3/group/member?groupname=%s", groupname)
		req, err := jiraClient.NewRequest("GET", apiEndpoint, nil)
		v := new(models.MembersFromGroup)
		_, err = jiraClient.Do(req, v)
		if err != nil {
			panic(err)
		}
		return &v.Values, err
	}
	return nil, nil
}
