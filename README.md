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

- KrakenD-CE runs in the port [8080](http://localhost:8080)

- The sample client at port [3000](http://localhost:3000)

- All local datasource endpoints under port [8000](http://localhost:8000)

If you use `docker-machine` you will need to access the services using something like `http://192.168.99.100:PORT` instead of `http://localhost:PORT`.

## Editing the endpoints

To add or remove endpoints, edit the file `krakend/krakend.json`. The easiest way to do it is by **dragging this file to the [KrakenD designer](http://www.krakend.io/designer/)** and download the edited file.

To change the data in the static server (simulating your backend API) edit, add or delete files in the **`data`** folder.

## Contribute!
This repository is the place for everyone to start using KrakenD. Maybe we are too used to KrakenD and we don't realize what would be good to include here for a starter. If it doesn't help you good enough or you think that you can add other demo endpoints or middleware integrations, please open a pull request!

Thanks!
