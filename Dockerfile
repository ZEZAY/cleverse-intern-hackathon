FROM golang:1.18

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o bin/server cmd/api/main.go
CMD [ "./bin/server" ]