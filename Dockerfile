FROM golang:1.22.3-alpine3.19

WORKDIR /app
COPY . /app

# creates webserver file
RUN cd webserver && go build .  
RUN chmod +x ./webserver/webserver

ENTRYPOINT [ "webserver/webserver" ]