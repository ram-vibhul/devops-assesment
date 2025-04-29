FROM golang:1.21-alpine AS builder
LABEL maintainer="Ram Vibhulaanandh"
# Add a non-root user for build
ENV USER_UID=1001
RUN adduser -D -u ${USER_UID} builder
USER builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server -ldflags="-s -w" main.go && chmod +x server

FROM golang:1.20-alpine

WORKDIR /app/

COPY --from=builder /app/server .
COPY --from=builder /app/config/ ./config/
COPY --from=builder /app/db ./db/

EXPOSE 8080

ENTRYPOINT ["/app/server"]
