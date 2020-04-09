FROM nanodatacenter/basic_python_container
RUN pip install SQLAlchemy
RUN pip install web3
RUN apt-get install -y python-psycopg2
RUN pip install psycopg2
ENV CONTAINER_NAME "log_stream_events"
COPY . .
