// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MemberInitParameters struct {

	// (String) The Room ID of the space
	// The Room ID of the space
	// +crossplane:generate:reference:type=github.com/estenrye/provider-netdata/apis/room/v1alpha1.Room
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RoomID *string `json:"roomId,omitempty" tf:"room_id,omitempty"`

	// Reference to a Room in room to populate roomId.
	// +kubebuilder:validation:Optional
	RoomIDRef *v1.Reference `json:"roomIdRef,omitempty" tf:"-"`

	// Selector for a Room in room to populate roomId.
	// +kubebuilder:validation:Optional
	RoomIDSelector *v1.Selector `json:"roomIdSelector,omitempty" tf:"-"`

	// (String) Space ID of the member
	// Space ID of the member
	// +crossplane:generate:reference:type=github.com/estenrye/provider-netdata/apis/space/v1alpha1.Space
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SpaceID *string `json:"spaceId,omitempty" tf:"space_id,omitempty"`

	// Reference to a Space in space to populate spaceId.
	// +kubebuilder:validation:Optional
	SpaceIDRef *v1.Reference `json:"spaceIdRef,omitempty" tf:"-"`

	// Selector for a Space in space to populate spaceId.
	// +kubebuilder:validation:Optional
	SpaceIDSelector *v1.Selector `json:"spaceIdSelector,omitempty" tf:"-"`

	// (String) The Space Member ID of the space
	// The Space Member ID of the space
	// +crossplane:generate:reference:type=github.com/estenrye/provider-netdata/apis/member/v1alpha1.Member
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SpaceMemberID *string `json:"spaceMemberId,omitempty" tf:"space_member_id,omitempty"`

	// Reference to a Member in member to populate spaceMemberId.
	// +kubebuilder:validation:Optional
	SpaceMemberIDRef *v1.Reference `json:"spaceMemberIdRef,omitempty" tf:"-"`

	// Selector for a Member in member to populate spaceMemberId.
	// +kubebuilder:validation:Optional
	SpaceMemberIDSelector *v1.Selector `json:"spaceMemberIdSelector,omitempty" tf:"-"`
}

type MemberObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The Room ID of the space
	// The Room ID of the space
	RoomID *string `json:"roomId,omitempty" tf:"room_id,omitempty"`

	// (String) Space ID of the member
	// Space ID of the member
	SpaceID *string `json:"spaceId,omitempty" tf:"space_id,omitempty"`

	// (String) The Space Member ID of the space
	// The Space Member ID of the space
	SpaceMemberID *string `json:"spaceMemberId,omitempty" tf:"space_member_id,omitempty"`
}

type MemberParameters struct {

	// (String) The Room ID of the space
	// The Room ID of the space
	// +crossplane:generate:reference:type=github.com/estenrye/provider-netdata/apis/room/v1alpha1.Room
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RoomID *string `json:"roomId,omitempty" tf:"room_id,omitempty"`

	// Reference to a Room in room to populate roomId.
	// +kubebuilder:validation:Optional
	RoomIDRef *v1.Reference `json:"roomIdRef,omitempty" tf:"-"`

	// Selector for a Room in room to populate roomId.
	// +kubebuilder:validation:Optional
	RoomIDSelector *v1.Selector `json:"roomIdSelector,omitempty" tf:"-"`

	// (String) Space ID of the member
	// Space ID of the member
	// +crossplane:generate:reference:type=github.com/estenrye/provider-netdata/apis/space/v1alpha1.Space
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpaceID *string `json:"spaceId,omitempty" tf:"space_id,omitempty"`

	// Reference to a Space in space to populate spaceId.
	// +kubebuilder:validation:Optional
	SpaceIDRef *v1.Reference `json:"spaceIdRef,omitempty" tf:"-"`

	// Selector for a Space in space to populate spaceId.
	// +kubebuilder:validation:Optional
	SpaceIDSelector *v1.Selector `json:"spaceIdSelector,omitempty" tf:"-"`

	// (String) The Space Member ID of the space
	// The Space Member ID of the space
	// +crossplane:generate:reference:type=github.com/estenrye/provider-netdata/apis/member/v1alpha1.Member
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpaceMemberID *string `json:"spaceMemberId,omitempty" tf:"space_member_id,omitempty"`

	// Reference to a Member in member to populate spaceMemberId.
	// +kubebuilder:validation:Optional
	SpaceMemberIDRef *v1.Reference `json:"spaceMemberIdRef,omitempty" tf:"-"`

	// Selector for a Member in member to populate spaceMemberId.
	// +kubebuilder:validation:Optional
	SpaceMemberIDSelector *v1.Selector `json:"spaceMemberIdSelector,omitempty" tf:"-"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MemberInitParameters `json:"initProvider,omitempty"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Member is the Schema for the Members API. Provides a Netdata Cloud Room Member resource. Use this resource to manage user membership to the room in the selected space. It is referring to the user created at the space level.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netdata}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MemberSpec   `json:"spec"`
	Status            MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}
