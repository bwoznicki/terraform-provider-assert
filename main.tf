resource "null_resource" "count" {
    provisioner "local-exec" {
        command = "echo 123"
    }
}


data "assert_test" "workspace" {
    depends_on = [
        null_resource.count
    ]
    test = terraform.workspace != "default"
    throw = "default workspace is not valid in this project"
}
