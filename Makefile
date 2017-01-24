build:
	docker build . -t jippi/consul-rundeck-resource-writer --network host

publish:
	docker push jippi/consul-rundeck-resource-writer
