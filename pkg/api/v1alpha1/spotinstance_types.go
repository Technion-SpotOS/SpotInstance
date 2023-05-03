/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SpotInstanceSpec defines the desired state of SpotInstance.
type SpotInstanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Region is the region of the instance.
	Region string `json:"region,omitempty"`
	// CPU is the number of CPU cores in the instance.
	CPU int32 `json:"cpu,omitempty"`
	// Memory is the amount of memory in the instance.
	Memory int32 `json:"memory,omitempty"`
	// TypeName is the type of the instance.
	TypeName string `json:"typeName,omitempty"`
	// TypeMajor is the major type of the instance.
	TypeMajor string `json:"typeMajor,omitempty"`
	// TypeMinor is the minor type of the instance.
	TypeMinor string `json:"typeMinor,omitempty"`
	// Storage is the storage of the instance.
	Storage string `json:"storage,omitempty"`
}

// SpotInstanceStatus defines the observed state of SpotInstance.
type SpotInstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Stage is the stage of the lifecycle of the instance
	Stage string `json:"stage,omitempty" enum:"PendingOrder|Ordered|Ready|Installed"`
	// NodeName is the name of the node assigned to the instance
	NodeName string `json:"nodeName,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SpotInstance is the Schema for the spotinstances API.
type SpotInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpotInstanceSpec   `json:"spec,omitempty"`
	Status SpotInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpotInstanceList contains a list of SpotInstance.
type SpotInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpotInstance{}, &SpotInstanceList{})
}
