export REPOSITORY=localtest

build:
	docker build -t ${REPOSITORY}/mqtt-laser .
run:
	docker run -d --privileged --env GPIO_PIN=12 --env MQTT_BROKER=192.168.10.245 --env MQTT_TOPIC=laser-light ${REPOSITORY}/mqtt-laser
