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

type CatalogSourceBlueprintObservation struct {
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Global *bool `json:"global,omitempty" tf:"global,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ItemsFound *string `json:"itemsFound,omitempty" tf:"items_found,omitempty"`

	ItemsImported *string `json:"itemsImported,omitempty" tf:"items_imported,omitempty"`

	LastImportCompletedAt *string `json:"lastImportCompletedAt,omitempty" tf:"last_import_completed_at,omitempty"`

	LastImportErrors []*string `json:"lastImportErrors,omitempty" tf:"last_import_errors,omitempty"`

	LastImportStartedAt *string `json:"lastImportStartedAt,omitempty" tf:"last_import_started_at,omitempty"`

	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty" tf:"last_updated_by,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	TypeID *string `json:"typeId,omitempty" tf:"type_id,omitempty"`
}

type CatalogSourceBlueprintParameters struct {

	// +kubebuilder:validation:Optional
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/ankasoftco/provider-vra/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

// CatalogSourceBlueprintSpec defines the desired state of CatalogSourceBlueprint
type CatalogSourceBlueprintSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogSourceBlueprintParameters `json:"forProvider"`
}

// CatalogSourceBlueprintStatus defines the observed state of CatalogSourceBlueprint.
type CatalogSourceBlueprintStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogSourceBlueprintObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogSourceBlueprint is the Schema for the CatalogSourceBlueprints API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type CatalogSourceBlueprint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   CatalogSourceBlueprintSpec   `json:"spec"`
	Status CatalogSourceBlueprintStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogSourceBlueprintList contains a list of CatalogSourceBlueprints
type CatalogSourceBlueprintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogSourceBlueprint `json:"items"`
}

// Repository type metadata.
var (
	CatalogSourceBlueprint_Kind             = "CatalogSourceBlueprint"
	CatalogSourceBlueprint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogSourceBlueprint_Kind}.String()
	CatalogSourceBlueprint_KindAPIVersion   = CatalogSourceBlueprint_Kind + "." + CRDGroupVersion.String()
	CatalogSourceBlueprint_GroupVersionKind = CRDGroupVersion.WithKind(CatalogSourceBlueprint_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogSourceBlueprint{}, &CatalogSourceBlueprintList{})
}
