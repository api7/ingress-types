package v1

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixRoute struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec              *ApisixRouteSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
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
	Path    string   `json:"path,omitempty"`
	Backend Backend  `json:"backend,omitempty"`
	Plugins []Plugin `json:"plugins,omitempty"`
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

// +genclient
// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixUpstream struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec              *ApisixUpstreamSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type ApisixUpstreamSpec struct {
	Ports []Port `json:"ports,omitempty"`
}

type Port struct {
	Port         int64        `json:"port,omitempty"`
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

// +genclient
// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixService struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec              *ApisixServiceSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ApisixService `json:"items,omitempty"`
}

type ApisixServiceSpec struct {
	Upstream string   `json:"upstream,omitempty"`
	Port     int64    `json:"port,omitempty"`
	Plugins  []Plugin `json:"plugins,omitempty"`
}

type Plugin struct {
	Name      string    `json:"name,omitempty"`
	Enable    bool      `json:"enable,omitempty"`
	Config    Config    `json:"config,omitempty"`
	ConfigSet ConfigSet `json:"config_set,omitempty"`
}

type ConfigSet []interface{}

func (p ConfigSet) DeepCopyInto(out *ConfigSet) {
	b, _ := json.Marshal(&p)
	_ = json.Unmarshal(b, out)
}

func (p *ConfigSet) DeepCopy() *ConfigSet {
	if p == nil {
		return nil
	}
	out := new(ConfigSet)
	p.DeepCopyInto(out)
	return out
}

type Config map[string]interface{}

func (p Config) DeepCopyInto(out *Config) {
	b, _ := json.Marshal(&p)
	_ = json.Unmarshal(b, out)
}

func (p *Config) DeepCopy() *Config {
	if p == nil {
		return nil
	}
	out := new(Config)
	p.DeepCopyInto(out)
	return out
}

// +genclient
// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixTls struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec              *ApisixTlsSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApisixTlsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ApisixTls `json:"items,omitempty"`
}

type ApisixTlsSpec struct {
	Hosts  []string     `json:"hosts,omitempty"`
	Secret ApisixSecret `json:"secret,omitempty"`
}

type ApisixSecret struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}
