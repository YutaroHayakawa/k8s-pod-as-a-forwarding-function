FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -o /bin/resolve-cilium-host-ip

ENTRYPOINT "/bin/sh"
