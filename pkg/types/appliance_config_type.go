package types

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApplianceConfigVersion is the version supported by this package.
const ApplianceConfigVersion = "v1beta1"

type ApplianceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	OcpReleaseImage string
	DiskSizeGB      int
	PullSecret      string
}
