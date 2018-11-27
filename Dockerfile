# build stage
FROM golang:1.11.2-alpine3.8 AS build-env
ADD . /go/src/github.com/sqshq/piggymetrics-go/
RUN cd /go/src/github.com/sqshq/piggymetrics-go/app && go build -o piggymetrics

# final stage
FROM alpine:3.8
COPY --from=build-env /go/src/github.com/sqshq/piggymetrics-go/app/ /app/
EXPOSE 80
ENTRYPOINT ./app/piggymetrics