package gocoprhd

import (
	"fmt"
	"log"
	"os"
	"testing"
  "crypto/tls"
  "net/http"

  "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/victorock/gocoprhd/client"
  "github.com/victorock/gocoprhd/client/block"
  "github.com/victorock/gocoprhd/client/vdc"
  "github.com/victorock/gocoprhd/models"
)

func Init() (*apiclient.CoprHD, runtime.ClientAuthInfoWriter) {
  // create the transport
  transport := httptransport.New("localhost:4443", "/", []string{"https"})
  authInfo := httptransport.BasicAuth( "root", "password")

  // Set Insecure SSL
  transport.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}

  // If not using the Vagrant image, set this environment variable to something other than localhost:4443
  if os.Getenv("GOCOPRHD_ENDPOINT") != "" {
      transport.Host = os.Getenv("GOCOPRHD_ENDPOINT")
  }

  // Get the token to populate header for requests
  if os.Getenv("GOCOPRHD_TOKEN") != "" {
      authInfo = httptransport.APIKeyAuth("X-SDS-AUTH-TOKEN",
                                          "header",
                                          os.Getenv("GOCOPRHD_TOKEN"))
  }

  // Basic Authentication to get the user Token
  if os.Getenv("GOCOPRHD_USERNAME") != "" {
    if os.Getenv("GOCOPRHD_PASSWORD") != "" {
      authInfo = httptransport.BasicAuth(os.Getenv("GOCOPRHD_USERNAME"),
                                          os.Getenv("GOCOPRHD_PASSWORD"))
    }
  }

  // create the API client, with the transport
  client := apiclient.New(transport, strfmt.Default)

  // Login to get our Token
  login, err := client.Authentication.Login(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }

  // Populate the Header with token from now on
  authInfo = httptransport.APIKeyAuth("X-SDS-AUTH-TOKEN",
                                      "header",
                                      login.XSDSAUTHTOKEN)
  return client, authInfo
}

// Test Get Volumes
func TestInit(t *testing.T) {
  client, authInfo := Init()

  //use any function to do REST operations
  resp, err := client.Block.ListVolumes(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)

  fmt.Printf("######################\n")
  for _, v := range resp.Payload.ID {
    fmt.Printf("Volume ID: %#v\n", v)
  }
  fmt.Printf("######################\n")
}

// Test Get Volumes
func TestListVolumes(t *testing.T) {

  // Init
  client, authInfo := Init()

  //use any function to do REST operations
  resp, err := client.Block.ListVolumes(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)
  fmt.Printf("######################\n")
  for _, v := range resp.Payload.ID {
    fmt.Printf("Volume ID: %#v\n", v)
  }
  fmt.Printf("######################\n")

}

// Test Get Snapshots
func TestListSnapshots(t *testing.T) {

  // Init
  client, authInfo := Init()

	//use any function to do REST operations
	resp, err := client.Block.ListSnapshots(nil, authInfo)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Printf("%#v\n", resp.Payload)
  fmt.Printf("######################\n")
  for _, v := range resp.Payload.ID {
    fmt.Printf("Snapshot ID: %#v\n", v)
  }
  fmt.Printf("######################\n")
}

// Test Get Snapshots
func TestCreateVolume(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Construct Request Parameters
  b := &models.CreateVolume {
          ConsistencyGroup: "",
          Count: 1,
          Name: "gocoprhd_test_vol01",
          Project: "urn:storageos:Project:7d46540b-140c-4f39-91b8-52d276356cf0:global",
          Size: 10,
          Varray: "urn:storageos:VirtualArray:ad18dd81-99c6-415d-9081-6091db3df599:vdc1",
          Vpool: "urn:storageos:VirtualPool:7e036b4a-9cba-4357-9afc-3aa7539f10c0:vdc1",
        }

  // Create the New Params and Populate the Body
  CreateVolumeParams := block.NewCreateVolumeParams().
                              WithBody(b)

	//use any function to do REST operations
	resp, err := client.Block.CreateVolume(CreateVolumeParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

// Test Get Tasks
func TestListTasks(t *testing.T) {

  // Init
  client, authInfo := Init()

	//use any function to do REST operations
	resp, err := client.Vdc.ListTasks(nil, authInfo)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Printf("%#v\n", resp.Payload)
  fmt.Printf("######################\n")
  for _, t := range resp.Payload.Task {
    fmt.Printf("Task ID: %#v\n", t.ID)
  }
  fmt.Printf("######################\n")
}

// Test Get Tasks
func TestShowTask(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Create Object to Request
  showTaskParams := vdc.NewShowTaskParams().
                    WithID("urn:storageos:Task:16030dd1-81b6-47ff-9243-a78032134691:vdc1")

	//use any function to do REST operations
	resp, err := client.Vdc.ShowTask(showTaskParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

// Test Get Tasks
func TestListVolumeExports(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Create Object to Request
  listVolumeExportsParams := block.NewListVolumeExportsParams().
                    WithID("urn:storageos:Volume:ab871dd4-e41c-4959-9b92-e9c600c2f612:vdc1")

	//use any function to do REST operations
	resp, err := client.Block.ListVolumeExports(listVolumeExportsParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("%#v\n", resp.Payload)
  fmt.Printf("######################\n")
  for _, exp := range resp.Payload.Itl {
    fmt.Printf("Initiator ID: %#v\n", exp.Initiator.ID)
    fmt.Printf("Initiator Port: %#v\n", exp.Initiator.Port)
    fmt.Printf("Export ID: %#v\n", exp.Export.ID)
    fmt.Printf("Export Name: %#v\n", exp.Export.ID)
    fmt.Printf("Device ID: %#v\n", exp.Device.ID)
    fmt.Printf("Target ID: %#v\n", exp.Target.ID)
  }
  fmt.Printf("######################\n")
}

func TestCreateVolumeSnapshot(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Construct Request Parameters
  b := &models.CreateVolumeSnapshot {
          Name: "coprhd_snap_vol01",
          CreateInactive: true,
          ReadOnly: false,
          Type: "rp",
        }

  // Create Object to Request
  createVolumeSnapshotParams := block.NewCreateVolumeSnapshotParams().
                                      WithID("urn:storageos:Volume:9435e860-e729-478f-8a1b-350da9ce6dd9:vdc1").
                                      WithBody(b)

	//use any function to do REST operations
	resp, err := client.Block.CreateVolumeSnapshot(createVolumeSnapshotParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("%#v\n", resp.Payload)
}

func TestListVolumeSnapshots(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Create Object to Request
  listVolumeSnapshotsParams := block.NewListVolumeSnapshotsParams().
                                      WithID("urn:storageos:Volume:9435e860-e729-478f-8a1b-350da9ce6dd9:vdc1")

	//use any function to do REST operations
	resp, err := client.Block.ListVolumeSnapshots(listVolumeSnapshotsParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("%#v\n", resp.Payload)
  fmt.Printf("######################\n")
  for _, snap := range resp.Payload.Snapshot {
    fmt.Printf("Snapshot ID: %#v\n", snap.ID)
  }
  fmt.Printf("######################\n")
}

func TestDeleteSnapshot(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Create Object to Request
  deleteSnapshotParams := block.NewDeleteSnapshotParams().
                                WithID("urn:storageos:BlockSnapshot:c18f641f-921e-419c-a8bd-4ed5ccb292e7:vdc1")

	//use any function to do REST operations
	resp, err := client.Block.DeleteSnapshot(deleteSnapshotParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("%#v\n", resp.Payload)
}

func TestShowVolume(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Create Object to Request
  showVolumeParams := block.NewShowVolumeParams().
                            WithID("urn:storageos:Volume:9435e860-e729-478f-8a1b-350da9ce6dd9:vdc1")

	//use any function to do REST operations
	resp, err := client.Block.ShowVolume(showVolumeParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("%#v\n", resp.Payload)
}

func TestShowSnapshot(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Create Object to Request
  showSnapshotParams := block.NewShowSnapshotParams().
                            WithID("urn:storageos:BlockSnapshot:b88d6b52-38c4-4456-a92b-110bb4265a2a:vdc1")

	//use any function to do REST operations
	resp, err := client.Block.ShowSnapshot(showSnapshotParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("%#v\n", resp.Payload)
}

func TestCreateExport(t *testing.T) {

  // Init
  client, authInfo := Init()

  // Construct Request Parameters
  b := &models.CreateExport {
          Name: "coprhd_snap_vol01",
          CreateInactive: true,
          ReadOnly: false,
          Type: "rp",
        }

  // Create Object to Request
  createVolumeSnapshotParams := block.NewCreateVolumeSnapshotParams().
                                      WithID("urn:storageos:Volume:9435e860-e729-478f-8a1b-350da9ce6dd9:vdc1").
                                      WithBody(b)

	//use any function to do REST operations
	resp, err := client.Block.CreateVolumeSnapshot(createVolumeSnapshotParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Printf("%#v\n", resp.Payload)
}
