package plain

const createIssue = `mutation createIssue($input: CreateIssueInput!) {
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

  const createIssueType = `mutation createIssueType($input: CreateIssueTypeInput!) {
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