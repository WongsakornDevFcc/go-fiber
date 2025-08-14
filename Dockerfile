FROM golang:1.25.0

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]
# CMD ["air"]