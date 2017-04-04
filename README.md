KrakenD Playground
====

The KrakenD Playground is a demo project that sets a KrakenD API using several
endpoints from a static API.

You can expand this static API just by storing more XML or JSON files in the `data`
folder.

The KrakenD configuration is stored under `krakend/krakend.json` and you can
drag this file anytime to the [KrakenD designer](http://www.krakend.io/designer/) and resume the edition from there.

## Start

    docker-compose up

## Use

- The KrakenD API runs in the port 8080
- The static API runs in the port 8000

E.g: [http://localhost:8080/splash]()

## Edit endpoints

The KrakenD configuration is stored under `krakend/krakend.json` and you can
drag this file anytime to the [KrakenD designer](http://www.krakend.io/designer/)

The backend API endpoints are just static files in the `data` folder. Add or remove there.

## Available demos

### KrakenD Free

This demo uses the [KrakenD free version](https://hub.docker.com/r/devopsfaith/krakend/), and is limited to 1000rps and 2 backend endpoints per KrakenD endpoint (if you add more they are ignored)

	$ curl -i http://${DOCKER_IP}:8080/splash

### OS KrakenD Gin

This demo uses the [Gin example](https://github.com/devopsfaith/krakend/blob/master/examples/gin/main.go) from the KrakenD OS

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8081/splash

### OS KrakenD Mux

This demo uses the [Mux example](https://github.com/devopsfaith/krakend/blob/master/examples/mux/main.go) from the KrakenD OS

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8082/splash

### OS KrakenD Gorilla

This demo uses the [Gorilla example](https://github.com/devopsfaith/krakend/blob/master/examples/gorilla/main.go) from the KrakenD OS

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8083/splash

### OS KrakenD Negroni

This demo uses the [Negroni example](https://github.com/devopsfaith/krakend/blob/master/examples/negroni/main.go) from the KrakenD OS

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8084/splash

## Add your demos!

Please add your own examples by doing a pull request!
