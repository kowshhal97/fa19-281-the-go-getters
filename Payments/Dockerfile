FROM golang:latest 
EXPOSE 9000
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
RUN cd /app ; go install goapi
CMD ["/app/bin/goapi"]
