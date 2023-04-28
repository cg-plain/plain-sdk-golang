package plain

import (
	"cg-plain/plain-sdk-golang/pkg/plain/graphql"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"


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

func (c *PlainClient) Query(operation, query, variables string) ([]byte, error) {
	fullQuery := Query{
		Query:         query,
		Variables:     variables,
		OperationName: operation,
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

	return io.ReadAll(res.Body)
}

type in struct {
	Input UpsertCustomerInput `json:"input,omitempty"`
}

func (c *PlainClient) UpsertCustomer(input UpsertCustomerInput) (*UpsertCustomerOutput, error) {
	mutationInput := in{
		Input: input,
	}
	marshalled, err := json.Marshal(&mutationInput)
	if err != nil {
		return nil, err
	}

	body, err := c.Query("upsertCustomer", graphql.UpsertCustomer, string(marshalled))

	target := GraphqlResponse{}
	json.Unmarshal(body, &target)
	if target.Data.UpsertCustomer.Error != nil {
		return nil, fmt.Errorf("Graphql error: %s", *target.Data.UpsertCustomer.Error)
	}
	return target.Data.UpsertCustomer, nil
}

type timelineIn struct {
	Input UpsertCustomTimelineEntryInput `json:"input,omitempty"`
}
	
func (c *PlainClient) UpsertCustomTimelineEntry(input UpsertCustomTimelineEntryInput) (*UpsertCustomTimelineEntryOutput, error) {
	mutationInput := timelineIn{
		Input: input,
	}
	marshalled, err := json.Marshal(&mutationInput)
	if err != nil {
		return nil, err
	}

	body, err := c.Query("upsertCustomTimelineEntry", graphql.UpsertCustomTimelineEntry, string(marshalled))
	if err != nil {
		return nil, err
	}

	target := GraphqlResponse{}
	json.Unmarshal(body, &target)
	if target.Data.UpsertCustomTimelineEntry.Error != nil {
		return nil, fmt.Errorf("Graphql error: %s", *target.Data.UpsertCustomTimelineEntry.Error)
	}
	
	return target.Data.UpsertCustomTimelineEntry, nil
}

type issueTypeIn struct {
	Input CreateIssueTypeInput `json:"input,omitempty"`
}
	
func (c *PlainClient) CreateIssueType(input CreateIssueTypeInput) (*CreateIssueTypeOutput, error) {
	mutationInput := issueTypeIn{
		Input: input,
	}
	marshalled, err := json.Marshal(&mutationInput)
	if err != nil {
		return nil, err
	}

	body, err := c.Query("createIssueType", graphql.CreateIssueType, string(marshalled))
	if err != nil {
		return nil, err
	}

	target := GraphqlResponse{}
	json.Unmarshal(body, &target)
	if target.Data.CreateIssueType.Error != nil {
		return nil, fmt.Errorf("Graphql error: %s", *target.Data.CreateIssueType.Error)
	}
	
	return target.Data.CreateIssueType, nil
}

type issueIn struct {
	Input CreateIssueInput `json:"input,omitempty"`
}
	
func (c *PlainClient) CreateIssue(input CreateIssueInput) (*CreateIssueOutput, error) {
	mutationInput := issueIn{
		Input: input,
	}
	marshalled, err := json.Marshal(&mutationInput)
	if err != nil {
		return nil, err
	}

	body, err := c.Query("createIssue", graphql.CreateIssue, string(marshalled))
	if err != nil {
		return nil, err
	}

	target := GraphqlResponse{}
	json.Unmarshal(body, &target)
	if target.Data.CreateIssue.Error != nil {
		return nil, fmt.Errorf("Graphql error: %s", *target.Data.CreateIssue.Error)
	}
	
	return target.Data.CreateIssue, nil
}