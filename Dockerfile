FROM golang:1.7
MAINTAINER Christian Winther <cw@bownty.com>

ADD . /go/src/github.com/jippi/consul-rundeck-resource-writer

RUN set -ex \
    && cd /go/src/github.com/jippi/consul-rundeck-resource-writer \
    && go get github.com/kardianos/govendor \
    && govendor install \
    && go install

CMD consul-rundeck-resource-writer
