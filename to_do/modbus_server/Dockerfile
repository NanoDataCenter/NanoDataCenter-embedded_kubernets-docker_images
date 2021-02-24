FROM nanodatacenter/basic_python_container
ENV CONTAINER_NAME "modbus_server"
RUN pip install  serial
RUN pip install pyserial 
RUN pip install uuid
VOLUME [/data/,/dev]
COPY . .



