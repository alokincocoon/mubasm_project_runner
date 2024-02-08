

FROM golang:1.18-alpine

WORKDIR /CREDIT_CARD_VALIDATOR

COPY go.mod go.sum ./

COPY . .

CMD ["go", "run", "main.go"]

EXPOSE 8080