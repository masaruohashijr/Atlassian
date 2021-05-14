package handlers

import (
	m "Atlassian/models"
	"strconv"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

func ReportIssues(issues []jira.Issue) {
	println("****************************************************")
	println("Unresolved issues for the SIAFE Alagoas (Brazil) project")
	for i := 0; i < len(issues); i++ {
		println(issues[i].Key + " - " + issues[i].Fields.Summary)
	}
}
func ReportUsers(users []m.User) {
	println("****************************************************")
	println("Members of the SIAFE Group of the State of Alagoas (Brazil)")
	println("AccountID|", "Display Name|", "EmailAddress")
	for i, u := range users {
		println(strconv.Itoa(i+1), u.AccountID, u.DisplayName, u.EmailAddress)
	}

}
