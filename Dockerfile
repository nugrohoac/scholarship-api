FROM golang:1.17-alpine as build

WORKDIR /app

COPY . .

# download dependencies
RUN go mod vendor

# build binary
RUN go build -o scholarship-api cmd/app/main.go

#
EXPOSE 7070

CMD [ "./scholarship-api" ]