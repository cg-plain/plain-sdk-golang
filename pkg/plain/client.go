package plain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"go.uber.org/zap"
)

type PlainClient struct {
	apiKey string
	logger *zap.SugaredLogger
}

type Query struct {
	Query         string `json:"query"`
	Variables     string `json:"variables"`
	OperationName string `json:"operationName"`
}

var (
	plainUrl string = "https://core-api.uk.plain.com/graphql/v1"
)

func New(logger *zap.SugaredLogger, apiKey string) *PlainClient {
	return &PlainClient{
		apiKey: apiKey,
		logger: logger,
	}
}

func (c *PlainClient) CreateCustomer(input UpsertCustomerInput) (string, error) {
	b, err := os.ReadFile("./pkg/plain/graphql/upsertCustomer.graphql") // just pass the file name
	if err != nil {
		return "", err
	}
	marshalled, err := json.Marshal(&input)
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
	target := UpsertCustomerOutput{}
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

func (c *PlainClient) CreateHelpscoutCustomTimelineEntry(customerId string, entries []TimelineData) error {
	b, err := os.ReadFile("./pkg/plain/upsertCustomTimelineEntry.graphql") // just pass the file name
	if err != nil {
		return err
	}

	mutationInput := UpsertCustomTimelineEntryInput{
		Input: UpsertCustomTimelineEntryData{
			CustomerID: customerId,
			Title:      "Helpscout History",
			ExternalID: "helpscout-import",
			Components: []Content{
				{
					ComponentContainer: &Container{
						ContainerContent: []Content{
							{
								ComponentText: &Text{
									Text: fmt.Sprintf("Most recent **%d** conversations from HelpScout:", len(entries)),
								},
							},
							{
								ComponentDivider: &Divider{
									DividerSpacingSize: "M",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, entry := range entries {
		component := Content{
			ComponentRow: &Row{
				RowMainContent: []Content{
					{
						ComponentText: &Text{
							Text: fmt.Sprintf("**%s**", entry.Subject),
						},
					},
					{
						ComponentText: &Text{
							Text: fmt.Sprintf("%s", entry.Preview),
						},
					},
					{
						ComponentSpacer: &Spacer{
							SpacerSize: "XS",
						},
					},
					{
						ComponentText: &Text{
							Text:      fmt.Sprintf("%s", entry.Date),
							TextColor: "MUTED",
							TextSize:  "S",
						},
					},
				},
				RowAsideContent: []Content{
					{
						ComponentLinkButton: &LinkButton{
							LinkButtonLabel: "🔗 Helpscout",
							LinkButtonUrl:   "http://helpscout.com",
						},
					},
				},
			},
		}
		container := mutationInput.Input.Components[0].ComponentContainer
		container.ContainerContent = append(container.ContainerContent, component)
	}
	marshalled, err := json.Marshal(&mutationInput)
	if err != nil {
		return err
	}

	fullQuery := Query{
		Query:         string(b),
		Variables:     string(marshalled),
		OperationName: "upsertCustomTimelineEntry",
	}

	marshalledQuery, err := json.Marshal(&fullQuery)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, plainUrl, bytes.NewReader(marshalledQuery))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	target := GraphqlResponse{}
	json.Unmarshal(resBody, &target)
	if target.Error != nil {
		return fmt.Errorf("Graphql error: %s", *target.Error)
	}
	c.logger.Debugf("Entry was %s, id: %s\n", target.Data.UpsertCustomTimelineEntry.Result, target.Data.UpsertCustomTimelineEntry.TimelineEntry.ID)
	return nil
}
