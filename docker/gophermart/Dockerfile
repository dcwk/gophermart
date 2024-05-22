FROM golang:1.21 as build

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .

RUN mkdir -p build
RUN go build -trimpath -ldflags="-w -s" -o build/gophermart ./cmd/gophermart
CMD ["/app/build/gophermart"]

