## Steps to Build Menu API

## Run the module locally -

#### 1. Set your GOPATH to the project directoy in your terminal

```
export GOPATH=" Path to your project directory"
```

#### 2.Install the dependencies. You can do that manually or use the Makefile provided.

#### 3. Build your module from within the appropriate directory
```
go build menu
```
This steps makes the 'menu' executable in your folder.

#### 4. Run your module locally on port 8001
```
./menu
```
This step will start the server on localhost on port 8001
```
[negroni] listening on :8001
```

#### 5. Test the API on Postman or on Terminal 
POSTMAN -
```
Request
http://localhost:8001/menu/ping
```
```
Response
{
"Test": "Pizza My Heart API Server is UP!: 172.17.0.2"
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
RUN cd /app ; go install menu
RUN apt-get update
RUN apt-get -y install vim
RUN ln -sf /dev/stdout /var/log/test.log
CMD ["/app/bin/menu"]
```

#### 5. Build Docker image locally 

```
sudo docker build -t menu . 
sudo docker images
```
#### 6. Tag the Image 

```
docker tag menu anjanamenoncherubala/menu-api:latest
```

#### 7. Push this new imgae to Docker Hub

```
docker anjanamenoncherubala/menu-api:latest
```





