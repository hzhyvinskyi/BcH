.PHONY: build push module-install

build:
	docker build -t hzhyvinskyi/bch-servicea:0.0.1 .

push:
	docker push hzhyvinskyi/bch-servicea:0.0.1

module-install: build push

.DEFAULT_GOAL := module-install
