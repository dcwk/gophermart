FROM golang:1.21 as build

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .
#
RUN mv cmd/accrual/accrual_darwin_amd64 /usr/local/bin/app
RUN chmod -R 700 /usr/local/bin/app

CMD ["app"]

EXPOSE 8080
