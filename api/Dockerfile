FROM golang:1.24.4-alpine3.22
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go run ingress.go
CMD [ "app" ]