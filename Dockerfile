FROM alpine:3.15
WORKDIR /app
COPY . .
ENTRYPOINT build/tilt-example-api
