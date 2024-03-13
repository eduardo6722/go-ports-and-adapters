FROM golang:1.22

USER root
RUN apt-get update && apt-get install -y sqlite3

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"

RUN go install github.com/golang/mock/mockgen@v1.6.0

RUN mkdir -p /var/www/.cache && \
    chown -R root:root /go && \
    chown -R root:root /var/www/.cache

CMD ["tail", "-f", "/dev/null"]