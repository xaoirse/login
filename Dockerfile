FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app

# <- COPY go.mod and go.sum files to the workspace
# COPY go.mod .
# COPY go.sum .
# Add this go mod download command to pull in any dependencies
RUN go get github.com/gorilla/sessions
RUN	go get github.com/labstack/echo-contrib/session
RUN	go get github.com/labstack/echo
# Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]