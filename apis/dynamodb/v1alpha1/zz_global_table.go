/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GlobalTableParameters defines the desired state of GlobalTable
type GlobalTableParameters struct {
	// Region is which region the GlobalTable will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The Regions where the global table needs to be created.
	// +kubebuilder:validation:Required
	ReplicationGroup            []*Replica `json:"replicationGroup"`
	CustomGlobalTableParameters `json:",inline"`
}

// GlobalTableSpec defines the desired state of GlobalTable
type GlobalTableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlobalTableParameters `json:"forProvider"`
}

// GlobalTableObservation defines the observed state of GlobalTable
type GlobalTableObservation struct {
	// The creation time of the global table.
	CreationDateTime *metav1.Time `json:"creationDateTime,omitempty"`
	// The unique identifier of the global table.
	GlobalTableARN *string `json:"globalTableARN,omitempty"`
	// The global table name.
	GlobalTableName *string `json:"globalTableName,omitempty"`
	// The current state of the global table:
	//
	//    * CREATING - The global table is being created.
	//
	//    * UPDATING - The global table is being updated.
	//
	//    * DELETING - The global table is being deleted.
	//
	//    * ACTIVE - The global table is ready for use.
	GlobalTableStatus *string `json:"globalTableStatus,omitempty"`
}

// GlobalTableStatus defines the observed state of GlobalTable.
type GlobalTableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlobalTableObservation `json:"atProvider"`
}

// +kubebuilder:object:root=true

// GlobalTable is the Schema for the GlobalTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GlobalTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalTableSpec   `json:"spec,omitempty"`
	Status            GlobalTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalTableList contains a list of GlobalTables
type GlobalTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalTable `json:"items"`
}

// Repository type metadata.
var (
	GlobalTableKind             = "GlobalTable"
	GlobalTableGroupKind        = schema.GroupKind{Group: Group, Kind: GlobalTableKind}.String()
	GlobalTableKindAPIVersion   = GlobalTableKind + "." + GroupVersion.String()
	GlobalTableGroupVersionKind = GroupVersion.WithKind(GlobalTableKind)
)

func init() {
	SchemeBuilder.Register(&GlobalTable{}, &GlobalTableList{})
}