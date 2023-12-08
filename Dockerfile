FROM alpine:3.19
COPY golang-cli-template /usr/bin/golang-cli-template
ENTRYPOINT ["/usr/bin/golang-cli-template"]