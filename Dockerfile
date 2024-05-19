FROM golang:latest

WORKDIR /app

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]