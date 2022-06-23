LOGIN ?= sans
PASS ?= 321
MAI_TOKEN ?= ${MAI_TOKEN}
TEXT ?= qq
TO ?= tilt

.PHONY: up
up:
	docker-compose up --remove-orphans

.PHONY: down
down:
	docker-compose down

.PHONY: build
build:
	docker-compose build

.PHONY: register
register:
	curl -sv -X POST -H 'Content-type: application/json' -d '{"login":"$(LOGIN)", "password":"$(PASS)"}' http://localhost:80/api/auth/register

.PHONY: login
login:
	curl -sv -X POST -H 'Content-type: application/json' -d '{"login":"$(LOGIN)", "password":"$(PASS)"}' http://localhost:80/api/auth/login

.PHONY: send_message
send_message:
	curl -sv -X POST -H 'Content-type: application/json' -H 'Mai-Backend-Auth-Header: $(MAI_TOKEN)' -d '{"text":"$(TEXT)"}' http://localhost:80/api/messages

.PHONY: get_messages
get_messages:
	curl -sv -X GET -H 'Content-type: application/json' -H 'Mai-Backend-Auth-Header: $(MAI_TOKEN)' http://localhost:80/api/messages

.PHONY: send_user_message
send_user_message:
	curl -sv -X POST -H 'Content-type: application/json' -H 'Mai-Backend-Auth-Header: $(MAI_TOKEN)' -d '{"text":"$(TEXT)"}' http://localhost:80/api/users/$(TO)/messages

.PHONY: get_user_messages
get_user_messages:
	curl -sv -X GET -H 'Content-type: application/json' -H 'Mai-Backend-Auth-Header: $(MAI_TOKEN)' http://localhost:80/api/users/me/messages
