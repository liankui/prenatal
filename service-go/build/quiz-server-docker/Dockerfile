FROM alpine:3.18.4

ARG GIT_HASH
ARG GIT_BRANCH
LABEL GitHash=$GIT_HASH GitBranch=$GIT_BRANCH

ADD cmd/quiz-server/quiz-server /quiz-server
ADD configs /configs

ENTRYPOINT [ "/quiz-server" ]
