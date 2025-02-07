/*
Copyright 2021 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"cuelang.org/go/cue"
)

// Source record the source of Capability
type Source struct {
	RepoName  string `json:"repoName"`
	ChartName string `json:"chartName,omitempty"`
}

// CRDInfo record the CRD info of the Capability
type CRDInfo struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

// Chart defines all necessary information to install a whole chart
type Chart struct {
	Repo      string                 `json:"repo"`
	URL       string                 `json:"url"`
	Name      string                 `json:"name"`
	Namespace string                 `json:"namespace,omitempty"`
	Version   string                 `json:"version"`
	Values    map[string]interface{} `json:"values"`
}

// Installation defines the installation method for this Capability, currently only helm is supported
type Installation struct {
	Helm Chart `json:"helm"`
	// TODO(wonderflow) add raw yaml file support for install capability
}

// CapType defines the type of capability
type CapType string

const (
	// TypeComponentDefinition represents OAM ComponentDefinition
	TypeComponentDefinition CapType = "componentDefinition"
	// TypeWorkload represents OAM Workload
	TypeWorkload CapType = "workload"
	// TypeTrait represents OAM Trait
	TypeTrait CapType = "trait"
	// TypeWorkflowStep represent OAM Workflow
	TypeWorkflowStep CapType = "workflowstep"
	// TypePolicy represent OAM Policy
	TypePolicy CapType = "policy"
)

// CapabilityConfigMapNamePrefix is the prefix for capability ConfigMap name
const CapabilityConfigMapNamePrefix = "schema-"

const (
	// OpenapiV3JSONSchema is the key to store OpenAPI v3 JSON schema in ConfigMap
	OpenapiV3JSONSchema string = "openapi-v3-json-schema"
	// UISchema is the key to store ui custom schema
	UISchema string = "ui-schema"
	// VelaQLConfigmapKey is the key to store velaql view
	VelaQLConfigmapKey string = "template"
)

// CapabilityCategory defines the category of a capability
type CapabilityCategory string

// categories of capability schematic
const (
	TerraformCategory CapabilityCategory = "terraform"

	CUECategory CapabilityCategory = "cue"
)

// Parameter defines a parameter for cli from capability template
type Parameter struct {
	Name     string      `json:"name"`
	Short    string      `json:"short,omitempty"`
	Required bool        `json:"required,omitempty"`
	Default  interface{} `json:"default,omitempty"`
	Usage    string      `json:"usage,omitempty"`
	Ignore   bool        `json:"ignore,omitempty"`
	Type     cue.Kind    `json:"type,omitempty"`
	Alias    string      `json:"alias,omitempty"`
	JSONType string      `json:"jsonType,omitempty"`
}

// Capability defines the content of a capability
type Capability struct {
	Name           string             `json:"name"`
	Type           CapType            `json:"type"`
	CueTemplate    string             `json:"template,omitempty"`
	CueTemplateURI string             `json:"templateURI,omitempty"`
	Parameters     []Parameter        `json:"parameters,omitempty"`
	CrdName        string             `json:"crdName,omitempty"`
	Center         string             `json:"center,omitempty"`
	Status         string             `json:"status,omitempty"`
	Description    string             `json:"description,omitempty"`
	Example        string             `json:"example,omitempty"`
	Labels         map[string]string  `json:"labels,omitempty"`
	Category       CapabilityCategory `json:"category,omitempty"`

	// trait only
	AppliesTo []string `json:"appliesTo,omitempty"`

	// Namespace represents it's a system-level or user-level capability.
	Namespace string `json:"namespace,omitempty"`

	// Plugin Source
	Source  *Source  `json:"source,omitempty"`
	CrdInfo *CRDInfo `json:"crdInfo,omitempty"`

	// Terraform
	TerraformConfiguration string `json:"terraformConfiguration,omitempty"`
	ConfigurationType      string `json:"configurationType,omitempty"`
	Path                   string `json:"path,omitempty"`
}
