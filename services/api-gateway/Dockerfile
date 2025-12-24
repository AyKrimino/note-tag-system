# ---------- build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# copy go mod files first
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# ---------- runtime stage ----------
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
