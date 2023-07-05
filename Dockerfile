FROM golang:alpine
RUN mkdir /app
COPY main /app
WORKDIR /app
CMD ["/main"]
EXPOSE 8080
