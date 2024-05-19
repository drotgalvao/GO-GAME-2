FROM golang:latest

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]