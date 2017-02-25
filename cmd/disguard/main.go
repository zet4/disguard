package main // import "go.zeta.pm/disguard/cmd/disguard"

import (
	"io/ioutil"
	"log"
	"net/http"

	"flag"

	"github.com/pressly/chi"
	"go.zeta.pm/disguard"
	"gopkg.in/yaml.v2"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "c", "config.yaml", "The path to the configuration file.")
	flag.Parse()
}

func main() {
	in, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	var conf disguard.Config
	err = yaml.Unmarshal(in, &conf)
	if err != nil {
		log.Fatal(err)
	}

	if conf.AuthRoot == "" {
		conf.AuthRoot = "/oauth"
	}

	sess := disguard.NewSessionRouter(&conf)

	root := chi.NewRouter()
	root.Route(conf.AuthRoot, sess.Route)
	root.Mount("/", sess.ReverseHandler())

	http.ListenAndServe(conf.ListenAddress, root)
}
