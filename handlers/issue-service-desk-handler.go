package handlers

import (
	c "Atlassian/config"
	m "Atlassian/models"

	"gopkg.in/andygrunwald/go-jira.v1"
)

func RemoveIssues(jiraClient jira.Client) {
	last := 0
	opt := &jira.SearchOptions{
		MaxResults: 1000,
		StartAt:    last,
	}
	AdminAccountIDs := "557058:2664a93a-3fc4-4407-a4a2-2984447d2e81, 557058:11e222cb-58ae-4910-b369-0e8a66a6e34d, qm:15189404-2985-49fe-ac35-a2fa15eadbd2:5e0f5e12-ecbd-4386-b36d-ac533318002c"
	search := "project = \"AL - Service Desk\" AND creator in (" + AdminAccountIDs + ")"
	resultIssues, _, _ := jiraClient.Issue.Search(search, opt)
	for _, r := range resultIssues {
		jiraClient.Issue.Delete(r.ID)
	}
}
func CreateIssues(jiraClient jira.Client, issues []jira.Issue, usersMap map[string]m.User) {
	newIssue := new(jira.Issue)
	fields := new(jira.IssueFields)
	newIssue.Fields = fields
	newIssue.Fields.Assignee = new(jira.User)
	newIssue.Fields.Reporter = new(jira.User)
	for _, i := range issues {
		newIssue.Fields.Summary = i.Fields.Summary
		newIssue.Fields.Description = i.Fields.Description + "\n#-----" + "\nStatus: " + i.Fields.Status.Name
		println("Assignee accountid: " + i.Fields.Assignee.AccountID)
		println("Reporter accountid: " + i.Fields.Reporter.AccountID)
		println(i.Fields.Reporter.AccountID)
		newIssue.Fields.Project.Key = c.SDeskConfig.SDeskProjectKey
		newIssue.Fields.Type.Name = c.SDeskConfig.SDeskIssueTypeName
		//nova, _, err := jiraClient.Issue.Create(newIssue)
		/*if err != nil {
			println(err.Error())
		} else {
			for _, c := range i.Fields.Comments.Comments {
				c.Body = c.Body + "\n# --------------------"
				c.Body = c.Body + "\nAutor: " + usersMap[c.Author.AccountID].DisplayName
				c.Body = c.Body + "\nCriado em: " + c.Created
				jiraClient.Issue.AddComment(nova.Key, c)
			}
		}*/
	}
}
