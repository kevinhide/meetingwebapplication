FROM golang:1.12

# add the source
COPY . /go/src/HelpNow
WORKDIR /go/src/HelpNow/

RUN go get github.com/gorilla/mux && \
go get gopkg.in/mgo.v2 && \
go build -o main .

ADD . /go/src/HelpNow

#build the go app
RUN GOOS=windows GOARCH=amd64 go build -o ./HelpNow ./main.go

EXPOSE 8080

ENTRYPOINT [ "bash", "entrypoint.sh" ]

CMD ["/app/main"]
