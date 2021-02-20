build:
	-docker build -t rex .

push:
	-docker tag rex:latest us.gcr.io/remote-exec-bot-305203/rex:latest
	-docker push us.gcr.io/remote-exec-bot-305203/rex:latest

infra:
	-terraform -chdir=./terraform apply 

run:
	-docker-compose up

stop:
	-docker-compose down

nuke:stop
	-docker image rm rex:latest

clean:
	-rm ./rex
