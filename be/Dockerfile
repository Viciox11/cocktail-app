FROM golang:1.14

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN mkdir -p /resources/images

COPY . .

RUN go build -o main .

EXPOSE 5000

CMD ["./main"]

