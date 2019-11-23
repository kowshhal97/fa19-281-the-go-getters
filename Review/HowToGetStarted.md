## Steps to Build Menu API

## Run the module locally -

#### 1. Set your GOPATH to the project directoy in your terminal

```
export GOPATH=" Path to your project directory"
```

#### 2.Install the dependencies. You can do that manually or use the Makefile provided.

#### 3. Build your module from within the appropriate directory
```
go build Review
```
This steps makes the 'menu' executable in your folder.

#### 4. Run your module locally on port 3000
```
./Review
```
This step will start the server on localhost on port 3000
```
[negroni] listening on :3000
```

#### 5. Test the API on Postman or on Terminal 
POSTMAN -
```
Request
http://localhost:3000/reviews/ping
```
```
Response
{
"The Go Getters Review API version 1.0 ALIVE!"
}
```

---

## Running the Order microservice in EC2 docker instance 

#### Prerequsites 
#### 1. Install docker on EC2 instance 
#### 2. Start Docker

sudo systemctl start docker<br>
sudo systemctl is-active docker

#### 3. Login to your docker hub account

sudo docker login

#### 4. Create Docker file 

sudo vi Dockerfile

```
  
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
```

#### 5. Build Docker image locally 

```
sudo docker build -t review . 
sudo docker images
```
#### 6. Tag the Image 

```
docker tag review namanagrawal54/review:2
```

#### 7. Push this new imgae to Docker Hub

```
docker push namanagrawal54/review:2
```
