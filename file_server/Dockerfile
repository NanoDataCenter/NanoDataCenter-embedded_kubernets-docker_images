FROM nanodatacenter/basic_python_container 
ENV CONTAINER_NAME "file_server" 
Run pip install redis 
Run pip install msgpack 
VOLUME [/data/,/files/] 
COPY . .
COPY common_tools/ .
