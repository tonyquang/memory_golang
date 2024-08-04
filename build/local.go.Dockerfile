FROM golang:1.20

RUN apt-get update

RUN GO111MODULE=on go install golang.org/x/tools/cmd/goimports@latest

RUN GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4@v4.14.0 && \
    GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.0 \
