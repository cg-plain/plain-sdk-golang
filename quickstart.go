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

func other() {
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
		OnCreate: plain.UpsertCustomerOnCreateInput{
			FullName: "hello world",
			Email: plain.EmailAddressInput{
				Email:      "hello@world.com",
				IsVerified: true,
			},
		},
		OnUpdate: plain.UpsertCustomerOnUpdateInput{
			FullName: &plain.StringInput{
				Value: "hello world",
			},
			Email: &plain.EmailAddressInput{
				Email:      "hello@world.com",
				IsVerified: true,
			},
		},
	}
	cust, err := plain.UpsertCustomer(client, context.Background(), custInput)
	if err != nil {
		panic(err)
	}
	if cust.Error != nil {
		panic(cust.Error.Message)
	}
	fmt.Printf(cust.Customer.FullName)

	// timelineEntry, err := plain.UpsertCustomTimelineEntry(client, context.Background(), plain.UpsertCustomTimelineEntryInput{
	// 	CustomerId: cust.Customer.Id,
	// 	Title:      "test",
	// 	Components: []plain.ComponentInput{
	// 		{
	// 			ComponentText: &plain.ComponentTextInput{
	// 				TextSize: textSizePointer(plain.ComponentTextSizeM),
	// 				Text:     "hello, world",
	// 			},
	// 		},
	// 	},
	// })
	// fmt.Printf("timelineEntry: %v", timelineEntry)
	// if err != nil {
	// 	panic(err)
	// }
	// if timelineEntry.Error != nil {
	// 	panic(timelineEntry.Error.Message)
	// }
	//fmt.Printf(timelineEntry.TimelineEntry.Id)

	issue, err := plain.CreateIssue(client, context.Background(), plain.CreateIssueInput{})
	if err != nil {
		panic(err)
	}
	if issue.Error != nil {
		panic(issue.Error.Message)
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
