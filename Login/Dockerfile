FROM golang:latest
EXPOSE 3000
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN go get -u	"github.com/dgrijalva/jwt-go"
RUN go get -u "github.com/gorilla/mux"
RUN go get -u "github.com/lib/pq"
RUN go get -u	"golang.org/x/crypto/bcrypt"
RUN cd /app ; go install Login
CMD ["/app/bin/Login"]
