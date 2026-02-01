package loadbalancer

type LoadBalancer struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	ProvisioningStatus string  `json:"provisioning_status"`
	OperatingStatus    string  `json:"operating_status"`
	AdminStateUp       bool    `json:"admin_state_up"`
	VipAddress         string  `json:"vip_address"`
	VipSubnetID        string  `json:"vip_subnet_id"`
	VipPortID          string  `json:"vip_port_id"`
	TenantID           string  `json:"tenant_id"`
	Listeners          []IDRef `json:"listeners"`
	Pools              []IDRef `json:"pools"`
	LoadbalancerType   string  `json:"loadbalancer_type"`
}

type IDRef struct {
	ID string `json:"id"`
}

type LoadBalancerListResponse struct {
	Loadbalancers []LoadBalancer `json:"loadbalancers"`
}

type LoadBalancerResponse struct {
	Loadbalancer LoadBalancer `json:"loadbalancer"`
}

type LoadBalancerCreateRequest struct {
	Loadbalancer LoadBalancerCreateBody `json:"loadbalancer"`
}

type LoadBalancerCreateBody struct {
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	VipSubnetID      string `json:"vip_subnet_id"`
	VipAddress       string `json:"vip_address,omitempty"`
	AdminStateUp     *bool  `json:"admin_state_up,omitempty"`
	LoadbalancerType string `json:"loadbalancer_type,omitempty"`
}

type LoadBalancerUpdateRequest struct {
	Loadbalancer LoadBalancerUpdateBody `json:"loadbalancer"`
}

type LoadBalancerUpdateBody struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	AdminStateUp *bool  `json:"admin_state_up,omitempty"`
}

type Listener struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Protocol           string `json:"protocol"`
	ProtocolPort       int    `json:"protocol_port"`
	LoadbalancerID     string `json:"loadbalancer_id,omitempty"`
	DefaultPoolID      string `json:"default_pool_id"`
	AdminStateUp       bool   `json:"admin_state_up"`
	TenantID           string `json:"tenant_id"`
	ProvisioningStatus string `json:"provisioning_status"`
	OperatingStatus    string `json:"operating_status"`
}

type ListenerListResponse struct {
	Listeners []Listener `json:"listeners"`
}

type ListenerResponse struct {
	Listener Listener `json:"listener"`
}

type ListenerCreateRequest struct {
	Listener ListenerCreateBody `json:"listener"`
}

type ListenerCreateBody struct {
	LoadbalancerID string `json:"loadbalancer_id"`
	Protocol       string `json:"protocol"`
	ProtocolPort   int    `json:"protocol_port"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	DefaultPoolID  string `json:"default_pool_id,omitempty"`
	AdminStateUp   *bool  `json:"admin_state_up,omitempty"`
}
