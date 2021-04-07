# Build Geth in a stock Go builder container
FROM arm32v7/golang

VOLUME [/data/]


ENV CONTAINER_NAME "flask_web"
COPY . .
RUN /bin/bash -c "echo 'export EDITOR=nano' >> ~/.bashrc"
RUN go get github.com/msgpack/msgpack-go 
RUN go get github.com/go-redis/redis/v8 
RUN go get github.com/polydawn/gosh
RUN go get gopkg.in/mail.v2

RUN go mod init site_control.com

##### step up go path

# Pull all binaries into a second stage deploy alpine container
#FROM alpine:latest
#COPY --from=builder /go/go-ethereum/build/bin/* /usr/local/bin/



