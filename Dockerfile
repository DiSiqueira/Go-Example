FROM golang:1.8.3-alpine

RUN apk add --update make git

# Install gcc (needed for goose)
RUN apk add --no-cache gcc musl-dev
# Migrations
RUN go get -u github.com/pressly/goose/cmd/goose

# Application
RUN go get -u github.com/Masterminds/glide/...

ADD ./ ./src/github.com/disiqueira/Go-Example
WORKDIR ./src/github.com/disiqueira/Go-Example

RUN make build-app
EXPOSE 80
CMD make run-app

