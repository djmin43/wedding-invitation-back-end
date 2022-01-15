FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build

EXPOSE 10000

CMD ["./marriage-invitation-back-end"]

