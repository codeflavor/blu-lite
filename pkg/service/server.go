package service

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"

	c "github.com/codeflavor/servops/pkg/config"
)

type HTTPService struct {
	config *c.AppConfig
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to ServOps API")
}

// Start instantiates a new server connection
func (h *HTTPService) Start() error {
	glog.V(0).Info("Trying to start Database instance...")
	// NOTE: by the time we get here, we should already have a config either self
	// generated, or user set,
	// loaded API resources
	// valid connection to the database
	// bind server according to specified configuration
	http.HandleFunc("/", rootHandle)
	// NOTE: this should be config assigned
	appURI := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	glog.V(0).Infof("Starting server on: %s", appURI)

	if err := http.ListenAndServe(appURI, nil); err != nil {
		return err
	}
	return nil
}

func (s *HTTPService) Status() ([]Status, error) {

}

func (s *HTTPService) Stop() error {}
