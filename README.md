# terraform-variable-checker

This small tool checks the current working directory for a Terraform Module and returns all undefined Variables without a default value.

## Usage

```sh
# Test the current working directory for undefined variables
go run main.go

# Test a specific path to a Terraform Module for undefined variables
go run main.go -path /code/iac/terraform/my-module/

# Example Output
go run main.go -path /code/iac/terraform/my-module/
The variable empty_string has no configured default Value!
The variable empty_int has no configured default Value!
```
