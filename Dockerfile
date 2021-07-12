FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags "-X github.com/FalcoSuessgott/golang-cli-template/cmd.version=1.0.0" -o golang-cli-template

CMD [ "./golang-cli-template" ]
