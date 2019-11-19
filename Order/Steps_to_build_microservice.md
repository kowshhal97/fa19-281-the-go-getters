## Order Microservice Building Steps

## Run the module locally -

#### 1. Set your GOPATH to the project directoy in your terminal

```
export GOPATH=" Path to your project directory"
```

#### 2.Install the dependencies. You can do that manually or use the Makefile provided.

#### 3. Build your module
```
go build order
```
This steps makes the 'order' executable in your folder.

#### 4. Run your module locally on port 3000
```
./order
```
This step will start the server on localhost on port 3000
```
[negroni] listening on :3000
```

#### 5. Test the API on Postman or on Terminal 
POSTMAN -
```
Request
http://localhost:3000/order/ping
```
```
Response
{
  "Test": "Order API is alive"
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
RUN go get -u   "github.com/gorilla/handlers"
RUN go get -u "github.com/gorilla/mux"
RUN go get -u "github.com/satori/go.uuid"
RUN go get -u   "github.com/unrolled/render"
RUN go get -u   "gopkg.in/mgo.v2"
RUN go get -u   "gopkg.in/mgo.v2/bson"
RUN go get -u "github.com/codegangsta/negroni"
RUN cd /app ; go install order
RUN apt-get update
RUN apt-get -y install vim
RUN ln -sf /dev/stdout /var/log/test.log
CMD ["/app/bin/order"]
```

#### 5. Build Docker image locally 

```
sudo docker build -t order . 
sudo docker images
```
#### 6. Tag the Image 

```
docker tag order 12wsed/order:1
```

#### 7. Push this new imgae to Docker Hub

```
docker push 12wsed/order:1
```





