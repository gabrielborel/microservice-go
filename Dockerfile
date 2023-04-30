FROM golang:1.20
WORKDIR /app
ENTRYPOINT [ "tail", "-f", "/dev/null" ]