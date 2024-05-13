/*
Copyright 2022.

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

package v1beta1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GrafanaContactPointSpec defines the desired state of GrafanaContactPoint
type GrafanaContactPointSpec struct {
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=duration
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$"
	// +kubebuilder:default="10m"
	ResyncPeriod metav1.Duration `json:"resyncPeriod,omitempty"`

	// selects Grafanas for import
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	InstanceSelector *metav1.LabelSelector `json:"instanceSelector"`

	// +optional
	DisableResolveMessage bool `json:"disableResolveMessage,omitempty"`

	// +kubebuilder:validation:type=string
	Name string `json:"name"`

	Settings *apiextensions.JSON `json:"settings"`

	// +kubebuilder:validation:Enum=alertmanager;dingding;discord;email;googlechat;kafka;line;opsgenie;pagerduty;pushover;sensugo;slack;teams;telegram;threema;victorops;webhook;wecom
	Type string `json:"type,omitempty"`

	// +optional
	AllowCrossNamespaceImport *bool `json:"allowCrossNamespaceImport,omitempty"`
}

// GrafanaContactPointStatus defines the observed state of GrafanaContactPoint
type GrafanaContactPointStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GrafanaContactPoint is the Schema for the grafanacontactpoints API
type GrafanaContactPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GrafanaContactPointSpec   `json:"spec,omitempty"`
	Status GrafanaContactPointStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GrafanaContactPointList contains a list of GrafanaContactPoint
type GrafanaContactPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafanaContactPoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GrafanaContactPoint{}, &GrafanaContactPointList{})
}
