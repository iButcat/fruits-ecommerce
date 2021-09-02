FROM golang:1.17

WORKDIR /api

COPY go.mod .
RUN go mod download ecommerce

COPY . .

RUN go build -o main .
ENTRYPOINT [ "/api/main"]