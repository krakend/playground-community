package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/devopsfaith/bloomfilter/rpc/client"
)

func main() {
	server := flag.String("server", "krakend_ce:1234", "ip:port of the remote bloomfilter to connect to")
	key := flag.String("key", "jti", "the name of the claim to inspect for revocations")
	flag.Parse()

	c, err := client.New(*server)
	if err != nil {
		log.Println("unable to create the rpc client:", err.Error())
		return
	}
	defer c.Close()

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		subject := *key + "-" + r.FormValue(*key)
		c.Add([]byte(subject))
		http.Redirect(w, r, "/", 302)
	})

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		subject := *key + "-" + r.FormValue(*key)
		fmt.Fprintf(w, "%v", c.Check([]byte(subject)))
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Add("Content-Type", "text/html")
		rw.Write(indexPageContent)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var indexPageContent = []byte(`
<!DOCTYPE html>
<html>
  <head>
    <title>KrakenD JWT Revoker</title>
    <meta name="description" content="KrakenD JWT Revoker - Revoke your JWT remotely.">
  </head>
  <body>
    <h1>KrakenD JWT Revoker - Revoke your JWT remotely.</h1>
    <h2>Add</h2>
	<form action="/add">
	  jti:<br>
	  <input type="text" name="jti" value="the jti to revoke">
	  <br><br>
	  <input type="submit" value="Submit">
	</form>
    <h2>Check</h2>
	<form action="/check">
	  jti:<br>
	  <input type="text" name="jti" value="the jti to revoke">
	  <br><br>
	  <input type="submit" value="Submit">
	</form>
  </body>
</html>
`)
