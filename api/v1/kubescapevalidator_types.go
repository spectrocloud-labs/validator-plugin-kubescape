/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubescapeValidatorSpec defines the desired state of KubescapeValidator
type KubescapeValidatorSpec struct {
	SeverityLimitRules []SeverityLimitRule `json:"severityLimitRules,omitempty" yaml:"severityLimitRules,omitempty"`
	// Ignore CVEs
	IgnoredVulnerabilities []string `json:"ignoredVulnerabilities,omitempty" yaml:"ignoredVulnerabilities,omitempty"`
}

// Increase for every rule
func (s KubescapeValidatorSpec) ResultCount() int {
	return len(s.SeverityLimitRules)
}

type SeverityLimitRule struct {
	RuleName   string `json:"name"`
	Critical   int    `json:"critical,omitempty"`
	High       int    `json:"high,omitempty"`
	Medium     int    `json:"medium,omitempty"`
	Low        int    `json:"low,omitempty"`
	Negligible int    `json:"negligible,omitempty"`
	Unknown    int    `json:"unknown,omitempty"`
}

func (r SeverityLimitRule) Name() string {
	return r.RuleName
}

// KubescapeValidatorStatus defines the observed state of KubescapeValidator
type KubescapeValidatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KubescapeValidator is the Schema for the kubescapevalidators API
type KubescapeValidator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubescapeValidatorSpec   `json:"spec,omitempty"`
	Status KubescapeValidatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KubescapeValidatorList contains a list of KubescapeValidator
type KubescapeValidatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubescapeValidator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubescapeValidator{}, &KubescapeValidatorList{})
}
