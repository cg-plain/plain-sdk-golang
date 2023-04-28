package main

import (
	"cg-plain/plain-sdk-golang/pkg/plain"
	"cg-plain/plain-sdk-golang/pkg/plain/plain_handwritten"
	"fmt"
	"os"
	"encoding/json"

	"go.uber.org/zap"
)

func strPointer(input string) *string {
	return &input
}

func textSizePointer(input plain.ComponentTextSize) *plain.ComponentTextSize {
	return &input
}

func main() {
	token := os.Getenv("PLAIN_API_KEY")
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
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	sugar := logger.Sugar()
	client := plain_handwritten.New(sugar, token)
	out, err := client.UpsertCustomer(custInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf(out.Customer.Id)

	timelineEntry, err := client.UpsertCustomTimelineEntry(plain.UpsertCustomTimelineEntryInput{
		CustomerId: out.Customer.Id,
		Title:      "test",
		Components: []plain.ComponentInput{
			{
				ComponentText: &plain.ComponentTextInput{
					TextSize: textSizePointer(plain.ComponentTextSizeM),
					Text:     "hello, world",
				},
			},
		},
	})
	fmt.Printf("timelineEntry: %v", timelineEntry)
	if err != nil {
		panic(err)
	}
	if timelineEntry.Error != nil {
		panic(timelineEntry.Error.Message)
	}
	fmt.Printf(timelineEntry.TimelineEntry.Id)

	issueType, err := client.CreateIssueType(plain.CreateIssueTypeInput{
		PublicName: "publicName",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf(issueType.IssueType.Id)
	issue, err := client.CreateIssue(plain.CreateIssueInput{
		CustomerId: out.Customer.Id,
		IssueTypeId: issueType.IssueType.Id,
	})
	if err != nil {
		panic(err)
	}
	if issue.Error != nil {
		panic(issue.Error.Message)
	}
	fmt.Printf(issue.Issue.Id)

	// or, roll your own:
	f, err := os.Open("./pkg/plain/graphql/upsertCustomer.graphql")
	if err != nil {
		panic(err)
	}

	queryBytes := make([]byte, 0)
	_, err = f.Read(queryBytes)
	if err != nil {
		panic(err)
	}
	
	type upsertIn struct {
		input plain.UpsertCustomerInput `json:"input,omitempty"`
	}
	wrappedInput := upsertIn{
		input: custInput,
	}

	inputBytes, err := json.Marshal(&wrappedInput)
	if err != nil {
		panic(err)
	}

	outputBytes, err := client.Query("upsertCustomer", string(queryBytes), string(inputBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("output: %s", outputBytes)
}
