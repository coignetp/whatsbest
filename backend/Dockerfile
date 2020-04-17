# We specify the base image we need for our
# go application
FROM golang:1.12.0-alpine3.9

RUN apk update && \
    apk upgrade && \
    apk add git build-base
# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /app

# We copy everything in the root directory
# into our /app directory
ADD . /app

# Used to wait for the database to start in docker-compose
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait

# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /app

# we run go build to compile the binary
# executable of our Go program
RUN go get "github.com/lib/pq"
RUN go build -o backend .

# Our start command which kicks off
# our newly created binary executable
CMD sleep 5s && "/app/backend" -dev -port=8081
