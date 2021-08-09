FROM nanodatacenter/bash_basic 
ENV CONTAINER_NAME "sqlite3_server" 
 
VOLUME [/data/,/files/] 
COPY ./sqlite3_server .
COPY ./sqlite3_server_test/sqlite3_check .
COPY ./sqlite3_shell .
COPY ./run_test.bsh .

