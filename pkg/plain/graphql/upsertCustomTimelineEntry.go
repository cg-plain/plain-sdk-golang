package graphql

const UpsertCustomTimelineEntry = `mutation upsertCustomTimelineEntry($input: UpsertCustomTimelineEntryInput!) {
    upsertCustomTimelineEntry(input: $input) {
        result
        timelineEntry {
            id
            customerId
            timestamp {
                iso8601
            }
            entry {
                ... on CustomEntry {
                    title
                    components {
                        ... on ComponentText {
                            __typename
                            text
                            textSize
                            textColor
                        }
                        ... on ComponentSpacer {
                            __typename
                            spacerSize
                        }
                        ... on ComponentDivider {
                            __typename
                            spacingSize
                        }
                        ... on ComponentLinkButton {
                            __typename
                            url
                            label
                        }
                    }
                }
            }
            actor {
                ... on MachineUserActor {
                    machineUser {
                        id
                        fullName
                        publicName
                    }
                }
            }
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