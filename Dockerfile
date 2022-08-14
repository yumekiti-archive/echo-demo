FROM golang:1.18

ENV TZ /usr/share/zoneinfo/Asia/Tokyo
ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .


WORKDIR /go/src/app/api
EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]