# go-seeweb
Seeweb API client in Go, primarily used by the [Seeweb](https://github.com/uwtrilogyseaward0m/terraform-provider-seeweb) provider in Terraform.

### Testing

Run all unit tests with `make test`

Run a specific subset of unit test by name using `make test TESTARGS="-v -run TestServer"` which will run all test functions with "TestServer" in their name while `-v` enables verbose output.
