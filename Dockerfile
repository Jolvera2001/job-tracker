FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o output-binary ./cmd/main

FROM scratch

COPY --from=builder /app/output-binary /app/output-binary

WORKDIR /app
EXPOSE 8080

ENTRYPOINT [ "/app/output-binary" ]