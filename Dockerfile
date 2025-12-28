# Tahap pertama: builder menggunakan Go 1.23
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -tags netgo -ldflags="-w -s" -o invoice-app ./cmd

# Tahap kedua: runtime minimal
FROM alpine:latest

WORKDIR /root/

# Copy binary dari tahap builder
COPY --from=builder /app/invoice-app .
HEALTHCHECK --interval=30s CMD wget --spider http://localhost:8080 || exit 1


# ✅ Copy folder templates agar available saat runtime
COPY --from=builder /app/templates ./templates

# ✅ Copy static files (CSS, JS, dll)
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./invoice-app"]
