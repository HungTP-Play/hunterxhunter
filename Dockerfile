FROM golang:1.19.6

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .