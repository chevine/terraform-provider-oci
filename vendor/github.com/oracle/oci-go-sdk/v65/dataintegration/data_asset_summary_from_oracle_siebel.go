// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataAssetSummaryFromOracleSiebel Summary details for the Oracle Siebel data asset type.
type DataAssetSummaryFromOracleSiebel struct {

	// The Oracle Siebel hostname.
	Host *string `mandatory:"true" json:"host"`

	// The Oracle Siebel port.
	Port *string `mandatory:"true" json:"port"`

	DefaultConnection *ConnectionSummaryFromOracleSiebel `mandatory:"true" json:"defaultConnection"`

	// Generated key that can be used in API calls to identify data asset.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The user-defined description of the data asset.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Additional properties for the data asset.
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	NativeTypeSystem *TypeSystem `mandatory:"false" json:"nativeTypeSystem"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The Oracle Siebel service name.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The Oracle Siebel driver class.
	DriverClass *string `mandatory:"false" json:"driverClass"`

	// The Oracle Siebel SID.
	Sid *string `mandatory:"false" json:"sid"`

	WalletSecret *SensitiveAttribute `mandatory:"false" json:"walletSecret"`

	WalletPasswordSecret *SensitiveAttribute `mandatory:"false" json:"walletPasswordSecret"`
}

// GetKey returns Key
func (m DataAssetSummaryFromOracleSiebel) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m DataAssetSummaryFromOracleSiebel) GetModelVersion() *string {
	return m.ModelVersion
}

// GetName returns Name
func (m DataAssetSummaryFromOracleSiebel) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m DataAssetSummaryFromOracleSiebel) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m DataAssetSummaryFromOracleSiebel) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m DataAssetSummaryFromOracleSiebel) GetIdentifier() *string {
	return m.Identifier
}

// GetExternalKey returns ExternalKey
func (m DataAssetSummaryFromOracleSiebel) GetExternalKey() *string {
	return m.ExternalKey
}

// GetAssetProperties returns AssetProperties
func (m DataAssetSummaryFromOracleSiebel) GetAssetProperties() map[string]string {
	return m.AssetProperties
}

// GetNativeTypeSystem returns NativeTypeSystem
func (m DataAssetSummaryFromOracleSiebel) GetNativeTypeSystem() *TypeSystem {
	return m.NativeTypeSystem
}

// GetObjectVersion returns ObjectVersion
func (m DataAssetSummaryFromOracleSiebel) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetParentRef returns ParentRef
func (m DataAssetSummaryFromOracleSiebel) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetMetadata returns Metadata
func (m DataAssetSummaryFromOracleSiebel) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DataAssetSummaryFromOracleSiebel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataAssetSummaryFromOracleSiebel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataAssetSummaryFromOracleSiebel) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataAssetSummaryFromOracleSiebel DataAssetSummaryFromOracleSiebel
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataAssetSummaryFromOracleSiebel
	}{
		"ORACLE_SIEBEL_DATA_ASSET",
		(MarshalTypeDataAssetSummaryFromOracleSiebel)(m),
	}

	return json.Marshal(&s)
}
