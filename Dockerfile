FROM golang:latest

WORKDIR /app

RUN go mod init user_api

COPY . .

RUN go build -o main .

EXPOSE 8181

CMD [ "./main" ]