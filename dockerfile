############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk upgrade && apk add --no-cache ca-certificates git
RUN update-ca-certificates

WORKDIR /go/src/cjlapao/ms-graph-collector

COPY . .

# Using go get.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/ms-graph-collector

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/ms-graph-collector /go/bin/ms-graph-collector
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run the hello binary.
EXPOSE 1000

ENTRYPOINT ["/go/bin/ms-graph-collector"]