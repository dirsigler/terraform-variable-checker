# terraform-variable-checker

This small tool checks the current working directory for a Terraform Module and returns all undefined Variables without a default value.

## Purpose

Terraform variables can be defined without declaring a default value.
During your `terraform plan` or `terraform apply` run Terraform will ask you to manually provide a value.

This approach is technically possible but should not be used, neither in a CI/CD setup nor in a manual management. A variable should always have a predefined value.

With the help of this tool you are able to find those variables without a default value even in big `variables.tf` files or in new big project where you have no clear overview.

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
