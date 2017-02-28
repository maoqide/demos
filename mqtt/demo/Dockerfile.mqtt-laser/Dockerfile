FROM maoqide/rpi-gpio

ENV PATH $PATH:/root
WORKDIR /root

RUN pip install paho-mqtt

ADD ./mqtt_laser.py /root/mqtt_laser.py
ADD ./mqtt_laser.sh /root/mqtt_laser.sh

ENTRYPOINT ["mqtt_laser.sh"]
