package main

import (
	"cg-plain/plain-sdk-golang/pkg/plain"
	"cg-plain/plain-sdk-golang/pkg/plain/plain_handwritten"
	"fmt"
	"os"

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
	// custInput := plain.UpsertCustomerInput{
	// 	Identifier: plain.UpsertCustomerIdentifierInput{
	// 		EmailAddress: strPointer("hello@world.com"),
	// 	},
	// 	OnCreate: plain.UpsertCustomerOnCreateInput{
	// 		FullName: "hello world",
	// 		Email: plain.EmailAddressInput{
	// 			Email:      "hello@world.com",
	// 			IsVerified: true,
	// 		},
	// 	},
	// 	OnUpdate: plain.UpsertCustomerOnUpdateInput{
	// 		FullName: &plain.StringInput{
	// 			Value: "hello world",
	// 		},
	// 		Email: &plain.EmailAddressInput{
	// 			Email:      "hello@world.com",
	// 			IsVerified: true,
	// 		},
	// 	},
	// }
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	sugar := logger.Sugar()
	client := plain_handwritten.New(sugar, token)
	out, err := client.CreateCustomer("hello@helloworld.com", "hello", "world")
	if err != nil {
		panic(err)
	}
	fmt.Printf(out)
	// output, err := client.UpsertCustomer(custInput)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf(output.Customer.FullName)

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

	// issue, err := plain.CreateIssue(client, context.Background(), plain.CreateIssueInput{})
	// if err != nil {
	// 	panic(err)
	// }
	// if issue.Error != nil {
	// 	panic(issue.Error.Message)
	// }
	// fmt.Printf(issue.Issue.Id)

	// or, roll your own:
	// f, err := os.Open("./graphql/upsertCustomer.graphql")
	// if err != nil {
	// 	panic(err)
	// }

	// queryBytes := make([]byte, 0)
	// _, err = f.Read(queryBytes)
	// if err != nil {
	// 	panic(err)
	// }
	// op := gqlclient.NewOperation(string(queryBytes))
	// op.Var("input", custInput)
	// var respData struct {
	// 	UpsertCustomer *plain.UpsertCustomerOutput
	// }
	// err = client.Execute(context.Background(), op, &respData)
}
