package handlers

import (
	jira "Atlassian/jiracloud"
	m "Atlassian/models"
)

func AddUser() {

}

func AddUserToOrganization(jiraClient jira.Client, allUsers []m.User) {
	/*	Authentication
		The Jira Service Management REST API uses the same authentication methods as Jira Cloud.
	*/
}
