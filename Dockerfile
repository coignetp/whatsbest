# Heroku free version only expose one container
# So the frontend must be built and expose via the backend

# frontend builder
FROM node:alpine AS front-builder

WORKDIR /web/

COPY ./frontend /web

RUN yarn && yarn build

# Java jar builder
FROM golang:1.12.0-alpine3.9

RUN apk update && \
    apk upgrade && \
    apk add git build-base

# files
RUN mkdir /app

ADD ./backend /app

WORKDIR /app

COPY ./backend /app
RUN pwd && ls -la
COPY --from=front-builder /web/dist/ /app/web/dist

RUN go get "github.com/mattn/go-sqlite3"
RUN go build -o backend .

CMD /app/backend -port=$PORT
