# terraform-provider-assert

Custom provider to enable assert like functionality in terraform.  

Example usage:  
Throw exception if user is not using workspaces ( current workspace is default )  
**main.tf**
```javascript
data "assert_test" "workspace" {
    test = terraform.workspace != "default" // must return TRUE or  FALSE
    throw = "default workspace is not valid in this project" // must be of type string
}
```

## dynamic vs static variables in test condition
Terraform validates values at two stages:  
* during validation ( static variables passed to terraform or computed during validation )
* during apply ( dynamic variables know after resource is created )

terraform-provider-assert works in both stages however if you pass condition based on computed values from resources it will only fail during apply stage.

## Instalation
Terraform reads the custom provider binaries from two locations:
* current location where terraform main.tf is run from 
* **%APPDATA%\terraform.d\plugins** on Windows and **~/.terraform.d/plugins** on all other platforms.

Place the binary in plugins directory for your platform and run **terraform init** in your project to load custom plagin configuration.

You can compile the Go binary your self