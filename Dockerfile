FROM golang:alpine
RUN mkdir /app
COPY main /app
WORKDIR /app
CMD ["/app/main"]
EXPOSE 8080
