# Pull base image
FROM resin/rpi-raspbian:jessie
#FROM balenalib/raspberrypi4-64-python:jessie
MAINTAINER Pascal de Vink <pascal.de.vink@gmail.com>


RUN apt-get update
RUN apt-get install -y wget

RUN wget -q -O - http://repo.mosquitto.org/debian/mosquitto-repo.gpg.key | apt-key add -
RUN wget -q -O /etc/apt/sources.list.d/mosquitto-jessie.list http://repo.mosquitto.org/debian/mosquitto-jessie.list
RUN apt-get install -y mosquitto

RUN adduser --system --disabled-password --disabled-login mosquitto

COPY config /mqtt/config
VOLUME ["/mqtt/config", "/mqtt/data", "/mqtt/log"]

#EXPOSE 8883 
CMD /usr/sbin/mosquitto -c /mqtt/config/mosquitto.conf
