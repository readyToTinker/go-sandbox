FROM golang:1.22.3-alpine3.19

WORKDIR /app
COPY webserver /app/webserver/
COPY go.* /app/

# outputs webserver binary
RUN cd webserver && go build .  
RUN chmod +x ./webserver/webserver

ENTRYPOINT [ "webserver/webserver" ]