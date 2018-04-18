FROM golang:alpine

ENV AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY

WORKDIR /go/src/app

COPY . .

RUN go get -u github.com/labstack/echo/
RUN go get -u github.com/aws/aws-sdk-go/

WORKDIR /go/src/app/src/main

EXPOSE 1323

CMD ["go", "build", "main.go"]