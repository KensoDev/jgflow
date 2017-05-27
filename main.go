package jgflow

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, err := jira.NewClient(nil, "https://your.jira-instance.com/")
	if err != nil {
		panic(err)
	}

	res, err := jiraClient.Authentication.AcquireSessionCookie("username", "password")
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	issue, _, err := jiraClient.Issue.Get("SYS-5156", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}
