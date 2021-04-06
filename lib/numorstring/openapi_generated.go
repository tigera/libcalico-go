// +build !ignore_autogenerated

// Copyright (c) 2016-2017 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package numorstring

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/projectcalico/libcalico-go/lib/numorstring.Port":          schema_projectcalico_libcalico_go_lib_numorstring_Port(ref),
		"github.com/projectcalico/libcalico-go/lib/numorstring.Protocol":      schema_projectcalico_libcalico_go_lib_numorstring_Protocol(ref),
		"github.com/projectcalico/libcalico-go/lib/numorstring.Uint8OrString": schema_projectcalico_libcalico_go_lib_numorstring_Uint8OrString(ref),
	}
}

func schema_projectcalico_libcalico_go_lib_numorstring_Port(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Port represents either a range of numeric ports or a named port.\n\n    - For a named port, set the PortName, leaving MinPort and MaxPort as 0.\n    - For a port range, set MinPort and MaxPort to the (inclusive) port numbers.  Set\n      PortName to \"\".\n    - For a single port, set MinPort = MaxPort and PortName = \"\".",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"minPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"maxPort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"portName": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"portName"},
			},
		},
	}
}

func schema_projectcalico_libcalico_go_lib_numorstring_Protocol(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"numVal": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "byte",
						},
					},
					"strVal": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"type", "numVal", "strVal"},
			},
		},
	}
}

func schema_projectcalico_libcalico_go_lib_numorstring_Uint8OrString(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UInt8OrString is a type that can hold an uint8 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"numVal": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "byte",
						},
					},
					"strVal": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"type", "numVal", "strVal"},
			},
		},
	}
}
