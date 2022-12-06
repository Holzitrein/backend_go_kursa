FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod tidy && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/gateway

CMD ["app"]