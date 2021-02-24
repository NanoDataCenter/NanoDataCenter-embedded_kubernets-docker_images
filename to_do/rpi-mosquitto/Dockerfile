# Pull base image
FROM python:3.7.5-slim-buster
#FROM nanodatacenter/panadas
MAINTAINER Pascal de Vink <pascal.de.vink@gmail.com>


RUN apt-get update

RUN apt-get install -y wget
RUN apt-get install -y gpg

RUN wget -q -O - http://repo.mosquitto.org/debian/mosquitto-repo.gpg.key | apt-key add -
RUN wget -q -O /etc/apt/sources.list.d/mosquitto-jessie.list http://repo.mosquitto.org/debian/mosquitto-jessie.list
RUN apt-get install -y mosquitto



RUN usermod -aG sudo mosquitto
COPY config /mqtt/config
VOLUME ["/mqtt/config", "/mqtt/data", "/mqtt/log"]

#EXPOSE 8883 
CMD /usr/sbin/mosquitto -c /mqtt/config/mosquitto.conf
