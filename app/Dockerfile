# Step 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev libc-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/bin/lol-team-maker .

# Step 2: Deploy
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/lol-team-maker /app/lol-team-maker
COPY meta/env.yaml /app/meta/env.yaml
CMD ["/app/lol-team-maker"]