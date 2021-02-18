variable "google-credentials-file" {
  type = string
  default = "credentials.json"
}

variable "ansible-playbook-file" {
  type = string
  default = "../ansible/rex.yml"
}

variable "vm-user" {
  type = string
  default = "rex"
}

variable "ssh-pub-key" {
  type = string
  default = "~/.ssh/id_rsa.pub"
}

variable "ssh-priv-key" {
  type = string
  default = "~/.ssh/id_rsa"
}

variable "project" {
  type = string
  default = "remote-exec-bot-305203"
}

variable "image-id" {
  type = string
  default = "us.gcr.io/remote-exec-bot-305203/rex:latest"
}
