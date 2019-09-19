FROM golang:alpine AS build-env

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/main

FROM alpine
COPY --from=build-env /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"]

