FROM nanodatacenter/bash_basic 
ENV CONTAINER_NAME "jquery_file_loader" 
 
VOLUME [/data/,/files/] 
COPY ./jquery_file_loader .
COPY ./run.bsh .
RUN mkdir jquery_files
COPY ./jquery_files jquery_files


