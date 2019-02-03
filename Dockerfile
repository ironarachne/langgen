# build stage
FROM golang:1.11 AS build-env
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/langgen

# final stage
FROM scratch
COPY --from=build-env /go/bin/langgen /go/bin/langgen
EXPOSE 7479
CMD ["/go/bin/langgen"]