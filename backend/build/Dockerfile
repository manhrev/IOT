FROM golang:1.18.3-alpine 
WORKDIR /app
COPY ./../ /app/
WORKDIR /app
RUN go build -o /server
RUN apk add --no-cache bash
EXPOSE 1883

CMD ["/server"]