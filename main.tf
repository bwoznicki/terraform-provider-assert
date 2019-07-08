data "assert_test" "workspace" {
    test = terraform.workspace != "default"
    throw = "default workspace is not valid in this project"
}
