build:
	-docker build -t rex .

push:
	-docker tag rex:latest us.gcr.io/remote-exec-bot-305203/rex:latest
	-docker push us.gcr.io/remote-exec-bot-305203/rex:latest

infra:
	-terraform -chdir=./terraform apply 

provision:
	-ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u rex -i '34.122.190.139,' --private-key '~/.ssh/id_rsa' -e '~/.ssh/id_rsa.pub' ansible/rex.yml --extra-vars "{\"imageID\":\"us.gcr.io/remote-exec-bot-305203/rex:latest\"}"

run:
	-docker-compose up -d

stop:
	-docker-compose down

nuke:stop
	-docker image rm rex:latest

clean:
	-rm ./rex
