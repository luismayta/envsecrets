# App
.PHONY: app

app.help:
	@echo '    App:'
	@echo ''
	@echo '        app                  show help'
	@echo '        app        	        command=() stage=(dev|stage|prod)'
	@echo '        app.run        	    run server 0.0.0.0 stage=(dev|stage|prod)'
	@echo ''

app: 
	@if [ -z "${command}" ]; then \
		make app.help;\
	fi
	@if [ -z "${stage}" ] && [ -n "${command}" ]; then \
		$(docker-compose) -f ${PATH_DOCKER_COMPOSE}/dev.yml run \
			--rm app bash -c "go run ./main.go ${command}" ; \
	elif [ -n "${stage}" ] && [ -n "${command}" ]; then \
		$(docker-compose) -f ${PATH_DOCKER_COMPOSE}/${stage}.yml run \
			--rm app bash -c "go run ./main.go ${command}" ; \
	fi

app.run: 
	@if [ -z "${stage}" ]; then \
		$(docker-compose) -f ${PATH_DOCKER_COMPOSE}/dev.yml run \
			--rm --service-ports app bash -c "go run ./main.go graphql" ; \
	else \
		$(docker-compose) -f ${PATH_DOCKER_COMPOSE}/${stage}.yml run \
			--rm --service-ports app bash -c "go run ./main.go graphql"; \
	fi


