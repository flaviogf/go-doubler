FROM golang:1.19-alpine AS build
WORKDIR /www
COPY go.mod ./
RUN go mod download
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o doubler ./cmd/doubler/main.go

FROM scratch
WORKDIR /www
COPY --from=build /www/doubler ./
ENTRYPOINT ["./doubler"]
