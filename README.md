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

## Installation
Terraform reads the custom provider binaries from two locations:
* current location where terraform main.tf is run from 
* and 
  * **%APPDATA%\terraform.d\plugins\windows_amd64** on Windows  
  * **~/.terraform.d/plugins/darwin_amd64** on Mac  
  * **~/.terraform.d/plugins/linux_amd64** on Linux  

Place the binary in plugins directory for your platform and run **terraform init** in your project to load custom plugin configuration.

If you want to compile the binary your self make sure you have GO installed and $PATH set correctly, clone this repo and run:
```
$ go build -o terraform-provider-assert
```
this will produce the binary in your current dir, just move it to your plugins dir and initialize terraform.