# terraform-provider-nomad

Get node id for a machine running the nomad agent in client mode

Usage:

    resource "nomad_node" "node" {
      count = "${var.cluster_size}"
      node_addr = "http://${element(aws_instance.coreos.*.public_ip, count.index)}:4646"
    }

    output "Nomad_Nodes" {
      value = "${formatlist("%v", nomad_node.node.*.id)}"
    }

