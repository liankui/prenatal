#!/usr/bin/env bash
# Dockerfile-action用于多平台构建，使用docker容器分阶段构建镜像
set -e

case "$1" in
"amd64" | "arm64") ARCH="$1" ;;
"") ARCH="amd64" ;;
*) echo "Unknown 1st option, ARCH:$1" && exit 1 ;;
esac

PWD=$(cd "$(dirname "$0")" && pwd)
serviceGoDir=$(cd "$PWD"/.. && pwd)
remotePrefix=ericyami
tag=latest

gitBranch=$(git branch --show-current)
gitHash=$(git rev-list -1 HEAD)

SERVERS=(quiz-server)
cd "$serviceGoDir"/build/quiz-server-docker
for server in "${SERVERS[@]}"; do
  cd ../"$server"-docker
  (
    set -x
    docker build -f Dockerfile-action --platform linux/"$ARCH" --build-arg GIT_BRANCH="$gitBranch" --build-arg GIT_HASH="$gitHash" -t "$remotePrefix"/prenatal-"$server":"$tag" "$serviceGoDir"/..
    docker push "$remotePrefix"/prenatal-"$server":"$tag"
  )
done
