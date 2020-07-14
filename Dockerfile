
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.10

# Copy the local package files to the container's workspace.
ADD . /go/src/mbrdi/food-aggregator

# Build the food-aggregator service inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install mbrdi/food-aggregator

# Run the food-aggregator service by default when the container starts.
ENTRYPOINT /go/bin/food-aggregator

# Document that the service listens on port 3000.
EXPOSE 3000
