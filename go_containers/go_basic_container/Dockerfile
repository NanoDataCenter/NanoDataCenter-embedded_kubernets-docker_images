FROM arm32v7/bash 

ENV TZ=US/Pacific
RUN apk add --no-cache libc6-compat
WORKDIR /pod_util

COPY ./process_control .
COPY ./process_control.bsh .
COPY error_logger_directory/error_logger .
