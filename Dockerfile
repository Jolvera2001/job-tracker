FROM golang:latest AS builder

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o output-binary ./cmd/main

FROM scratch

COPY --from=builder /output-binary /output-binary

WORKDIR /app
EXPOSE 8080

ENTRYPOINT [ "/app/output-binary" ]