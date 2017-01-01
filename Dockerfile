# Use the golang image
FROM golang
MAINTAINER landon.wainwright@gmail.com

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/landonia/landotube
ADD ./posts /usr/local/landotube/posts
ADD ./templates /usr/local/landotube/templates
ADD ./assets /usr/local/landotube/assets

# Go and fetch all the dependencies
WORKDIR /go/src/github.com/landonia/landotube/landoblog
RUN go get

# Install the actual blog
RUN go install github.com/landonia/landotube/landoblog

# Run the bootstrap command by default when the container starts.
ENTRYPOINT ["/go/bin/landoblog", "-address=:8080", "-pdir=/usr/local/landotube/posts", "-tdir=/usr/local/landotube/templates", "-adir=/usr/local/landotube/assets"]

# Document that the service listens on port 8080.
EXPOSE 8080
