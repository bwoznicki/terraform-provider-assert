---
page_title: "assert_test Data Source - terraform-provider-assert"
subcategory: "Utility"
description: |-
  Assert functionality for testing.
---

# assert_test (Data Source)

The assert provider is a utility for testing terraform configuration with assert like functionality found in common programing languages for testing.

## Example Usage

```hcl
terraform {
  required_providers {
    assert = {
      source  = "bwoznicki/assert"
      version = "0.0.1"
    }
  }
}
// Any condition that evaluates to false will throw error

// make sure all users are using workspaces
// if current workspace is default throw error
data "assert_test" "workspace" {
    test = terraform.workspace != "default"
    throw = "default workspace is not valid in this project"
}

data "aws_region" "current" {}

// region restrictions
// make sure this project is only deployable in certain region
data "assert_test" "region" {
    test = data.aws_region.current.name == "eu-west-1"
    throw = "You cannot deploy this resource in any other region but eu-west-1"
}
```

### Required

- **test** (Boolean) Any valid terraform condition that will return **bool**
- **throw** (String) String printed as terraform error when the test returns **false**

### Optional

- **id** (String) The ID of this resource.


