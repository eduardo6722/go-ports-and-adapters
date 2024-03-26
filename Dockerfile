FROM golang:1.22

RUN apt-get update && apt-get install -y sqlite3

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"

RUN go install github.com/golang/mock/mockgen@v1.6.0 && \
    go install github.com/spf13/cobra/cobra@v1.1.3

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]