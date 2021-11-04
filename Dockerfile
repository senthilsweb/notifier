
#
# First stage: 
# Building a backend.
#

FROM golang:1.16-alpine AS backend

# Move to a working directory (/build).
WORKDIR /build

# Copy and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy a source code to the container.
COPY . .

# Set necessary environmet variables needed for the image and build the server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Run go build (with ldflags to reduce binary size).
RUN go build -ldflags="-s -w" -o notifier .

#
# Second stage: 
# Creating and running a new scratch container with the backend binary.
#

FROM golang:1.16-alpine


# Copy binary from /build to the root folder of the scratch container.
COPY --from=backend ["/build/notifier", "/notifier"]

EXPOSE 3000

# Command to run when starting the container.
CMD ["/notifier", "-p", "3000"]