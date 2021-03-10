FROM golang:1.13-alpine3.11 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/heru-wijaya/go-graphql-skeleton
COPY . .
RUN go mod vendor
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app .
COPY .env /go/bin

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8081
CMD ["app"]