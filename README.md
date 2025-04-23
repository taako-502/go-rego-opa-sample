# go-rego-opa-sample

A Go demo project demonstrating authorization using Open Policy Agent (OPA) Rego based on user attributes.

## Prerequisites

- Go 1.18 or later installed

## Setup

1. Initialize modules and install dependencies (already done):
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

## Project Structure

- `main.go`: Sample code loading and evaluating `data.demo.allow` rule.
- `policy/policy.rego`: Rego policy defining rules based on user role, groups, and age.

## Next Steps

- Enhance `policy/policy.rego` with additional attributes and rules.
- Update `main.go` to accept dynamic input (e.g., JSON via flags or stdin).
- Integrate CLI flags for user input and policy path.

