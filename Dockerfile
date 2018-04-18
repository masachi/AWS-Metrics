FROM golang:alpine

ENV AWS_ACCESS_KEY_ID ""
ENV AWS_SECRET_ACCESS_KEY ""

WORKDIR /go/src/app

COPY . /usr/local/go/src/

RUN apk add --no-cache ca-certificates git wget

RUN go get -u github.com/labstack/echo/
RUN go get -u github.com/aws/aws-sdk-go/

EXPOSE 1323

CMD ["go", "build", "/usr/local/go/src/AWS-Metrics/src/main/main.go"]