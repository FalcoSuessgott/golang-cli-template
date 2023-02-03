FROM alpine:3.17
COPY golang-cli-template /usr/bin/golang-cli-template
ENTRYPOINT ["/usr/bin/golang-cli-template"]