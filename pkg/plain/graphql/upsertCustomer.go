package graphql

const UpsertCustomer = `mutation upsertCustomer($input: UpsertCustomerInput!) {
  upsertCustomer(input: $input) {
    result
    customer {
      id
      externalId
      shortName
      fullName
      email {
        email
        isVerified
      }
      status
    }
    error {
      message
      type
      code
      fields {
        field
        message
        type
      }
    }
  }
}`