FROM nanodatacenter/basic_python_container
ENV CONTAINER_NAME "eto"
Run pip install redis
Run pip install msgpack       
VOLUME [/data/]


COPY . .
COPY common_tools/ .

