FROM alpine:3.15
WORKDIR /app
COPY build build
COPY dev.env dev.env
ENTRYPOINT build/tilt-example-api
