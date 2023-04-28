SCHEMA_URL=https://core-api.uk.plain.com/graphql/v1/schema.graphql

go install git.sr.ht/\~emersion/gqlclient/cmd/gqlclientgen@latest

curl $SCHEMA_URL --output schema.graphql
gqlclientgen -s schema.graphql -o pkg/plain/generated_types.go -n plain