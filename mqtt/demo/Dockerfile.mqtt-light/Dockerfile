FROM maoqide/rpi-gpio

ENV PATH $PATH:/root
WORKDIR /root

RUN pip install paho-mqtt

ADD ./light.py /root/light.py
ADD ./light.sh /root/light.sh

ENTRYPOINT ["light.sh"]
