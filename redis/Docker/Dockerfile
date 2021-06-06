FROM arm32v7/redis
ENV CONTAINER_NAME  "redis"

#RUN apt-get install -y libatomic1
                   
VOLUME [/data/]

WORKDIR  /pod_util

COPY redis-cli  .
COPY redis.conf .
COPY dbx.so .
COPY redistimeseries.so .
COPY redis-check-rdb .
COPY redis-check-aof .
COPY redis_control.bsh .

