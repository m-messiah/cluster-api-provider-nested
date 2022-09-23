/*
Copyright 2022 The Kubernetes Authors.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/client.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=vcn
// +k8s:openapi-gen=true

// VirtualClusterNodeSelector is the Schema for the virtualclusters nodeselector isolation mechanism
// +k8s:openapi-gen=true
type VirtualClusterNodeSelector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VirtualClusterNodeSelectorSpec `json:"spec,omitempty"`
}

// VirtualClusterNodeSelectorSpec defines the desired selectors of VirtualClusterNodeSelector
type VirtualClusterNodeSelectorSpec struct {
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// If specified, list of taints that should be tolerated in pod spec
	// to assign to NoSchedule nodes for better isolation
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/client.Object

// VirtualClusterNodeSelectorList contains a list of VirtualClusterNodeSelector
type VirtualClusterNodeSelectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualClusterNodeSelector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualClusterNodeSelector{}, &VirtualClusterNodeSelectorList{})
}
