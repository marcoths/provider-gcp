/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SSHPublicKeyObservation struct {

	// The SHA-256 fingerprint of the SSH public key.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// an identifier for the resource with format users/{{user}}/sshPublicKeys/{{fingerprint}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SSHPublicKeyParameters struct {

	// An expiration time in microseconds since epoch.
	// +kubebuilder:validation:Optional
	ExpirationTimeUsec *string `json:"expirationTimeUsec,omitempty" tf:"expiration_time_usec,omitempty"`

	// Public key text in SSH format, defined by RFC4253 section 6.6.
	// +kubebuilder:validation:Required
	KeySecretRef v1.SecretKeySelector `json:"keySecretRef" tf:"-"`

	// The project ID of the Google Cloud Platform project.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The user email.
	// +kubebuilder:validation:Required
	User *string `json:"user" tf:"user,omitempty"`
}

// SSHPublicKeySpec defines the desired state of SSHPublicKey
type SSHPublicKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSHPublicKeyParameters `json:"forProvider"`
}

// SSHPublicKeyStatus defines the observed state of SSHPublicKey.
type SSHPublicKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSHPublicKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SSHPublicKey is the Schema for the SSHPublicKeys API. The SSH public key information associated with a Google account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SSHPublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SSHPublicKeySpec   `json:"spec"`
	Status            SSHPublicKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSHPublicKeyList contains a list of SSHPublicKeys
type SSHPublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHPublicKey `json:"items"`
}

// Repository type metadata.
var (
	SSHPublicKey_Kind             = "SSHPublicKey"
	SSHPublicKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSHPublicKey_Kind}.String()
	SSHPublicKey_KindAPIVersion   = SSHPublicKey_Kind + "." + CRDGroupVersion.String()
	SSHPublicKey_GroupVersionKind = CRDGroupVersion.WithKind(SSHPublicKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SSHPublicKey{}, &SSHPublicKeyList{})
}
