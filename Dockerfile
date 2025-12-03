# ============================================
# Multi-stage Dockerfile for Attendance Backend
# ============================================

# ============================================
# Stage 1: Builder
# ============================================
FROM golang:1.20-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
RUN go mod verify

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 for static binary
# -ldflags="-w -s" to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /app/bin/server \
    cmd/server/main.go

# ============================================
# Stage 2: Production
# ============================================
FROM alpine:latest AS production

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Set working directory
WORKDIR /home/appuser

# Copy binary from builder
COPY --from=builder /app/bin/server ./server

# Copy config files (if needed)
COPY --from=builder /app/config ./config

# Change ownership
RUN chown -R appuser:appuser /home/appuser

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./server"]

# ============================================
# Stage 3: Development (with hot reload)
# ============================================
FROM golang:1.20-alpine AS development

# Install development tools
RUN apk add --no-cache git make

# Install Air for hot reload
RUN go install github.com/cosmtrek/air@latest

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Run with Air (hot reload)
CMD ["air", "-c", ".air.toml"]
