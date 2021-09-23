FROM nanodatacenter/bash_basic 
ENV CONTAINER_NAME "error_detection" 

VOLUME [/data/] 

COPY ./watchdog_incidents/wd_monitoring .
COPY ./incident_streams/incident_monitoring .
COPY ./check_rpc/rpc_monitoring             .
