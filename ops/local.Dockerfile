FROM golang:1.20.11-bullseye

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go install
RUN go install github.com/cosmtrek/air@latest
#RUN air init

EXPOSE 3002
CMD ["air", "-c", ".air.docker.toml"]
