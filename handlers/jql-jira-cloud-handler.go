package handlers

import (
	"gopkg.in/andygrunwald/go-jira.v1"
)

func GetIssuesByJql(client jira.Client, searchString string) ([]jira.Issue, error) {
	last := 0
	var issues []jira.Issue
	opt := &jira.SearchOptions{
		MaxResults: 1000, // Max results can go up to 1000
		StartAt:    last,
		Expand:     "renderedBody",
		Fields: []string{
			"summary", "description", "reporter",
			"status", "assignee", "author",
			"created", "Creator", "comment",
			"worklog"},
	}

	chunk, resp, err := client.Issue.Search(searchString, opt)
	if err != nil {
		return nil, err
	}

	total := resp.Total
	if issues == nil {
		issues = make([]jira.Issue, 0, total)
	}
	issues = append(issues, chunk...)
	last = resp.StartAt + len(chunk)
	if last >= total {
		return issues, nil
	}
	return issues, nil
}
