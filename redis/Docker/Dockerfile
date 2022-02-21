FROM redis
ENV CONTAINER_NAME  "redis"

                   
VOLUME [/data/]

WORKDIR  /pod_util

COPY redis-cli  .
COPY redis.conf .
COPY redis-check-rdb .
COPY redis-check-aof .
COPY redis_control.bsh .

