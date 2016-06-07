FROM golang:1.6
RUN git clone https://github.com/Mcronk/golang_sample.git
RUN go get github.com/gorilla/mux
EXPOSE 8000
EXPOSE 8080
WORKDIR /go/golang_sample/gApp
RUN go build gorilla_server.go
CMD ["./gorilla_server"]