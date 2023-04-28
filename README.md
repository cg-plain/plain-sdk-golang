# plain-sdk-golang
A fairly light weight sdk for working with plain's graphql api in golang.

## Updating the types
2. run `./scripts/generate.sh`
3. `git add pkg/plain/generated_types.go`
4. `git commit && git push`

## Adding new queries
While many generators offered "support" for generating functions to handle queries, none of them seemed particularly robust or to generate good quality code.
The pattern for how to add a new one can be seen in the existing methods.

## Using the SDK
See quickstart.go for the current implemented methods, as well as how to roll and use your own queries.