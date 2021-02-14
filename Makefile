build:
	-docker build -t remote-exec-bot .

run:
	-docker-compose up -d

stop:
	-docker-compose down

nuke:stop
	-docker image rm remote-exec-bot:latest
