FROM golang:1.22.1-alpine as build
WORKDIR /app/

COPY . .
RUN go mod download
RUN go build -o app ./cmd/postgres-test/main.go

FROM scratch
WORKDIR /app/

COPY --from=build /app/app /app

EXPOSE 8080
ENTRYPOINT [ "./app", "-h", "db" ]