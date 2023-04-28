package plain_handwritten

import (
	"bytes"
	"cg-plain/plain-sdk-golang/pkg/plain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"go.uber.org/zap"
)

var (
	plainUrl string = "https://core-api.uk.plain.com/graphql/v1"
)

type PlainClient struct {
	apiKey string
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, apiKey string) *PlainClient {
	return &PlainClient{
		apiKey: apiKey,
		logger: logger,
	}
}

type CustInput struct {
	Input plain.UpsertCustomerInput `json:"input,omitempty"`
}

type QueryOutput struct {
	Data string `json:"data,omitempty"`
}

type QueryOutputData struct {
	UpsertCustomer *plain.UpsertCustomerOutput `json:"upsert_customer,omitempty"`
}

func (c *PlainClient) UpsertCustomer(input plain.UpsertCustomerInput) (*plain.UpsertCustomerOutput, error) {
	b, err := os.ReadFile("./pkg/plain/graphql/upsertCustomer.graphql") // just pass the file name
	if err != nil {
		return nil, err
	}
	custInput := CustInput{
		Input: input,
	}
	marshalled, err := json.Marshal(&custInput)
	if err != nil {
		return nil, err
	}

	fullQuery := Query{
		Query:         string(b),
		Variables:     string(marshalled),
		OperationName: "upsertCustomer",
	}

	marshalledQuery, err := json.Marshal(&fullQuery)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, plainUrl, bytes.NewReader(marshalledQuery))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("output %s", string(resBody))
	outputStr := []byte("{\"data\":{\"upsertCustomer\":{\"result\":\"NOOP\",\"customer\":{\"id\":\"c_01GZ4AS090QYR8WDQ0CBG6Z5CS\",\"externalId\":null,\"shortName\":null,\"fullName\":\"hello world\",\"email\":{\"email\":\"hello@world.com\",\"isVerified\":true},\"status\":\"IDLE\"},\"error\":null}}}")
	output := QueryOutput{}
	json.Unmarshal(outputStr, &output)
	fmt.Printf("%v", output)
	//msg := plain.UpsertCustomerOutput{}
	// json.Unmarshal(output["data"], &msg)
	// fmt.Printf("msg %v", msg)
	// if output.Data.UpsertCustomer.Error != nil {
	// 	return nil, fmt.Errorf("Graphql error: %s", output.Data.UpsertCustomer.Error.Message)
	// }
	return nil, nil
}

func (c *PlainClient) CreateCustomer(email, fullName, shortName string) (string, error) {
	b, err := os.ReadFile("./pkg/plain/graphql/upsertCustomer.graphql") // just pass the file name
	if err != nil {
		return "", err
	}

	mutationInput := CreateCustomerMutationInput{
		Input: CreateCustomerInput{
			Identifier: Identifier{
				EmailAddress: email,
			},
			OnCreate: CreateCustomerData{
				FullName:  fullName,
				ShortName: shortName,
				Email: Email{
					Email: email,
				},
			},
			OnUpdate: UpdateCustomerData{
				FullName: UpdateObject{
					Value: fullName,
				},
				ShortName: UpdateObject{
					Value: shortName,
				},
				Email: Email{
					Email: email,
				},
			},
		},
	}
	marshalled, err := json.Marshal(&mutationInput)
	if err != nil {
		return "", err
	}

	fullQuery := Query{
		Query:         string(b),
		Variables:     string(marshalled),
		OperationName: "upsertCustomer",
	}

	marshalledQuery, err := json.Marshal(&fullQuery)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, plainUrl, bytes.NewReader(marshalledQuery))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	target := GraphqlResponse{}
	json.Unmarshal(resBody, &target)
	if target.Error != nil {
		return "", fmt.Errorf("Graphql error: %s", *target.Error)
	}
	c.logger.Debugf("Customer was %s, id: %s\n", target.Data.UpsertCustomer.Result, target.Data.UpsertCustomer.Customer.ID)
	return target.Data.UpsertCustomer.Customer.ID, nil
}

type TimelineData struct {
	Subject string
	Preview string
	Date    string
}

func (c *PlainClient) UpsertCustomTimelineEntry(input plain.UpsertCustomTimelineEntryInput) (*plain.UpsertCustomTimelineEntryOutput, error) {
	b, err := os.ReadFile("./pkg/plain/graphql/upsertCustomTimelineEntry.graphql") // just pass the file name
	if err != nil {
		return nil, err
	}

	marshalled, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}

	fullQuery := Query{
		Query:         string(b),
		Variables:     string(marshalled),
		OperationName: "upsertCustomTimelineEntry",
	}

	marshalledQuery, err := json.Marshal(&fullQuery)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, plainUrl, bytes.NewReader(marshalledQuery))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var respData struct {
		UpsertCustomTimelineEntry *plain.UpsertCustomTimelineEntryOutput
	}
	json.Unmarshal(resBody, &respData)
	if respData.UpsertCustomTimelineEntry.Error != nil {
		return nil, fmt.Errorf("Graphql error: %s", respData.UpsertCustomTimelineEntry.Error.Message)
	}
	return respData.UpsertCustomTimelineEntry, nil
}
