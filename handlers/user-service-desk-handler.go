package handlers

import (
	c "Atlassian/config"
	jira "Atlassian/jira"
	m "Atlassian/models"
	"strconv"
)

func AddUser() {

}

func AddUserToOrganization(jiraClient jira.Client, allUsers []m.User) {
	/*	Authentication
		The Jira Service Management REST API uses the same authentication methods as Jira Cloud.
	*/
	organizationID, _ := strconv.Atoi(c.JiraConfig.JiraOrganizationId)
	var organizationUsers jira.OrganizationUsersDTO
	// var accountIDs []string
	/*for i, u := range allUsers {
	accountIDs[i] = u.AccountID
	}*/
	accountIDs := []string{"557058:11e222cb-58ae-4910-b369-0e8a66a6e34d"}
	organizationUsers.AccountIds = accountIDs
	_, err := jiraClient.Organization.AddUsers(organizationID, organizationUsers)
	if err == nil {
		println("**********************")
		println("User added to group \"Siafe-AL Estado\" with success.")
		println("**********************")

	}
}
