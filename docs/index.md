---
page_title: "assert Provider"
subcategory: "Utility"
description: |-
  Assert functionality for testing.
---

# assert Provider

The assert provider is a utility for testing terraform configuration with assert like functionality found in common programing languages for testing.

## Example Usage

```hcl
// Any condition that evaluates to false will throw error

// make sure all users are using workspaces
// if current workspace is default throw error
data "assert_test" "workspace" {
    test = terraform.workspace != "default"
    throw = "default workspace is not valid in this project"
}
```
