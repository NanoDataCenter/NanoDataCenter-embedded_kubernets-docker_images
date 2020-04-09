FROM nanodatacenter/basic_python_container
ENV CONTAINER_NAME "postgres_setup"
RUN pip install SQLAlchemy
RUN pip install redis
RUN apt-get install -y python-psycopg2
RUN pip install psycopg2
COPY . .
