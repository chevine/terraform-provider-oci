// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetAttributeTagRequest wrapper for the GetAttributeTag operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetAttributeTag.go.html to see an example of how to use GetAttributeTagRequest.
type GetAttributeTagRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Unique attribute key.
	AttributeKey *string `mandatory:"true" contributesTo:"path" name:"attributeKey"`

	// Unique tag key.
	TagKey *string `mandatory:"true" contributesTo:"path" name:"tagKey"`

	// Specifies the fields to return in an entity attribute tag response.
	Fields []GetAttributeTagFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAttributeTagRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAttributeTagRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAttributeTagRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAttributeTagRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAttributeTagRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetAttributeTagFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetAttributeTagFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAttributeTagResponse wrapper for the GetAttributeTag operation
type GetAttributeTagResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AttributeTag instance
	AttributeTag `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAttributeTagResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAttributeTagResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAttributeTagFieldsEnum Enum with underlying type: string
type GetAttributeTagFieldsEnum string

// Set of constants representing the allowable values for GetAttributeTagFieldsEnum
const (
	GetAttributeTagFieldsKey             GetAttributeTagFieldsEnum = "key"
	GetAttributeTagFieldsName            GetAttributeTagFieldsEnum = "name"
	GetAttributeTagFieldsTermkey         GetAttributeTagFieldsEnum = "termKey"
	GetAttributeTagFieldsTermpath        GetAttributeTagFieldsEnum = "termPath"
	GetAttributeTagFieldsTermdescription GetAttributeTagFieldsEnum = "termDescription"
	GetAttributeTagFieldsLifecyclestate  GetAttributeTagFieldsEnum = "lifecycleState"
	GetAttributeTagFieldsTimecreated     GetAttributeTagFieldsEnum = "timeCreated"
	GetAttributeTagFieldsCreatedbyid     GetAttributeTagFieldsEnum = "createdById"
	GetAttributeTagFieldsUri             GetAttributeTagFieldsEnum = "uri"
	GetAttributeTagFieldsAttributekey    GetAttributeTagFieldsEnum = "attributeKey"
)

var mappingGetAttributeTagFieldsEnum = map[string]GetAttributeTagFieldsEnum{
	"key":             GetAttributeTagFieldsKey,
	"name":            GetAttributeTagFieldsName,
	"termKey":         GetAttributeTagFieldsTermkey,
	"termPath":        GetAttributeTagFieldsTermpath,
	"termDescription": GetAttributeTagFieldsTermdescription,
	"lifecycleState":  GetAttributeTagFieldsLifecyclestate,
	"timeCreated":     GetAttributeTagFieldsTimecreated,
	"createdById":     GetAttributeTagFieldsCreatedbyid,
	"uri":             GetAttributeTagFieldsUri,
	"attributeKey":    GetAttributeTagFieldsAttributekey,
}

var mappingGetAttributeTagFieldsEnumLowerCase = map[string]GetAttributeTagFieldsEnum{
	"key":             GetAttributeTagFieldsKey,
	"name":            GetAttributeTagFieldsName,
	"termkey":         GetAttributeTagFieldsTermkey,
	"termpath":        GetAttributeTagFieldsTermpath,
	"termdescription": GetAttributeTagFieldsTermdescription,
	"lifecyclestate":  GetAttributeTagFieldsLifecyclestate,
	"timecreated":     GetAttributeTagFieldsTimecreated,
	"createdbyid":     GetAttributeTagFieldsCreatedbyid,
	"uri":             GetAttributeTagFieldsUri,
	"attributekey":    GetAttributeTagFieldsAttributekey,
}

// GetGetAttributeTagFieldsEnumValues Enumerates the set of values for GetAttributeTagFieldsEnum
func GetGetAttributeTagFieldsEnumValues() []GetAttributeTagFieldsEnum {
	values := make([]GetAttributeTagFieldsEnum, 0)
	for _, v := range mappingGetAttributeTagFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAttributeTagFieldsEnumStringValues Enumerates the set of values in String for GetAttributeTagFieldsEnum
func GetGetAttributeTagFieldsEnumStringValues() []string {
	return []string{
		"key",
		"name",
		"termKey",
		"termPath",
		"termDescription",
		"lifecycleState",
		"timeCreated",
		"createdById",
		"uri",
		"attributeKey",
	}
}

// GetMappingGetAttributeTagFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAttributeTagFieldsEnum(val string) (GetAttributeTagFieldsEnum, bool) {
	enum, ok := mappingGetAttributeTagFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
