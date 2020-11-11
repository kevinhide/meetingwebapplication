FROM golang:1.12

# add the source
WORKDIR /go/src/HelpNow/

COPY . /go/src/HelpNow

RUN go get github.com/gorilla/mux && \
go get gopkg.in/mgo.v2 && \
go build -o main .

ADD . /go/src/HelpNow

#build the go app
RUN GOOS=windows GOARCH=amd64 go build -o ./HelpNow ./main.go

EXPOSE 8080

COPY . /go/src/HelpNow/

ENTRYPOINT ["/go/src/HelpNow/-entrypoint.sh"]

CMD ["run"]
