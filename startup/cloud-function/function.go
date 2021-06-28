// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"fmt"
	"google.golang.org/api/compute/v1"
	"log"
	"net/http"
	"os"
)

var ProjectID = ""
var Zone = ""
var Region = ""
var InstanceName = ""

func StartInstance(w http.ResponseWriter, r *http.Request) {
	ProjectID = os.Getenv("PROJECT_ID")
	Zone = os.Getenv("ZONE")
	Region = os.Getenv("REGION")
	InstanceName = os.Getenv("INSTANCE_NAME")

	cs, err := InitComputeService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	_, err = GetInstance(cs)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error() + " instance does not exist"))
		log.Print(err)
	}

	cs.Instances.Start(ProjectID, Zone, InstanceName).Do()

	fmt.Fprint(w, "Factorio server instance Started!")
}

// GetInstance passes in the instance name supplied and retrieves it.
// An error indicates an instance that was never created.
// A non-nil error indicates an instance is present whether in the RUNNING or TERMINATED state.
func GetInstance(computeService *compute.Service) (*compute.Instance, error) {
	return computeService.Instances.Get(ProjectID, Zone, InstanceName).Do()
}

// InitComputeService obtains the compute service that allows us to use the compute API
func InitComputeService() (*compute.Service, error) {
	ctx := context.Background()
	return compute.NewService(ctx)
}
