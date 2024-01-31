FROM golang:alpine
WORKDIR /app
COPY . /app
RUN go build main.go
EXPOSE 8080
CMD ["/app/main"]

