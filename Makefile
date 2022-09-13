.PHONY: start stop restart logs compile-flexible-config elastic

start:
	docker-compose up -d

stop:
	docker-compose down --volumes

restart:
	docker-compose restart

logs:
	docker-compose logs -f logstash

compile-flexible-config:
	docker run \
        -v $(PWD)/config/krakend/:/etc/krakend/ \
        -e FC_ENABLE=1 \
        -e FC_SETTINGS=/etc/krakend/settings/dev \
        -e FC_PARTIALS=/etc/krakend/partials \
        -e FC_TEMPLATES=/etc/krakend/templates \
        -e FC_OUT=/etc/krakend/krakend-flexible-config.compiled.json \
        devopsfaith/krakend \
        check -c krakend-flexible-config.tmpl

elastic:
	curl -X POST "localhost:5601/api/saved_objects/_import" -H "kbn-xsrf: true" --form file=@config/elastic/dashboard.ndjson -H "kbn-xsrf: true"