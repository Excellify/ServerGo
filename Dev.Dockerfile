FROM golang:1.16.5-alpine3.13 AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download && go get -u github.com/gobuffalo/packr/v2/packr2
RUN apk add pkgconfig imagemagick-dev build-base

COPY . .

# Build the Go app
RUN packr2 && go build -o seventv

# Start fresh from a smaller image
FROM alpine:3.14
ENV MAGICK_HOME=/usr
RUN apk add --update ca-certificates pkgconfig imagemagick libwebp-tools libwebp-dev libpng-dev jpeg-dev giflib-dev

WORKDIR /app

COPY . /app
COPY --from=build_base /tmp/app/seventv /app/seventv

# Run the binary program produced by `go install`
CMD ["/app/seventv"]
