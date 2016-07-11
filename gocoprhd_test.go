package gocoprhd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/victorock/gocoprhd/client"
)

// Test Get Volumes
func TestBlockListVolumes(t *testing.T) {

  // create the transport
  transport := httptransport.New("localhost:4443", "/", []string{"https"})
  authInfo := httptransport.BasicAuth( "root", "password")


  // If not using the Vagrant image, set this environment variable to something other than localhost:4443
  if os.Getenv("GOCOPRHD_ENDPOINT") != "" {
      transport.Host = os.Getenv("GOCOPRHD_ENDPOINT")
  }

  // Get the token to populate header for requests
  if os.Getenv("GOCOPRHD_TOKEN") != "" {
      authInfo = transport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
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
  resp, err := client.Authentication.Login(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }

  // Populate the Header with token from now on
  authInfo = transport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
                                    "header",
                                    resp.XSDSAUTHTOKEN)

  //use any function to do REST operations
  resp, err = client.Block.ListVolumes(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)
}

// Test Get Snapshots
func TestBlockListSnapshots(t *testing.T) {

  // create the transport
  transport := httptransport.New("localhost:4443", "/", []string{"https"})
  authInfo := httptransport.BasicAuth( "root", "password")


  // If not using the Vagrant image, set this environment variable to something other than localhost:4443
  if os.Getenv("GOCOPRHD_ENDPOINT") != "" {
      transport.Host = os.Getenv("GOCOPRHD_ENDPOINT")
  }

  // Get the token to populate header for requests
  if os.Getenv("GOCOPRHD_TOKEN") != "" {
      authInfo = transport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
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
  resp, err := client.Authentication.Login(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }

  // Populate the Header with token from now on
  authInfo = transport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
                                    "header",
                                    resp.XSDSAUTHTOKEN)

	//use any function to do REST operations
	resp, err = client.Snapshot.ListSnapshots(nil, authInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

// Test Get Snapshots
func TestBlockCreateVolume(t *testing.T) {

  // create the transport
  transport := httptransport.New("localhost:4443", "/", []string{"https"})
  authInfo := httptransport.BasicAuth( "root", "password")


  // If not using the Vagrant image, set this environment variable to something other than localhost:4443
  if os.Getenv("GOCOPRHD_ENDPOINT") != "" {
      transport.Host = os.Getenv("GOCOPRHD_ENDPOINT")
  }

  // Get the token to populate header for requests
  if os.Getenv("GOCOPRHD_TOKEN") != "" {
      authInfo = transport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
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
  resp, err := client.Authentication.Login(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }

  // Populate the Header with token from now on
  authInfo = transport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
                                    "header",
                                    resp.XSDSAUTHTOKEN)
  // Construct Request Parameters
  b := &CreateVolume {
          ConsistencyGroup: nil,
          Count: "1",
          Name: "test_vol01",
          Project: "project01",
          Size: "10GB",
          Varray: "varray01",
          Vpool: "vpool01",
        }
  CreateVolumeParams := client.Block.NewCreateVolumeParams().
                        WithBody(&b);

	//use any function to do REST operations
	resp, err = client.Block.CreateVolume(&CreateVolumeParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
