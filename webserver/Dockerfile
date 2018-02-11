FROM golang
RUN go get github.com/aws/aws-sdk-go
RUN go get github.com/jeevanullas/go-examples/webserver
ENTRYPOINT /go/bin/webserver
EXPOSE 5000
