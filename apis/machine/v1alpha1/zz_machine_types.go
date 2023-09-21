/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BootConfigObservation struct {

	// A valid cloud config data in json-escaped yaml syntax.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`
}

type BootConfigParameters struct {

	// A valid cloud config data in json-escaped yaml syntax.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`
}

type ConstraintsObservation struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`
}

type ConstraintsParameters struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	// +kubebuilder:validation:Required
	Mandatory *bool `json:"mandatory" tf:"mandatory,omitempty"`
}

type DisksListObservation struct {
	BlockDeviceID *string `json:"blockDeviceId,omitempty" tf:"block_device_id,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DisksListParameters struct {
}

type DisksObservation struct {

	// The id of the existing block device.
	BlockDeviceID *string `json:"blockDeviceId,omitempty" tf:"block_device_id,omitempty"`

	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A human-friendly block-device name used as an identifier in APIs that support this option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DisksParameters struct {

	// The id of the existing block device.
	// +kubebuilder:validation:Required
	BlockDeviceID *string `json:"blockDeviceId" tf:"block_device_id,omitempty"`

	// A human-friendly description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A human-friendly block-device name used as an identifier in APIs that support this option.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ImageDiskConstraintsObservation struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`
}

type ImageDiskConstraintsParameters struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// Indicates whether this constraint should be strictly enforced or not.
	// +kubebuilder:validation:Required
	Mandatory *bool `json:"mandatory" tf:"mandatory,omitempty"`
}

type LinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type LinksParameters struct {
}

type MachineObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Machine boot config that will be passed to the instance that can be used to perform common automated configuration tasks and even run scripts after the instance starts.
	BootConfig []BootConfigObservation `json:"bootConfig,omitempty" tf:"boot_config,omitempty"`

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	Constraints []ConstraintsObservation `json:"constraints,omitempty" tf:"constraints,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Additional custom properties that may be used to extend the machine.
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// The id of the deployment that is associated with this resource.
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specification for attaching/detaching disks to a machine.
	Disks []DisksObservation `json:"disks,omitempty" tf:"disks,omitempty"`

	// List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
	DisksList []DisksListObservation `json:"disksList,omitempty" tf:"disks_list,omitempty"`

	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	ExternalZoneID *string `json:"externalZoneId,omitempty" tf:"external_zone_id,omitempty"`

	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of image used for this machine.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	ImageDiskConstraints []ImageDiskConstraintsObservation `json:"imageDiskConstraints,omitempty" tf:"image_disk_constraints,omitempty"`

	ImageRef *string `json:"imageRef,omitempty" tf:"image_ref,omitempty"`

	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Nics []NicsObservation `json:"nics,omitempty" tf:"nics,omitempty"`

	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	PowerState *string `json:"powerState,omitempty" tf:"power_state,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type MachineParameters struct {

	// Machine boot config that will be passed to the instance that can be used to perform common automated configuration tasks and even run scripts after the instance starts.
	// +kubebuilder:validation:Optional
	BootConfig []BootConfigParameters `json:"bootConfig,omitempty" tf:"boot_config,omitempty"`

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	// +kubebuilder:validation:Optional
	Constraints []ConstraintsParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Additional custom properties that may be used to extend the machine.
	// +kubebuilder:validation:Optional
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// The id of the deployment that is associated with this resource.
	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specification for attaching/detaching disks to a machine.
	// +kubebuilder:validation:Optional
	Disks []DisksParameters `json:"disks,omitempty" tf:"disks,omitempty"`

	// +kubebuilder:validation:Optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor,omitempty"`

	// Type of image used for this machine.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Constraints that are used to drive placement policies for entities such as image, network, storage, etc. Constraint expressions are matched against tags on existing placement targets.
	// +kubebuilder:validation:Optional
	ImageDiskConstraints []ImageDiskConstraintsParameters `json:"imageDiskConstraints,omitempty" tf:"image_disk_constraints,omitempty"`

	// +kubebuilder:validation:Optional
	ImageRef *string `json:"imageRef,omitempty" tf:"image_ref,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Nics []NicsParameters `json:"nics,omitempty" tf:"nics,omitempty"`

	// +crossplane:generate:reference:type=github.com/ankasoftco/provider-vra/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NicsObservation struct {
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type NicsParameters struct {

	// +kubebuilder:validation:Optional
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// +kubebuilder:validation:Optional
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkID *string `json:"networkId" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type TagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// MachineSpec defines the desired state of Machine
type MachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MachineParameters `json:"forProvider"`
}

// MachineStatus defines the observed state of Machine.
type MachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Machine is the Schema for the Machines API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type Machine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.flavor)",message="flavor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   MachineSpec   `json:"spec"`
	Status MachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MachineList contains a list of Machines
type MachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Machine `json:"items"`
}

// Repository type metadata.
var (
	Machine_Kind             = "Machine"
	Machine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Machine_Kind}.String()
	Machine_KindAPIVersion   = Machine_Kind + "." + CRDGroupVersion.String()
	Machine_GroupVersionKind = CRDGroupVersion.WithKind(Machine_Kind)
)

func init() {
	SchemeBuilder.Register(&Machine{}, &MachineList{})
}
