FROM arm32v7/python 
ENV TZ America/Los_Angeles
ENV INSTALL_PATH /pod_util
RUN mkdir -p $INSTALL_PATH
RUN apt-get update && apt-get install -qq -y \
  build-essential libpq-dev --no-install-recommends
RUN apt-get install -y  gcc python3-dev musl-dev --no-install-recommends
WORKDIR $INSTALL_PATH

RUN pip install --upgrade pip



COPY . .




VOLUME [/data/]



