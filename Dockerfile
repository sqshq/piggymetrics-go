# build stage
FROM golang:alpine AS build-env
ADD . /go/src/github.com/sqshq/piggymetrics-go/
RUN cd /go/src/github.com/sqshq/piggymetrics-go/app && go build -o piggymetrics

# final stage
FROM alpine
COPY --from=build-env /go/src/github.com/sqshq/piggymetrics-go/app/ /app/
EXPOSE 8080
ENTRYPOINT ./app/piggymetrics