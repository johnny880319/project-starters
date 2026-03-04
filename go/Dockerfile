# syntax=docker/dockerfile:1

# Build stage
FROM docker.io/library/golang:1.25-alpine AS builder

RUN apk --no-cache add ca-certificates
# If your app depends on local timezone data, install tzdata here.
# RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/my-app ./cmd/my-app

# Runtime stage
FROM scratch

# Include CA certs so HTTPS calls work.
COPY --from=builder --chown=65532:65532 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder --chown=65532:65532 /out/my-app /my-app

# Add EXPOSE only when this image runs a network service.
# EXPOSE 8080

# Run as non-root for safer defaults.
USER 65532:65532

ENTRYPOINT ["/my-app"]
