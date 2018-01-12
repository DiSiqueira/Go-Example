FROM golang:1.8.3-alpine

RUN apk add --update make git

# Install gcc (needed for goose)
RUN apk add --no-cache gcc musl-dev
# Migrations
RUN go get -u github.com/pressly/goose/cmd/goose

# Application
RUN go get -u github.com/Masterminds/glide/...
RUN go get -u github.com/githubnemo/CompileDaemon

COPY . ./src/github.com/disiqueira/Go-Example
WORKDIR ./src/github.com/disiqueira/Go-Example

EXPOSE 80

CMD CompileDaemon -build="make build-app" -command="./bin/Go-Example" -exclude-dir="./bin"
