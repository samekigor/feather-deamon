package registryclient

import (
	"log"
	"os"

	"github.com/samekigor/feather-deamon/internal/config"
	"oras.land/oras-go/v2/content/oci"
)

var store *oci.Store

func InitOciLayoutStore() {
	var err error

	layoutPath := config.GetValueFromConfig("registry.oci-layout-root")

	if _, err := os.Stat(layoutPath); os.IsNotExist(err) {
		log.Printf("Directory %s does not exist. Creating it...", layoutPath)

		err = os.MkdirAll(layoutPath, 0755)
		if err != nil {
			log.Panicf("Failed to create directory: %v", err)
		}
	}

	store, err = oci.New(layoutPath)
	if err != nil {
		log.Panicf("Failed to create OCI store: %v", err)
	}

}
