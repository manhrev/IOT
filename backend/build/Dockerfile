FROM golang:1.18.3-alpine 
WORKDIR /app
COPY ./../ /app/
WORKDIR /app/cmd/
RUN go build -o /server
EXPOSE 1883

CMD ["/server"]