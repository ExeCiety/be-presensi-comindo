FROM golang:1.20.11-alpine3.18 as builder

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go install
RUN go install github.com/cosmtrek/air@latest
#RUN air init

EXPOSE 3002
CMD ["air", "-c", ".air.docker.toml"]
