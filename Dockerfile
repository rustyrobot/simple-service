FROM golang:alpine

ENV service /service
WORKDIR ${service}
ADD . ${service}
RUN go build main.go

EXPOSE 80

ENTRYPOINT ["main", "--bind", ":80"]
