# terraform-provider-assert

Custom provider to enable assert like functionality in terraform.  

Example usage:  
Throw exception if user is not using workspaces ( current workspace is default )  
**main.tf**
```javascript
data "assert_test" "workspace" {
    test = terraform.workspace != "default" // must return TRUE or FALSE
    throw = "default workspace is not valid in this project" // must be of type string
}
```
there is no need to explicitly initialize provider in your .tf files with the **provider "name" {}** block, just use data source directly.

## dynamic vs static variables in test condition
Terraform validates values at two stages:  
* during validation ( static variables passed to terraform or computed during validation )
* during apply ( dynamic variables known after resource is created )

terraform-provider-assert works in both stages, however, if you pass condition based on value computed from resource data it will only **throw** during apply stage.

# Installation  
## terraform < 13.0  
Terraform reads the custom provider binaries from two locations:
* current location where terraform main.tf is run from 
* and 
  * **%APPDATA%\terraform.d\plugins\windows_amd64** on Windows  
  * **~/.terraform.d/plugins/darwin_amd64** on Mac  
  * **~/.terraform.d/plugins/linux_amd64** on Linux  

Place the binary in plugins directory for your platform and run **terraform init** in your project to load custom plugin configuration.

## terraform >= 13.0

from version 13+ the way terraform fetches providers has changed. In your main configuration add required_providers block:  
```
terraform {
  required_providers {
    assert = {
      source  = "github.com/bwoznicki/assert"
    }
  }
}
```
Implicitly terraform still checks your local dir eg. **~/terraform.d/plugins** but you must place the binary in a sligltly different path eg. **.terraform.d/plugins/github.com/bwoznicki/assert/0.0.1/linux_amd64/** for linux.  
You can also specify a custom path where your local providers are fetched from using **~/terraformrc** or **~/terraform.rc** for more info pleace check official documentation https://www.terraform.io/docs/cli/config/config-file.html#provider-installation
