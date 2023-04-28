package plain_handwritten

type GraphqlResponse struct {
	Data  GraphqlData `json:"data,omitempty"`
	Error *string     `json:"error,omitempty"`
}

type GraphqlData struct {
	UpsertCustomer            GraphqlObject `json:"upsertCustomer,omitempty"`
	UpsertCustomTimelineEntry GraphqlObject `json:"upsertCustomTimelineEntry,omitempty"`
}
type GraphqlObject struct {
	Customer      Customer      `json:"customer,omitempty"`
	TimelineEntry TimelineEntry `json:"timelineEntry,omitempty"`
	Result        string        `json:"result,omitempty"`
}

type TimelineEntry struct {
	ID string `json:"id,omitempty"`
}
type Customer struct {
	ID        string `json:"id,omitempty"`
	ShortName string `json:"shortName,omitempty"`
	FullName  string `json:"fullFame,omitempty"`
}

type Query struct {
	Query         string `json:"query"`
	Variables     string `json:"variables"`
	OperationName string `json:"operationName"`
}
type CreateCustomerMutationInput struct {
	Input CreateCustomerInput `json:"input"`
}

type CreateCustomerInput struct {
	Identifier Identifier         `json:"identifier"`
	OnCreate   CreateCustomerData `json:"onCreate"`
	OnUpdate   UpdateCustomerData `json:"onUpdate"`
}

type Identifier struct {
	EmailAddress string `json:"emailAddress"`
}

type CreateCustomerData struct {
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
	Email     Email  `json:"email"`
}

type UpdateCustomerData struct {
	FullName  UpdateObject `json:"fullName"`
	ShortName UpdateObject `json:"shortName"`
	Email     Email        `json:"email"`
}

type UpdateObject struct {
	Value string `json:"value,omitempty"`
}

type Email struct {
	Email      string `json:"email,omitempty"`
	IsVerified bool   `json:"isVerified"`
}

type UpsertCustomTimelineEntryInput struct {
	Input UpsertCustomTimelineEntryData `json:"input,omitempty"`
}

type UpsertCustomTimelineEntryData struct {
	CustomerID string    `json:"customerId,omitempty"`
	Title      string    `json:"title,omitempty"`
	ExternalID string    `json:"externalId,omitempty"`
	Components []Content `json:"components,omitempty"`
}

type Container struct {
	ContainerContent []Content `json:"containerContent,omitempty"`
}

type Content struct {
	ComponentContainer  *Container  `json:"componentContainer,omitempty"`
	ComponentText       *Text       `json:"componentText,omitempty"`
	ComponentDivider    *Divider    `json:"componentDivider,omitempty"`
	ComponentRow        *Row        `json:"componentRow,omitempty"`
	ComponentSpacer     *Spacer     `json:"componentSpacer,omitempty"`
	ComponentLinkButton *LinkButton `json:"componentLinkButton,omitempty"`
}

type LinkButton struct {
	LinkButtonLabel string `json:"linkButtonLabel,omitempty"`
	LinkButtonUrl   string `json:"linkButtonUrl,omitempty"`
}

type Spacer struct {
	SpacerSize string `json:"spacerSize,omitempty"`
}
type Text struct {
	Text      string `json:"text,omitempty"`
	TextColor string `json:"textColor,omitempty"`
	TextSize  string `json:"textSize,omitempty"`
}

type Divider struct {
	DividerSpacingSize string `json:"dividerSpacingSize,omitempty"`
}

type Row struct {
	RowMainContent  []Content `json:"rowMainContent,omitempty"`
	RowAsideContent []Content `json:"rowAsideContent,omitempty"`
}

//-d '{"query":"query getWorkspace($workspaceId: ID!) { workspace(workspaceId: $workspaceId) { id name publicName } }","variables":{"workspaceId":"'"$PLAIN_WORKSPACE_ID"'"},"operationName":"getWorkspace"}'
