// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "subscription_id" {

}

variable "compartment_id" {

}

provider "oci" {
	tenancy_ocid     = var.tenancy_ocid
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

data "oci_tenantmanagercontrolplane_subscription" "test_subscription" {
    subscription_id = var.subscription_id
}

data "oci_tenantmanagercontrolplane_subscriptions" "test_subscriptions" {
    compartment_id = var.compartment_id
}