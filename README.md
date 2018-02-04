![KrakenD Playground logo](logo.png)

KrakenD Playground
====

Based on the [KrakenD framework](https://github.com/devopsfaith/krakend), we build and distribute the [KrakenD API Gateway](http://wwww.krakend.io) (or KrakenD Community Edition). But the framework allows you to easily build other API Gateways running different engines.

Since API Gateways feed from APIs, we have also included a web server with fake data that you can modify to test the product. You can expand this static API just by adding more XML or JSON files in the `data` folder.

The KrakenD configuration is stored under `krakend/krakend.json` and you can drag this file anytime to the [KrakenD designer](http://www.krakend.io/designer/) and resume the edition from there.

## Start!

In order to start all the services just run:

    docker-compose up

## Play!

Fire up your browser, curl, postman, httpie or anything else you like to interact with any of the following ports.

Different versions of KrakenD:

- KrakenD-CE runs in the port [8080](http://localhost:8080)
- A custom KrakenD using **Gin** runs in the port [8081](http://localhost:8081)
- A custom KrakenD using **Mux** runs in the port [8082](http://localhost:8082)
- A custom KrakenD using **Gorilla** runs in the port [8083](http://localhost:8083)
- A custom KrakenD using **Negroni** runs in the port [8084](http://localhost:8084)
- A custom KrakenD using **Gin** + **JWT** runs in the port [8085](http://localhost:8085) (the token issuer is exposed here: http://localhost:8090/token/random_user_id)

The backend data ([LWAN](https://github.com/lpereira/lwan) server):

- All local datasource endpoints under port [8000](http://localhost:8000)

![KrakenD Playground logo](playground.jpg)

If you use `docker-machine` you will need to access the services using something like `http://192.168.99.100:PORT` instead of `http://localhost:PORT`.

## Editing the endpoints

Initially the different KrakenD gateways present the following endpoints:

* `/splash` composes responses from several local datasources
* `/showrss/{id}` composes responses from two RSS feeds
* `/nick/{nick}` composes responses from actual github and bitbucket api endpoints

To add more endpoints, edit the file `krakend/krakend.json`. The easiest way to do it is by **dragging this file to the [KrakenD designer](http://www.krakend.io/designer/)** and download the edited file.

To change the data in the static server (simulating your backend API) edit, add or delete files in the **`data`** folder.

## Available demos

### KrakenD CE

This demo uses the official docker image for the [KrakenD](https://hub.docker.com/r/devopsfaith/krakend/) gateway

	$ curl -i http://${DOCKER_IP}:8080/splash

### Custom KrakenD using Gin

This demo uses the [Gin example](https://github.com/devopsfaith/krakend/blob/master/examples/gin/main.go) provided in the KrakenD framework.

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8081/splash

### Custom KrakenD using Mux

This demo uses the [Mux example](https://github.com/devopsfaith/krakend/blob/master/examples/mux/main.go) provided in the KrakenD framework.

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8082/splash

### Custom KrakenD using Gorilla

This demo uses the [Gorilla example](https://github.com/devopsfaith/krakend/blob/master/examples/gorilla/main.go) provided in the KrakenD framework.

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8083/splash

### Custom KrakenD using Negroni

This demo uses the [Negroni example](https://github.com/devopsfaith/krakend/blob/master/examples/negroni/main.go) provided in the KrakenD framework.

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8084/splash

### Custom KrakenD using Gin + JWT

This demo uses the [JWT example](https://github.com/devopsfaith/krakend/blob/master/examples/jwt/main.go) provided in the KrakenD framework.

	$ curl -i -H'Host: ssl.example.com' http://${DOCKER_IP}:8085/splash
	HTTP/1.1 401 Unauthorized
	Content-Security-Policy: default-src 'self'
	Strict-Transport-Security: max-age=315360000; includeSubdomains
	X-Content-Type-Options: nosniff
	X-Frame-Options: DENY
	X-Xss-Protection: 1; mode=block
	Date: Sat, 24 Jun 2017 12:19:14 GMT
	Content-Length: 0
	Content-Type: text/plain; charset=utf-8

If you want to access an endpoint on this service, you must add an `Authorization` header with a valid token issued by the dummy issuer. The following requests demonstrates how the flow works:

	$ curl -i http://$(docker-machine ip dev):8090/token/myUser
	HTTP/1.1 200 OK
	Content-Type: application/json; charset=utf-8
	Date: Sat, 24 Jun 2017 12:22:15 GMT
	Content-Length: 174

	{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Im15VXNlciIsImV4cCI6MTQ5ODMxMDUzNSwiaXNzIjoiaHR0cDovL2V4YW1wbGUuY29tLyJ9.YJgp2qLaPkQ0DVxqGAJ95RBL3e6rEMEY_L-jlqWNrxU"}

	$ curl -iH'Host: ssl.example.com' -H'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6Im15VXNlciIsImV4cCI6MTQ5ODMxMDUzNSwiaXNzIjoiaHR0cDovL2V4YW1wbGUuY29tLyJ9.YJgp2qLaPkQ0DVxqGAJ95RBL3e6rEMEY_L-jlqWNrxU'  http://${DOCKER_IP}:8085/splash
	HTTP/1.1 200 OK
	Cache-Control: public, max-age=300
	Content-Security-Policy: default-src 'self'
	Content-Type: application/json; charset=utf-8
	Strict-Transport-Security: max-age=315360000; includeSubdomains
	X-Content-Type-Options: nosniff
	X-Frame-Options: DENY
	X-Krakend: Version undefined
	X-Xss-Protection: 1; mode=block
	Date: Sat, 24 Jun 2017 12:24:12 GMT
	Transfer-Encoding: chunked

## Contribute!
This repository is the place for everyone to start using KrakenD. Maybe we are too used to KrakenD and we don't realize what would be good to include here for a starter. If it doesn't help you good enough or you think that you can add other demo endpoints or middleware integrations, please open a pull request!

Thanks!
