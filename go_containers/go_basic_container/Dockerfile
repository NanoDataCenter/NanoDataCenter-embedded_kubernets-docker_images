FROM bash 

ENV TZ=America/Los_Angeles
RUN apk add --no-cache libc6-compat
WORKDIR /pod_util
COPY   ./zoneinfo /usr/share/zoneinfo/
COPY ./process_control .
COPY ./process_control.bsh .
COPY error_logger_directory/error_logger .
