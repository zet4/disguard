package main // import "go.zeta.pm/disguard"

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pressly/chi"
	"gopkg.in/yaml.v2"
)

func main() {
	in, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var conf Config
	err = yaml.Unmarshal(in, &conf)
	if err != nil {
		log.Fatal(err)
	}

	sess := NewSessionRouter(&conf)

	root := chi.NewRouter()
	root.Route("/oauth", sess.Route)
	root.Mount("/", sess.ReverseHandler())

	http.ListenAndServe(conf.ListenAddress, root)
}
