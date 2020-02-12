package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"encoding/json"
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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixUpstream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *ApisixUpstreamSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixUpstreamSpec struct {
	Ports []Port `json:"ports,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Port struct {
	Port int64 `json:"port,omitempty"`
	Loadbalancer Loadbalancer `json:"loadbalancer,omitempty"`
}

type Loadbalancer map[string]interface{}

func (p Loadbalancer) DeepCopyInto(out *Loadbalancer) {
	b, _ := json.Marshal(&p)
	_ = json.Unmarshal(b, out)
}

func (p *Loadbalancer) DeepCopy() *Loadbalancer {
	if p == nil {
		return nil
	}
	out := new(Loadbalancer)
	p.DeepCopyInto(out)
	return out
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixUpstreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ApisixUpstream `json:"items,omitempty"`
}
