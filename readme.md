# GO-GRAPHQL-SKELETON

this is a simple graphql skeleton using golang, gqlgen, and graphql

## how to test
1. git clone this project first
> git clone https://github.com/heru-wijaya/go-grpc-skeleton.git
2. do the step on readme at go-grpc-skeleton
3. build the docker using this command
> docker build -t graphql .
4. run docker
> docker run --network="host" graphql
5. open localhost:8081/playground and input mutation or query.
you can test using postman too with localhost:8081/dummy or localhost:8081/account 

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

## how to add new schema
1. create new folder
2. copy schema.graphql on other folder and edit
3. run this command
> go run github.com/99designs/gqlgen generate

or you can go to new folder and initialize
command for initialize
> go run github.com/99designs/gqlgen init
