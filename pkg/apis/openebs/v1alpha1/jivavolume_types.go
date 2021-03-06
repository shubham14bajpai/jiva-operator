/*
Copyright 2019 The OpenEBS Authors
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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ISCSISpec struct {
	TargetIP   string `json:"targetIP"`
	TargetPort int32  `json:"targetPort"`
	Iqn        string `json:"iqn"`
}

type MountInfo struct {
	Path       string `json:"mountPath"`
	FSType     string `json:"fsType"`
	DevicePath string `json:"devicePath"`
}

// JivaVolumeSpec defines the desired state of JivaVolume
// +k8s:openapi-gen=true
type JivaVolumeSpec struct {
	// ReplicaSC represents the storage class used for
	// creating the pvc for the replicas (provisioned by localpv provisioner)
	ReplicaSC string `json:"replicaSC"`
	PV        string `json:"pv"`
	Capacity  string `json:"capacity"`
	// ReplicationFactor represents the actual replica count for the underlying
	// jiva volume
	ReplicationFactor string                  `json:"replicationFactor"`
	ISCSISpec         ISCSISpec               `json:"iscsiSpec"`
	MountInfo         MountInfo               `json:"mountInfo"`
	ReplicaResource   v1.ResourceRequirements `json:"replicaResource"`
	TargetResource    v1.ResourceRequirements `json:"targetResource"`
}

// JivaVolumeStatus defines the observed state of JivaVolume
// +k8s:openapi-gen=true
type JivaVolumeStatus struct {
	Status          string          `json:"status"`
	ReplicaCount    int             `json:"replicaCount"`
	ReplicaStatuses []ReplicaStatus `json:"replicaStatus"`
	// Phase represents the current phase of JivaVolume.
	Phase JivaVolumePhase `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JivaVolume is the Schema for the jivavolumes API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type JivaVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JivaVolumeSpec   `json:"spec,omitempty"`
	Status JivaVolumeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JivaVolumeList contains a list of JivaVolume
type JivaVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JivaVolume `json:"items"`
}

// ReplicaStatus stores the status of replicas
type ReplicaStatus struct {
	Address string `json:"address"`
	Mode    string `json:"mode"`
}

// JivaVolumePhase represents the current phase of JivaVolume.
type JivaVolumePhase string

const (
	// JivaVolumePhasePending indicates that the jivavolume is still waiting for
	// the jivavolume to be created and bound
	JivaVolumePhasePending JivaVolumePhase = "Pending"

	// JivaVolumePhaseSyncing indicates that the jivavolume has been
	// provisioned and replicas are syncing
	JivaVolumePhaseSyncing JivaVolumePhase = "Syncing"

	// JivaVolumePhaseFailed indicates that the jivavolume provisioning
	// has failed
	JivaVolumePhaseFailed JivaVolumePhase = "Failed"

	// JivaVolumePhaseReady indicates that the jivavolume provisioning
	// has Created
	JivaVolumePhaseReady JivaVolumePhase = "Ready"

	// JivaVolumePhaseDeleting indicates the the jivavolume is deprovisioned
	JivaVolumePhaseDeleting JivaVolumePhase = "Deleting"
)

func init() {
	SchemeBuilder.Register(&JivaVolume{}, &JivaVolumeList{})
}
