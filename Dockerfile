FROM golang:1.24.0

WORKDIR /app

RUN apt-get update && apt-get install bash

ENV GO111MODULE on

RUN go install github.com/air-verse/air@latest \
	&& go install github.com/x-motemen/gore/cmd/gore@latest \
	&& go install github.com/spf13/cobra-cli@latest

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . /app

RUN go build -o bin/siggy main.go

RUN echo 'PATH=$PATH:/app/bin' > /root/.bashrc

RUN curl -SL https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait -o /wait
RUN chmod +x /wait
