FROM nanodatacenter/basic_python_container
ENV CONTAINER_NAME  "irrigation_scheduling"
Run pip install redis
Run pip install msgpack             
VOLUME [/data/]


COPY . .
COPY common_tools/ .

