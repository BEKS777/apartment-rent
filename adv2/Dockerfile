
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o apartment-renting .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/apartment-renting .

CMD ["./apartment-renting"]
