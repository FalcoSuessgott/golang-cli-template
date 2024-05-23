FROM alpine:3.20
COPY golang-cli-template /usr/bin/golang-cli-template
ENTRYPOINT ["/usr/bin/golang-cli-template"]