package service

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"

	"github.com/codeflavor/servops/pkg/config"
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to ServOps API")
}

// StartServer instantiates a new server connection
func StartServer(appConfig *config.AppConfig) error {
	glog.V(0).Info("Trying to start Database instance...")
	if err := InstatiateDbService(appConfig); err != nil {
		return fmt.Errorf("error: An error occured while instantiating the database service %v", err)
	}
	// NOTE: by the time we get here, we should already have a config either self
	// generated, or user set,
	// loaded API resources
	// valid connection to the database
	// bind server according to specified configuration
	http.HandleFunc("/", rootHandle)
	// NOTE: this should be config assigned
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}
	return nil
}
