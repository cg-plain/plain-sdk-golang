package graphql

const CreateIssueType = `mutation createIssueType($input: CreateIssueTypeInput!) {
	createIssueType(input:$input) {
	  issueType{
		  id
			publicName
			isArchived
			defaultIssuePriority {
			  label
			  value
			}
			icon
	  }
	  error {
		message
		type
		code
	  }
	}
  }`