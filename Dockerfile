FROM golang:1.19-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /go/src

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o app .

######## Start a new stage from scratch #######
FROM alpine

RUN apk update \
    && apk add --no-cache bash

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/app .

COPY config/profiles/default.env ./config/profiles/
COPY ./infra/entrypoint-parent.sh .

RUN chmod 500 entrypoint-parent.sh

# Expose port 80 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./entrypoint-parent.sh"]