# cloudserver-go-client
Seeweb API client in Go, primarily used by the [Seeweb](https://github.com/Seeweb/terraform-provider) provider in Terraform.

## Installation

In order for the installation to work while having this library in a private repository to will need to execute the following `export GOPRIVATE="github.com/Seeweb/cloudserver-go-client"` before trying to download the library. However if you get the error `fatal: could not read Username for 'https://github.com': terminal prompts disabled`, then you will need to do an additional step, which is described in the following [link.](https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project#providing-private-module-credentials-for-https)

```bash
go get github.com/Seeweb/cloudserver-go-client/seeweb
```

## Example usage
```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
)

func main() {
	client, err := seeweb.NewClient(&seeweb.Config{Token: os.Getenv("SEEWEB_TOKEN"), Debug: true})
	if err != nil {
		panic(err)
	}

	//
	// Server Create
	respSC, raw, err := client.Server.Create(&seeweb.SeewebServerCreateRequest{
		Plan:     "ECS1",
		Location: "it-fr2",
		Image:    "centos-7",
		Notes:    "Server created while executing example",
	})
	if err != nil {
		fmt.Printf("Error Server.Create::> %+v\n", err)
		panic(err)
	}

	fmt.Println("Server name...")
	fmt.Println(respSC.Server.Name)

	// All calls returns the raw *http.Response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Server List
	resp1, raw, err := client.Server.List()
	if err != nil {
		panic(err)
	}

	fmt.Println("Servers...")
	for _, server := range resp1.Server {
		fmt.Println(server.Name)
	}

	// All calls returns the raw *http.Response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Action List
	resp2, raw, err := client.Action.List()
	if err != nil {
		panic(err)
	}

	fmt.Println("Actions...")
	for _, action := range resp2.Actions {
		fmt.Println(action.ID)
	}

	// All calls returns the raw *http.response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Templates List
	resp3, raw, err := client.Template.List()
	if err != nil {
		panic(err)
	}

	fmt.Println("Templates...")
	for _, action := range resp3.Templates {
		fmt.Println(action.ID)
	}

	// All calls returns the raw *http.response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Group Create
	respGC, raw, err := client.Group.Create(&seeweb.SeewebGroupCreateRequest{
		Notes:    "Group created while executing example",
		Password: "secret",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Group name...")
	fmt.Println(respGC.Group.Name)

	// All calls returns the raw *http.response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Group List
	resp4, raw, err := client.Group.List()
	if err != nil {
		panic(err)
	}

	fmt.Println("Groups...")
	for _, action := range resp4.Groups {
		fmt.Println(action.ID)
	}

	// All calls returns the raw *http.response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Region List
	resp5, raw, err := client.Region.List()
	if err != nil {
		panic(err)
	}

	fmt.Println("Regions...")
	for _, action := range resp5.Regions {
		fmt.Println(action.ID)
	}

	// All calls returns the raw *http.response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Plan List
	resp6, raw, err := client.Plan.List()
	if err != nil {
		panic(err)
	}

	fmt.Println("Plans...")
	for _, action := range resp6.Plans {
		fmt.Println(action.ID)
	}

	// All calls returns the raw *http.response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	waitForDeletion := 120
	fmt.Printf("\nGiving some time to Server Creation to finish and proceed with Server.Update and deletion of Server and Group created.\nThe remaining API calls will be executed in %d seconds.\n\nIf the deletion fails. Please go to the console and explicitly delete Server name %q and Group name %q\n\n", waitForDeletion, respSC.Server.Name, respGC.Group.Name)
	time.Sleep(time.Duration(waitForDeletion) * time.Second)

	//
	// Server Update
	respSU, raw, err := client.Server.Update(respSC.Server.Name, &seeweb.SeewebServerUpdateRequest{
		Note: "Server updated while executing the example",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Server update result...")
	fmt.Println(respSU.Status)

	// All calls returns the raw *http.Response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Group Delete
	respGD, raw, err := client.Group.Delete(respGC.Group.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Group deletion status...")
	fmt.Println(respGD.Status)

	// All calls returns the raw *http.response for further inspection.
	fmt.Println(raw.Response.StatusCode)

	//
	// Server Delete
	respSD, raw, err := client.Server.Delete(respSC.Server.Name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server deletion status...")
	fmt.Println(respSD.Action.Status)

	// All calls returns the raw *http.Response for further inspection.
	fmt.Println(raw.Response.StatusCode)
}
```

### Testing

Run all unit tests with `make test`

Run a specific subset of unit test by name using `make test TESTARGS="-v -run TestServer"` which will run all test functions with "TestServer" in their name while `-v` enables verbose output.
