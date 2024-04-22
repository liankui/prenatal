#!/bin/bash
PWD=$(cd "$(dirname "$0")" && pwd)
cd "$PWD"

docker run -d -p 8091:8080 -v "$PWD"/example:/tmp -e SWAGGER_FILE=/tmp/quiz.yaml swaggerapi/swagger-editor
