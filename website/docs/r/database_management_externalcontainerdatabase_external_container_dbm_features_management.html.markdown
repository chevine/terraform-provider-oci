---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_externalcontainerdatabase_external_container_dbm_features_management"
sidebar_current: "docs-oci-resource-database_management-externalcontainerdatabase_external_container_dbm_features_management"
description: |-
  Provides the Externalcontainerdatabase External Container Dbm Features Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_externalcontainerdatabase_external_container_dbm_features_management
This resource provides the Externalcontainerdatabase External Container Dbm Features Management resource in Oracle Cloud Infrastructure Database Management service.

Enables a Database Management feature for the specified external container database.


## Example Usage

```hcl
resource "oci_database_management_externalcontainerdatabase_external_container_dbm_features_management" "test_externalcontainerdatabase_external_container_dbm_features_management" {
	#Required
	external_container_database_id = oci_database_external_container_database.test_external_container_database.id
	enable_external_container_dbm_feature = var.enable_external_container_dbm_feature

	#Optional
	feature_details {
		#Required
		feature = var.externalcontainerdatabase_external_container_dbm_features_management_feature_details_feature
		enable_external_container_dbm_feature = var.enable_external_container_dbm_feature

		#Optional
		connector_details {

			#Optional
			connector_type = var.externalcontainerdatabase_external_container_dbm_features_management_feature_details_connector_details_connector_type
			database_connector_id = oci_database_management_database_connector.test_database_connector.id
			management_agent_id = oci_management_agent_management_agent.test_management_agent.id
			private_end_point_id = oci_database_management_private_end_point.test_private_end_point.id
		}
		can_enable_all_current_pdbs = var.externalcontainerdatabase_external_container_dbm_features_management_feature_details_can_enable_all_current_pdbs
		is_auto_enable_pluggable_database = var.externalcontainerdatabase_external_container_dbm_features_management_feature_details_is_auto_enable_pluggable_database
		license_model = var.externalcontainerdatabase_external_container_dbm_features_management_feature_details_license_model
	}
}
```

## Argument Reference

The following arguments are supported:

* `external_container_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external container database.
* `feature_details` - (Optional) The details required to enable the specified Database Management feature.
	* `can_enable_all_current_pdbs` - (Applicable when feature=DIAGNOSTICS_AND_MANAGEMENT) Indicates whether Diagnostics & Management should be enabled for all the current pluggable databases in the container database.
	* `connector_details` - (Optional) The connector details required to connect to an Oracle cloud database.
		* `connector_type` - (Optional) The list of supported connection types:
			* PE: Private endpoint
			* MACS: Management agent
			* EXTERNAL: External database connector
			* DIRECT: Direct connection 
		* `database_connector_id` - (Applicable when connector_type=EXTERNAL) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database connector.
		* `management_agent_id` - (Applicable when connector_type=MACS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent.
		* `private_end_point_id` - (Applicable when connector_type=PE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
	* `feature` - (Required) The name of the Database Management feature.
	* `is_auto_enable_pluggable_database` - (Applicable when feature=DIAGNOSTICS_AND_MANAGEMENT) Indicates whether Diagnostics & Management should be enabled automatically for all the pluggable databases in the container database.
	* `license_model` - (Applicable when feature=DB_LIFECYCLE_MANAGEMENT | DIAGNOSTICS_AND_MANAGEMENT) The Oracle license model that applies to the external database. 
* `enable_external_container_dbm_feature` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Externalcontainerdatabase External Container Dbm Features Management
	* `update` - (Defaults to 20 minutes), when updating the Externalcontainerdatabase External Container Dbm Features Management
	* `delete` - (Defaults to 20 minutes), when destroying the Externalcontainerdatabase External Container Dbm Features Management
