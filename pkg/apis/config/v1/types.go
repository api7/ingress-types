package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *ApisixRouteSpec `json:"spec,omitempty"`
}

type ApisixRouteSpec struct {
	Rules []Rule `json:"rules,omitempty"`
}

type Rule struct {
	Host string `json:"host,omitempty"`
	Http Http   `json:"http,omitempty"`
}

type Http struct {
	Paths []Path `json:"paths,omitempty"`
}

type Path struct {
	Path    string  `json:"path,omitempty"`
	Backend Backend `json:"backend,omitempty"`
}

type Backend struct {
	ServiceName string `json:"serviceName,omitempty"`
	ServicePort int64  `json:"servicePort,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ApisixRoute `json:"items,omitempty"`
}
