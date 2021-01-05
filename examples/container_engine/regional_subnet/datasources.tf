// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domain" "ad1" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = var.tenancy_ocid
  ad_number      = 2
}

data "oci_containerengine_cluster_option" "test_cluster_option" {
  cluster_option_id = "all"
}

data "oci_containerengine_node_pool_option" "test_node_pool_option" {
  node_pool_option_id = "all"
}

locals {
  all_sources = "${data.oci_containerengine_node_pool_option.test_node_pool_option.sources}"

  oracle_linux_images = [for source in local.all_sources : source.image_id if length(regexall("Oracle-Linux-[0-9]*.[0-9]*-20[0-9]*",source.source_name)) > 0]
}

output "cluster_kubernetes_versions" {
  value = [data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions]
}

output "node_pool_kubernetes_version" {
  value = [data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions]
}

