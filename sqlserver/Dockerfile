FROM nanodatacenter/basic_python_container
ENV CONTAINER_NAME "sqlite_server"
Run pip install redis
Run pip install msgpack

VOLUME [/data/,/sqlite/]
COPY . .
COPY common_tools/ .

