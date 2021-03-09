# GO-GRAPHQL-SKELETON

this is a simple graphql skeleton using golang, gqlgen, and graphql

## how to test
1. git clone this project first
> git clone https://github.com/heru-wijaya/go-grpc-skeleton.git
2. do the step on readme at go-grpc-skeleton
3. use this command
> go run main.go
4. open localhost:8081/playground and input mutation or query

### example of mutation and query
> mutation {
  createAccount(account: { name: "client A" }) {
    id
    name
  }
}

> query {
  accounts(account: { name: "John" }) {
    id
    name
  }
}