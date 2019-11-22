FROM golang:latest
EXPOSE 3000
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN go get -u	"github.com/dgrijalva/jwt-go"
RUN go get -u "github.com/gorilla/mux"
RUN go get -u "github.com/lib/pq"
RUN go get -u	"github.com/codegangsta/negroni"
RUN go get -u	"github.com/rs/cors"
RUN go get -u	"github.com/unrolled/render"
RUN go get -u	"github.com/gorilla/handlers"
RUN go get -u "github.com/satori/go.uuid"
RUN go get -u	"gopkg.in/mgo.v2/bson"
RUN go get -u	"gopkg.in/mgo.v2"
RUN cd /app ; go install review
CMD ["/app/bin/review"]
