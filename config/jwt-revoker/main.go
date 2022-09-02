package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/devopsfaith/bloomfilter/rpc/client"
)

func main() {
	server := flag.String("server", "krakend_ce:1234", "ip:port of the remote bloomfilter to connect to")
	key := flag.String("key", "jti", "the name of the claim to inspect for revocations")
	port := flag.Int("port", 8080, "port to expose the service")
	flag.Parse()

	c, err := client.New(*server)
	if err != nil {
		log.Println("unable to create the rpc client:", err.Error())
		return
	}
	defer c.Close()

	tmpl, err := template.New("home").Parse(indexPageContent)

	http.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		subject := *key + "-" + r.FormValue(*key)
		c.Add([]byte(subject))
		log.Printf("adding [%s] %s", *key, subject)
		http.Redirect(w, r, "/", 302)
	})

	http.HandleFunc("/check/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		subject := *key + "-" + r.FormValue(*key)
		res := c.Check([]byte(subject))
		log.Printf("checking [%s] %s => %v", *key, subject, res)
		fmt.Fprintf(w, "%v", res)
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Add("Content-Type", "text/html")
		tmpl.Execute(rw, *key)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

const indexPageContent = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>KrakenD JWT Revoker</title>
    <meta name="description" content="KrakenD JWT Revoker - Revoke your JWT remotely.">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
  </head>
  <body>
  <div class="container">
    <h1>KrakenD JWT Revoker</h1>
    <h2>Revoke your JWT remotely.</h2>
    <div class="row">
    <div class="col-sm">
	    <h3>Add</h3>
		<form action="/add/">
		  {{.}}:<br>
		  <input type="text" name="{{.}}" value="the {{.}} to revoke">
		  <br><br>
		  <input type="submit" value="Submit">
		</form>
    </div>
    <div class="col-sm">
	    <h3>Check</h3>
		<form action="/check/">
		  {{.}}:<br>
		  <input type="text" name="{{.}}" value="the {{.}} to revoke">
		  <br><br>
		  <input type="submit" value="Submit">
		</form>
    </div>
  </div>
    

  </div>
  </body>
</html>
`
