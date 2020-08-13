FROM golang:1.10-alpine3.8 AS multistage
WORKDIR /src
COPY main.go main_test.go ./
RUN go build main.go
#FROM alpine:3.8
#COPY --from=multistage /src/main /src/main
EXPOSE 8000
CMD ["/src/main"]
