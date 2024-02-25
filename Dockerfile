FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
COPY data /app/data
COPY static /app/static
COPY templates /app/templates
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/data /app/data
COPY --from=builder /app/static /app/static
COPY --from=builder /app/templates /app/templates

EXPOSE 55555
CMD ["./main"]
