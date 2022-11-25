# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl git
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init #--parseDependency

RUN go build -o /birdazzone-api

ENV HOST 0.0.0.0
ENV PORT 8080
EXPOSE 8080

CMD [ "/birdazzone-api" ]
