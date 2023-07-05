FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
CMD ["/app/main"]
EXPOSE 8080
