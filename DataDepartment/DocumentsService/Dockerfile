# Étape 1 : Construction de l'application avec l'image officielle de Go
FROM golang:1.23.3 AS builder

WORKDIR /DocumentsService
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod tidy

COPY ./ ./
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o main .

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /DocumentsService/main .

EXPOSE 7000
CMD ["./main"]
