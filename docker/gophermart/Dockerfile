FROM golang:1.21 as build

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .
#
RUN mkdir -p build
EXPOSE 8081
