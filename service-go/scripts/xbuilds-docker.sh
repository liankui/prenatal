#!/usr/bin/env bash
# Dockerfile用于快速构建，使用本地机器构建二进制，并copy到镜像中
set -e

case "$1" in
"amd64" | "arm64") ARCH="$1" ;;
"") ARCH="amd64" ;;
*) echo "Unknown 1st option, ARCH:$1" && exit 1 ;;
esac

PWD=$(cd "$(dirname "$0")" && pwd)
serviceGoDir=$(cd "$PWD"/.. && pwd)
cmdDir="$serviceGoDir"/cmd
remotePrefix=prenatal
tag=latest

gitBranch=$(git branch --show-current)
gitHash=$(git rev-list -1 HEAD)

(cd "$cmdDir"/quiz-server && ./xbuild.sh linux "$ARCH")

SERVERS=(quiz-server)
cd "$serviceGoDir"/build/quiz-server-docker
for server in "${SERVERS[@]}"; do
  cd ../"$server"-docker
  (
    set -x
    docker build --platform linux/"$ARCH" --build-arg GIT_BRANCH="$gitBranch" --build-arg GIT_HASH="$gitHash" -f Dockerfile -t "$remotePrefix"/"$server":"$tag" "$serviceGoDir"
  )
  #  docker push "$remotePrefix"/"$server":"$tag"
done
