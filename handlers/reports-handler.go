package handlers

import (
	m "Atlassian/models"
	"fmt"
	"strconv"

	jira "Atlassian/jira"
)

func ReportIssues(issues []jira.Issue) {
	println("****************************************************")
	println("Unresolved issues for the SIAFE Alagoas (Brazil) project")
	for i := 0; i < len(issues); i++ {
		fmt.Println(issues[i].Key + " - " + issues[i].Fields.Summary)
	}
}
func ReportUsers(users []m.User) {
	println("****************************************************")
	println("Members of the SIAFE Group of the State of Alagoas (Brazil)")
	println("AccountID|", "Display Name|", "EmailAddress")
	for i, u := range users {
		fmt.Println(strconv.Itoa(i+1), u.AccountID, u.DisplayName, u.EmailAddress)
	}

}
