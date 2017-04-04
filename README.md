The KrakenD Playground is a demo project that sets a KrakenD API using several
endpoints from a static API.

You can expand this static API just by storing more XML or JSON files in the `data`
folder.

The KrakenD configuration is stored under `krakend/krakend.json` and you can
drag this file anytime to the [KrakenD designer](http://www.krakend.io/designer/) and resume the edition from there.

This demos uses the KrakenD free version, and is limited to 1000rps and 2 backend endpoints per KrakenD endpoint (if you add more they are ignored)

Start
===

    docker-compose up

Use
===
- The KrakenD API runs in the port 8080
- The static API runs in the port 8000

E.g: [http://localhost:8080/splash]()

Edit endpoints
==============
The KrakenD configuration is stored under `krakend/krakend.json` and you can
drag this file anytime to the [KrakenD designer](http://www.krakend.io/designer/)

The backend API endpoints are just static files in the `data` folder. Add or remove there.


Add your demos!
===============
Please add your own examples by doing a pull request!
