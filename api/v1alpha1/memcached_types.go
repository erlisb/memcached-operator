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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// MemcachedSpec defines the desired state of Memcached
type MemcachedSpec struct {

	// Size defines the number of Memcached instances
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Size int32 `json:"size,omitempty"`
	// MaxSurge is an optional field that specifies the maximum number of Pods that can be created over the desired number of Pods.
	// The value can be an absolute number (for example, 5) or a percentage of desired Pods (for example, 10%).
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	MaxSurge intstr.IntOrString `json:"maxsurge,omitempty"`
	// MaxUnavailable is an optional field that specifies the maximum number of Pods that can be unavailable during the update process.
	// The value can be an absolute number (for example, 5) or a percentage of desired Pods (for example, 10%).
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	MaxUnavailable intstr.IntOrString `json:"maxunavailable,omitempty"`

	// Image specifies Docker container image. This field is optional to allow higher level config management to default or override
	// container images in workload controllers like Deployments and StatefulSets.
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Image string `json:"image"`

	// The maximum time in seconds for a deployment to make progress before it is considered to be failed.
	// The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason
	// will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused.
	// Defaults to 600s.
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ProgressDeadlineSeconds *int32 `json:"progressDeadlineSeconds,omitempty" protobuf:"varint,9,opt,name=progressDeadlineSeconds"`
}

// MemcachedStatus defines the observed state of Memcached
type MemcachedStatus struct {
	// Nodes store the name of the pods which are running Memcached instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Nodes []string `json:"nodes,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +operator-sdk:csv:customresourcedefinitions:resources={{Deployment,v1,memcached-deployment}}

// Memcached is the Schema for the memcacheds API
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MemcachedList contains a list of Memcached
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}
