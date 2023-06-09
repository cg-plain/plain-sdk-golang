package graphql

const CreateIssue = `mutation createIssue($input: CreateIssueInput!) {
	createIssue(input:$input) {
	  issue {
		id
		status
		issueType {
		  id
		  publicName
		}
		updatedAt {
		  iso8601
		}
	  }
	  error {
		message
		type
		code
		fields {
		  field
		  message
		}
	  }
	}
  }`