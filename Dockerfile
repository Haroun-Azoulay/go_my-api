FROM debian:trixie


RUN apt-get update && \
    apt-get install -y \
    ca-certificates \
    golang-go

WORKDIR /app


RUN go mod init my-api

Run go mod tidy

COPY . .

EXPOSE 8080

ENTRYPOINT [ "go", "run", "main.go" ]