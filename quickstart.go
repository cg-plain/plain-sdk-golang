package main

import (
	"cg-plain/plain-sdk-golang/pkg/plain"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	gqlclient "git.sr.ht/~emersion/gqlclient"
)

func main() {
	endpoint := "https://core-api.uk.plain.com/graphql/v1"
	token := os.Getenv("PLAIN_API_KEY")
	client := gqlclient.New(endpoint, &http.Client{
		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
				return nil, nil
			},
		},
	})
	custInput := plain.UpsertCustomerInput{
		Identifier: plain.UpsertCustomerIdentifierInput{
			EmailAddress: strPointer("hello@world.com"),
		},
	}
	cust, err := plain.UpsertCustomer(client, context.Background(), custInput)
	fmt.Printf("customer: %v", cust)
	if err != nil {
		panic(err)
	}
	if cust.Error != nil {
		panic(cust.Error)
	}
	fmt.Printf(cust.Customer.FullName)

	timelineEntry, err := plain.UpsertCustomTimelineEntry(client, context.Background(), plain.UpsertCustomTimelineEntryInput{
		CustomerId: cust.Customer.Id,
		Components: []plain.ComponentInput{
			{
				ComponentText: &plain.ComponentTextInput{
					TextSize: textSizePointer(plain.ComponentTextSizeM),
					Text:     "hello, world",
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf(timelineEntry.TimelineEntry.Id)

	issue, err := plain.CreateIssue(client, context.Background(), plain.CreateIssueInput{})
	if err != nil {
		panic(err)
	}
	fmt.Printf(issue.Issue.Id)

	// or, roll your own:
	f, err := os.Open("./graphql/upsertCustomer.graphql")
	if err != nil {
		panic(err)
	}

	queryBytes := make([]byte, 0)
	_, err = f.Read(queryBytes)
	if err != nil {
		panic(err)
	}
	op := gqlclient.NewOperation(string(queryBytes))
	op.Var("input", custInput)
	var respData struct {
		UpsertCustomer *plain.UpsertCustomerOutput
	}
	err = client.Execute(context.Background(), op, &respData)
}

func strPointer(input string) *string {
	return &input
}

func textSizePointer(input plain.ComponentTextSize) *plain.ComponentTextSize {
	return &input
}
