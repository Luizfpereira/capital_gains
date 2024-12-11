FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

CMD cd cmd && /bin/bash






