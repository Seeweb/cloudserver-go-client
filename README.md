# go-seeweb
Seeweb API client in Go, primarily used by the [Seeweb](https://github.com/uwtrilogyseaward0m/terraform-provider-seeweb) provider in Terraform.

## Installation
```bash
go get github.com/uwtrilogyseaward0m/go-seeweb/seeweb
```

## Example usage
```go
package main

import (
	"fmt"
	"os"

	"github.com/uwtrilogyseaward0m/go-seeweb/seeweb"
)

func main() {
  client, err := seeweb.NewClient(&seeweb.Config{Token: os.Getenv("SEEWEB_TOKEN")})
  if err != nil {
    panic(err)
  }

  resp, raw, err := client.Server.List()
  if err != nil {
    panic(err)
  }

  fmt.Println("Servers...")
  for _, server := range resp.Server {
    fmt.Println(server.Name)
  }

  // All calls returns the raw *http.Response for further inspection.
  fmt.Println(raw.Response.StatusCode)

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
}
```

### Testing

Run all unit tests with `make test`

Run a specific subset of unit test by name using `make test TESTARGS="-v -run TestServer"` which will run all test functions with "TestServer" in their name while `-v` enables verbose output.
