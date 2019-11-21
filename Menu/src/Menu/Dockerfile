FROM golang:latest
EXPOSE 8001
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN go get -u "github.com/gorilla/handlers"
RUN go get -u "github.com/gorilla/mux"
RUN go get -u "github.com/satori/go.uuid"
RUN go get -u	"github.com/unrolled/render"
RUN go get -u	"gopkg.in/mgo.v2"
RUN go get -u	"gopkg.in/mgo.v2/bson"
RUN go get -u "github.com/codegangsta/negroni"
RUN cd /app ; go install menu
RUN apt-get update
RUN apt-get -y install vim
RUN ln -sf /dev/stdout /var/log/test.log
CMD ["/app/bin/menu"]

