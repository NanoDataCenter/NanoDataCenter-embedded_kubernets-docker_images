FROM nanodatacenter/basic_python_container
ENV CONTAINER_NAME "logs_to_blockchain"
RUN apt-get install -y wget
RUN wget https://nodejs.org/dist/v9.9.0/node-v9.9.0-linux-armv7l.tar.xz
RUN tar -xJf node-v9.9.0-linux-armv7l.tar.xz
RUN cp  -R node-v9.9.0-linux-armv7l/* /usr/local/
RUN npm install -g solc
RUN solcjs --help
VOLUME [/ipc,/key_store/]
RUN pip install web3
RUN pip install Jinja2
COPY . .
RUN /bin/bash ./compile_contracts.bsh
RUN python store_build_contracts.py EventHandler  HelloWorld


