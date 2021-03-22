# Build Geth in a stock Go builder container
FROM arm32v7/golang


RUN apt-get install git
COPY . .
ENV GOPATH=/go/src

RUN go get github.com/msgpack/msgpack-go
RUN go get github.com/go-redis/redis/v8
RUN go get "github.com/docker/docker/api/types"
RUN go get	"github.com/docker/docker/api/types/container"
RUN go get	"github.com/docker/docker/client"
RUN go get	"github.com/docker/docker/pkg/stdcopy"


# Pull all binaries into a second stage deploy alpine container
#FROM alpine:latest
#COPY --from=builder /go/go-ethereum/build/bin/* /usr/local/bin/

#Run go mod init github.com/glenn-edgar/go_module
#Run go get github.com/vmihailenco/msgpack/v5
#Run go build main.go

VOLUME [/data/]

