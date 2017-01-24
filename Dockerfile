FROM golang:1.7
MAINTAINER Christian Winther <cw@bownty.com>

ADD . /go/src/github.com/jippi/consul-rundeck-resource-writer

RUN set -ex \
    && go get -u github.com/kardianos/govendor \
    && cd /go/src/github.com/jippi/consul-rundeck-resource-writer \
    && govendor install \
    && go install

CMD consul-rundeck-resource-writer

