FROM alpine:3.18.4

ENV LANG C.UTF-8
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata

ARG GIT_HASH
ARG GIT_BRANCH
LABEL GitHash=$GIT_HASH GitBranch=$GIT_BRANCH

ADD cmd/quiz-server/quiz-server /quiz-server
ADD configs /configs

CMD [ "/quiz-server" ]
