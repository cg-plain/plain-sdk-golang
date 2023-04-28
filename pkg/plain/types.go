package plain

// helper structs which are used in queries, not autogenerated

type GraphqlResponse struct {
	Data  GraphqlData `json:"data,omitempty"`
}

// to add a new supported query add another line in here with the appropriate type
type GraphqlData struct {
	UpsertCustomer            *UpsertCustomerOutput `json:"upsertCustomer,omitempty"`
	UpsertCustomTimelineEntry *UpsertCustomTimelineEntryOutput `json:"upsertCustomTimelineEntry,omitempty"`
	CreateIssue *CreateIssueOutput `json:"createIssue,omitempty"`
	CreateIssueType *CreateIssueTypeOutput `json:"createIssueType,omitempty"`
}

type Query struct {
	Query         string `json:"query"`
	Variables     string `json:"variables"`
	OperationName string `json:"operationName"`
}