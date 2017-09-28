FROM golang:latest
RUN go get -v -u github.com/kataras/iris
RUN go get -v -u github.com/wirelessregistry/gora
EXPOSE 6060:6060
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]
