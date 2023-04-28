SCHEMA_URL=https://core-api.uk.plain.com/graphql/v1/schema.graphql

go install git.sr.ht/\~emersion/gqlclient/cmd/gqlclientgen@latest

curl $SCHEMA_URL --output schema.graphql
rm -rf queries.graphql
for FILE in graphql/*; do cat $FILE >> ./queries.graphql; done
gqlclientgen -s schema.graphql -q ./queries.graphql -o pkg/plain/gql.go -n plain