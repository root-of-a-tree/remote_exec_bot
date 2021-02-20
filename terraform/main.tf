terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.57.0"
    }
  }
}

provider "google" {

  credentials = file(var.google-credentials-file)

  project = var.project
  region  = "us-central1"
  zone    = "us-central1-c"
}

resource "google_compute_instance" "vm_instance" {
  name         = "terraform-instance"
  machine_type = "f1-micro"
  allow_stopping_for_update = "true"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
    access_config {
    }
  }

  metadata = {
    ssh-keys = "${var.vm-user}:${file(var.ssh-pub-key)}"
  }
}

# separate provisioning steps into a null resource, so they get triggered every time the image version updates
resource "null_resource" "ansible-provisioner" {
  triggers = {
    image_id = var.image-id
  }

  provisioner "remote-exec" {
    inline = ["sudo apt update", "sudo apt install python3 -y", "echo Done!"]

    connection {
      host        = google_compute_instance.vm_instance.network_interface.0.access_config.0.nat_ip
      type        = "ssh"
      user        = var.vm-user
      private_key = file(var.ssh-priv-key)
    }
  }

  provisioner "local-exec" {
    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u ${var.vm-user} -i '${google_compute_instance.vm_instance.network_interface.0.access_config.0.nat_ip},' --private-key ${var.ssh-priv-key} -e 'pub_key=${var.ssh-pub-key}' ${var.ansible-playbook-file} --extra-vars '{\"imageID\":\"${var.image-id}\"}'"
  }
}

output "vm_ip_addresses" {
  value = {
    (google_compute_instance.vm_instance.name) = google_compute_instance.vm_instance.network_interface.0.access_config.0.nat_ip
  }
}
