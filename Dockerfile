FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk add --no-cache git

RUN mkdir -p /opt/app

WORKDIR /opt/app

COPY . /opt/app/

RUN go build


FROM alpine

EXPOSE 3333

COPY --from=builder /opt/app/clean-car /opt/
COPY swaggerui/ /opt/swaggerui/

CMD ["/opt/clean-car"]
