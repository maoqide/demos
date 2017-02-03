From python:2.7-alpine

ENV PATH $PATH:/root
WORKDIR /root

ADD ./smtp.py /root/smtp.py
ADD ./smtp.sh /root/smtp.sh

ENTRYPOINT ["smtp.sh"]
