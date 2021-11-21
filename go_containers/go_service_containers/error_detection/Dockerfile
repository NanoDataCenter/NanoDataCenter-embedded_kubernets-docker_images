FROM nanodatacenter/bash_basic 
ENV CONTAINER_NAME "system_error_detection" 

VOLUME [/data/] 

COPY ./watchdog_incidents/wd_monitoring .
COPY ./incident_streams/incident_monitoring .
COPY ./check_rpc/rpc_monitoring             .
COPY ./streams/stream_monitoring .
COPY ./web_services/web_services .
COPY ./debug_start.bsh .
