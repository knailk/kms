FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN go build -o /kms ./cmd/server/...

# Set environment variable
ENV ENVIRONMENT=local

# Expose and run
EXPOSE 8080
CMD ["sh", "-c", "ENV=$ENVIRONMENT /kms"]
