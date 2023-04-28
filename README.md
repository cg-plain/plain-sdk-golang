# plain-sdk-golang
A fairly light weight sdk for working with plain's graphql api in golang.

## Updating the SDK
1. Create a graphql file in the graphql folder. This should contain a valid graphql mutation or query.
2. run `./scripts/generate.sh`
3. `git add pkg/plain/gql.go`
4. `git commit && git push`

## Using the SDK
See quickstart.go for the current implemented methods, as well as how to roll and use your own queries.