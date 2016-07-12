package gocoprhd

import (
	"fmt"
	"log"
	"os"
	"testing"
  "crypto/tls"
  "net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/victorock/gocoprhd/client"
  "github.com/victorock/gocoprhd/client/block"
  "github.com/victorock/gocoprhd/client/vdc"
  "github.com/victorock/gocoprhd/models"
)

// Test Get Volumes
func TestListVolumes(t *testing.T) {

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

  //use any function to do REST operations
  resp, err := client.Block.ListVolumes(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)

}

// Test Get Snapshots
func TestListSnapshots(t *testing.T) {

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

	//use any function to do REST operations
	resp, err := client.Block.ListSnapshots(nil, authInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

// Test Get Snapshots
func TestCreateVolume(t *testing.T) {

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

  // create the transport
  transport := httptransport.New("localhost:4443", "/", []string{"https"})
  authInfo := httptransport.BasicAuth( "root", "password")
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

	//use any function to do REST operations
	resp, err := client.Vdc.ListTasks(nil, authInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

// Test Get Tasks
func TestShowTask(t *testing.T) {

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

  // Create Object to Request
  showTaskParams := vdc.NewShowTaskParams()
  showTaskParams.ID = "urn:storageos:Task:16030dd1-81b6-47ff-9243-a78032134691:vdc1"

	//use any function to do REST operations
	resp, err := client.Vdc.ShowTask(showTaskParams, authInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

//DeleteVolume urn:storageos:Volume:e8b68ad0-bb87-41ed-9dee-7b815cd5e15b:vdc1
