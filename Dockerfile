FROM golang:1.25.0-alpine

WORKDIR /app

# airをインストールすることでホットリロードで開発できるよ
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum* ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd

CMD ["air", "-c", ".air.toml"]