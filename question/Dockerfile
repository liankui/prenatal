FROM golang

RUN mkdir -p /go/src/github.com/liankui/prenatal-server

ADD . /go/src/github.com/liankui/prenatal-server

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/liankui/prenatal-server/question/cmd  -watch github.com/liankui/prenatal-server/question
