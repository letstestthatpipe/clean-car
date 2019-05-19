FROM golang as builder

ENV GO111MODULE=on

RUN mkdir -p /opt/app

WORKDIR /opt/app

COPY * /opt/app

RUN go build


FROM alpine

EXPOSE 3333

COPY  --from=builder /opt/app/clean-car /usr/bin/

CMD ["/bin/bash", "-c", "/usr/bin/clean-car"]