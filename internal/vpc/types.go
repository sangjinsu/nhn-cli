package vpc

type VPC struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CIDRv4         string `json:"cidrv4"`
	TenantID       string `json:"tenant_id"`
	State          string `json:"state"`
	CreateTime     string `json:"create_time"`
	Shared         bool   `json:"shared"`
	RouterExternal bool   `json:"router:external"`
}

type VPCListResponse struct {
	VPCs []VPC `json:"vpcs"`
}

type VPCResponse struct {
	VPC VPC `json:"vpc"`
}

type VPCCreateRequest struct {
	VPC VPCCreateBody `json:"vpc"`
}

type VPCCreateBody struct {
	Name   string `json:"name"`
	CIDRv4 string `json:"cidrv4"`
}

type VPCUpdateRequest struct {
	VPC VPCUpdateBody `json:"vpc"`
}

type VPCUpdateBody struct {
	Name   string `json:"name,omitempty"`
	CIDRv4 string `json:"cidrv4,omitempty"`
}

type Subnet struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	VPCID            string `json:"vpc_id"`
	CIDR             string `json:"cidr"`
	TenantID         string `json:"tenant_id"`
	State            string `json:"state"`
	CreateTime       string `json:"create_time"`
	AvailableIPCount int    `json:"available_ip_count"`
	VPC              *VPC   `json:"vpc,omitempty"`
	Shared           bool   `json:"shared"`
	Gateway          string `json:"gateway"`
	RoutingTableID   string `json:"routingtable_id"`
}

type SubnetListResponse struct {
	VPCSubnets []Subnet `json:"vpcsubnets"`
}

type SubnetResponse struct {
	VPCSubnet Subnet `json:"vpcsubnet"`
}

type SubnetCreateRequest struct {
	VPCSubnet SubnetCreateBody `json:"vpcsubnet"`
}

type SubnetCreateBody struct {
	VPCID string `json:"vpc_id"`
	Name  string `json:"name"`
	CIDR  string `json:"cidr"`
}

type SecurityGroup struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	TenantID    string              `json:"tenant_id"`
	Rules       []SecurityGroupRule `json:"security_group_rules"`
}

type SecurityGroupRule struct {
	ID              string `json:"id"`
	SecurityGroupID string `json:"security_group_id"`
	Direction       string `json:"direction"`
	EtherType       string `json:"ethertype"`
	Protocol        string `json:"protocol"`
	PortRangeMin    *int   `json:"port_range_min"`
	PortRangeMax    *int   `json:"port_range_max"`
	RemoteIPPrefix  string `json:"remote_ip_prefix"`
	RemoteGroupID   string `json:"remote_group_id"`
	Description     string `json:"description"`
}

type SecurityGroupListResponse struct {
	SecurityGroups []SecurityGroup `json:"security_groups"`
}

type SecurityGroupResponse struct {
	SecurityGroup SecurityGroup `json:"security_group"`
}

type SecurityGroupCreateRequest struct {
	SecurityGroup SecurityGroupCreateBody `json:"security_group"`
}

type SecurityGroupCreateBody struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type SecurityGroupRuleCreateRequest struct {
	SecurityGroupRule SecurityGroupRuleCreateBody `json:"security_group_rule"`
}

type SecurityGroupRuleCreateBody struct {
	SecurityGroupID string `json:"security_group_id"`
	Direction       string `json:"direction"`
	EtherType       string `json:"ethertype"`
	Protocol        string `json:"protocol,omitempty"`
	PortRangeMin    *int   `json:"port_range_min,omitempty"`
	PortRangeMax    *int   `json:"port_range_max,omitempty"`
	RemoteIPPrefix  string `json:"remote_ip_prefix,omitempty"`
	RemoteGroupID   string `json:"remote_group_id,omitempty"`
	Description     string `json:"description,omitempty"`
}

type FloatingIP struct {
	ID                string `json:"id"`
	FloatingIPAddress string `json:"floating_ip_address"`
	FloatingNetworkID string `json:"floating_network_id"`
	FixedIPAddress    string `json:"fixed_ip_address"`
	PortID            string `json:"port_id"`
	TenantID          string `json:"tenant_id"`
	Status            string `json:"status"`
}

type FloatingIPListResponse struct {
	FloatingIPs []FloatingIP `json:"floatingips"`
}

type FloatingIPResponse struct {
	FloatingIP FloatingIP `json:"floatingip"`
}

type FloatingIPCreateRequest struct {
	FloatingIP FloatingIPCreateBody `json:"floatingip"`
}

type FloatingIPCreateBody struct {
	FloatingNetworkID string `json:"floating_network_id"`
	PortID            string `json:"port_id,omitempty"`
}

type FloatingIPUpdateRequest struct {
	FloatingIP FloatingIPUpdateBody `json:"floatingip"`
}

type FloatingIPUpdateBody struct {
	PortID *string `json:"port_id"`
}

type RoutingTable struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	TenantID     string  `json:"tenant_id"`
	DefaultTable bool    `json:"default_table"`
	Distributed  bool    `json:"distributed"`
	State        string  `json:"state"`
	VPCID        string  `json:"vpc_id"`
	SubnetCount  int     `json:"subnet_count"`
	Routes       []Route `json:"routes"`
}

type Route struct {
	ID              string `json:"id"`
	RoutingTableID  string `json:"routingtable_id"`
	DestinationCIDR string `json:"cidr"`
	GatewayID       string `json:"gateway_id"`
	TenantID        string `json:"tenant_id"`
}

type RoutingTableListResponse struct {
	RoutingTables []RoutingTable `json:"routingtables"`
}

type RoutingTableResponse struct {
	RoutingTable RoutingTable `json:"routingtable"`
}

type Port struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	NetworkID    string    `json:"network_id"`
	TenantID     string    `json:"tenant_id"`
	MACAddress   string    `json:"mac_address"`
	AdminStateUp bool      `json:"admin_state_up"`
	Status       string    `json:"status"`
	DeviceID     string    `json:"device_id"`
	DeviceOwner  string    `json:"device_owner"`
	FixedIPs     []FixedIP `json:"fixed_ips"`
}

type FixedIP struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
}

type PortListResponse struct {
	Ports []Port `json:"ports"`
}

type PortResponse struct {
	Port Port `json:"port"`
}

type PortCreateRequest struct {
	Port PortCreateBody `json:"port"`
}

type PortCreateBody struct {
	NetworkID    string `json:"network_id"`
	Name         string `json:"name,omitempty"`
	AdminStateUp bool   `json:"admin_state_up"`
}
