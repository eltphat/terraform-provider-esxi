#
#  See https://www.terraform.io/intro/getting-started/variables.html for more details.
#

<<<<<<< HEAD
=======
#  Change these defaults to fit your needs!

>>>>>>> a09975692ab4114aef08427f9b410b63842981c3
variable "esxi_hostname" { default = "esxi" }
variable "esxi_hostport" { default = "22" }
variable "esxi_username" { default = "root" }
variable "esxi_password" { } # Unspecified will prompt

variable "disk_store"    { default = "DS_001" }
