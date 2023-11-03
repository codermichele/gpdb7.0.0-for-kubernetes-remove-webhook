/*
.
*/

// Package v1 contains API Schema definitions for the greenplum v1 API group
// +kubebuilder:object:generate=true
// +groupName=greenplum.pivotal.io
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "greenplum.pivotal.io", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

const (
	// AppName is the generic name for our application (for `app=AppName` labels)
	AppName = "greenplum"
)
