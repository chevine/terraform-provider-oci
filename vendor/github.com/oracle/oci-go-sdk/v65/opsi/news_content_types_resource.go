// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// NewsContentTypesResourceEnum Enum with underlying type: string
type NewsContentTypesResourceEnum string

// Set of constants representing the allowable values for NewsContentTypesResourceEnum
const (
	NewsContentTypesResourceHost     NewsContentTypesResourceEnum = "HOST"
	NewsContentTypesResourceDatabase NewsContentTypesResourceEnum = "DATABASE"
	NewsContentTypesResourceExadata  NewsContentTypesResourceEnum = "EXADATA"
)

var mappingNewsContentTypesResourceEnum = map[string]NewsContentTypesResourceEnum{
	"HOST":     NewsContentTypesResourceHost,
	"DATABASE": NewsContentTypesResourceDatabase,
	"EXADATA":  NewsContentTypesResourceExadata,
}

var mappingNewsContentTypesResourceEnumLowerCase = map[string]NewsContentTypesResourceEnum{
	"host":     NewsContentTypesResourceHost,
	"database": NewsContentTypesResourceDatabase,
	"exadata":  NewsContentTypesResourceExadata,
}

// GetNewsContentTypesResourceEnumValues Enumerates the set of values for NewsContentTypesResourceEnum
func GetNewsContentTypesResourceEnumValues() []NewsContentTypesResourceEnum {
	values := make([]NewsContentTypesResourceEnum, 0)
	for _, v := range mappingNewsContentTypesResourceEnum {
		values = append(values, v)
	}
	return values
}

// GetNewsContentTypesResourceEnumStringValues Enumerates the set of values in String for NewsContentTypesResourceEnum
func GetNewsContentTypesResourceEnumStringValues() []string {
	return []string{
		"HOST",
		"DATABASE",
		"EXADATA",
	}
}

// GetMappingNewsContentTypesResourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNewsContentTypesResourceEnum(val string) (NewsContentTypesResourceEnum, bool) {
	enum, ok := mappingNewsContentTypesResourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
