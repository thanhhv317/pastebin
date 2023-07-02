# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /Pastebin

EXPOSE 3000

CMD ["go", "run", "Pastebin"]