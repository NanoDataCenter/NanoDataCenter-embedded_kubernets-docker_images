FROM nanodatacenter/basic_python_container
RUN pip install msgpack
RUN pip install redis
 
VOLUME [/data/]


COPY . .
COPY common_tools/ .

