#gocoprhd [![Build Status](https://travis-ci.org/victorock/gocoprhd.svg?branch=dev)](https://travis-ci.org/victorock/gocoprhd)
### Go bindings for CoprHD

## Description
gocoprhd represents API bindings for Go that allow you to manage [CoprHD](https://github.com/CoprHD). Not all functions available on CoprHD API were implemented on this library. The intended functions are a direct implementation of what's available within the [ViPR REST API] (https://www.emc.com/techpubs/api/vipr/v2-4-0-0/index.htm).

The methods always start with a verb, such as:
- List
- Show
- Create
- Update
- Delete
- Expand

The following REST and methods implementations are part of this release:

/login.json:

    operationId: Login

/block/volumes.json:

    operationId: CreateVolume

/block/volumes/{id}/protection/snapshots.json:

    operationId: CreateVolumeSnapshot
    operationId: ListVolumeSnapshots

/block/volumes/bulk.json:

    operationId: ListVolumes

/block/volumes/{id}.json:

    operationId: ShowVolume

/block/volumes/{id}/exports.json:

    operationId: ListVolumeExports

/block/volumes/{id}/expand.json:

    operationId: ExpandVolume

/block/volumes/{id}/deactivate.json:

    operationId: DeleteVolume

/block/exports.json:

    operationId: CreateExport

/block/exports/{id}.json:

    operationId: UpdateExport

/block/exports/{id}/deactivate.json:

    operationId: DeleteExport

/block/snapshots/bulk.json:

    operationId: ListSnapshots

/block/snapshots/{id}.json:

    operationId: ShowSnapshot

/block/snapshots/{id}/deactivate.json:

    operationId: DeleteSnapshot

/block/snapshots/{id}/protection/full-copies.json:

    operationId: CreateSnapshotFullCopy

/vdc/tasks.json:

    operationId: ListTasks

/vdc/tasks/{id}.json:

    operationId: ShowTask

/block/volumes/search.json?{item}={name}:

    operationId: ListVolumeSearch


## Compatibility
gocoprhd is created using [go-swagger](https://github.com/go-swagger/go-swagger) client generator. The local swagger specification is created adopting [EMC ViPR API documentation] (https://www.emc.com/techpubs/api/vipr/v2-4-0-0/index.htm).

## Installation of the client

To install the client:
```
$ go get github.com/victorock/gocoprhd
```

## Generation of the client

The client relies on a swagger spec file that is manually specified at `/swagger-spec/coprhd.yml`. As the REST API evolves, this specification requires
manual update.

The client is found within the `/client` folder which is generated by go-swagger. To generate the client:
```
$ cd $GOPATH/src/github.com/victorock/gocoprhd
$ make
```

The Makefile utilizes [glide](https://github.com/Masterminds/glide) to reference git commit `6b712512cbe1f` of [go-swagger](https://github.com/go-swagger/go-swagger) that has been tested and known to be working. The Makefile will generate a go-swagger directory in `/vendor/github.com/` and then create the client.

## Environment Variables
| Name        | Description           |
| ------------- |:-------------:|
| `GOCOPRHD_ENDPOINT`      | the API endpoint. `localhost:4443` is used by default if not set             |
| `GOCOPRHD_USERNAME`      | The Username for Authentication/Login             |
| `GOCOPRHD_PASSWORD`      | The password for authentication/Login             |
| `GOCOPRHD_TOKEN`      | The token provided by CoprHD/ViPR administrator             |


## Using the client

This sample script will retrieve (List) all volumes. Take a look at `gocoprhd_test.go` for more

```
package main

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

  // Set insecure SSL
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
      authInfo = httptransport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
                                            "header",
                                            os.Getenv("GOCOPRHD_TOKEN"))
  }

  // Basic Authentication to get the user Token
  if os.Getenv("GOCOPRHD_USERNAME") != "" {
    if os.Getenv("GOCOPRHD_PASSWORD") != "" {
      authInfo := httptransport.BasicAuth(os.Getenv("GOCOPRHD_USERNAME"),
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
  authInfo := httptransport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
                                        "header",
                                        login.XSDSAUTHTOKEN)

  //use any function to do REST operations
  resp, err := client.Block.ListVolumes(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)
}
```

### Use the client to import the Service

Import the library for the Block Service:
```
import (
    ...
    client "github.com/emccode/gocoprhd/client/block"
    ...
)
```

Import the Models Library ([See Definitions] (./swagger-spec/coprhd.yml))
```
import (
    ...
    "github.com/emccode/gocoprhd/models"
    ...
)
```

### Create the transport objects

Create the transport and point to the API Gateway Service:
```
  // create the transport
  transport := httptransport.New("localhost:4443", "/", []string{"https"})

  // Set insecure SSL
  transport.Transport = &http.Transport{
        TLSClientConfig: &tls.Config{
          InsecureSkipVerify: true,
        },
      }


  // create the API client, with the transport
  client := apiclient.New(transport, strfmt.Default)
```

### Create the authentication object

Login with Basic Authentication and get the TOKEN for further requests:
```
  authInfo := httptransport.BasicAuth( "root", "password")

  // Login to get our Token
  resp, err := client.Authentication.Login(nil, authInfo)
  if err != nil {
      log.Fatal(err)
  }

  // Populate the Header with token from now on
  authInfo := transport.APIKeyAuth( "X-SDS-AUTH-TOKEN",
                                    "header",
                                    resp.XSDSAUTHTOKEN)

```

### Create the Body with the object of your request

Populate the struct params, and create the object to use in your request:
```
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
```

### Execute the request

Execute the request with the populated struct:
```
    resp, err := client.Block.CreateVolume(CreateVolumeParams, authInfo)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%#v\n", resp.Payload)
```

## Contribution
Create a fork of the project into your own repository.

Run the tests provided in `gocoprhd_test.go` to verify GET/POST/DELETE still function:
```
env DEBUG=1 go test -run TestListVolumes -v
env DEBUG=1 go test -run TestListVolumeExports -v
env DEBUG=1 go test -run TestListSnapshots -v
env DEBUG=1 go test -run TestCreateVolume -v
env DEBUG=1 go test -run TestListTasks -v
env DEBUG=1 go test -run TestShowTask -v
```

If tests do not pass, please create an issue so it can be addressed or fix and create a PR.

If all tests pass, make changes and create a pull request with a description on what was added or removed and details explaining the changes in lines of code. If approved, project owners will merge it.

## Licensing
gocoprhd is freely distributed under the [MIT License](http://emccode.github.io/sampledocs/LICENSE "LICENSE"). See LICENSE for details.

##Support
Please file bugs and issues on the Github issues page for this project. This is to help keep track and document everything related to this repo. For general discussions and further support you can join the [EMC {code} Community slack team](http://community.emccode.com/) and join the **#coprhd** channel. Lastly, for questions asked on [Stackoverflow.com](https://stackoverflow.com) please tag them with **EMC**. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
