package compute

import "time"

type Instance struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Status      string            `json:"status"`
	TenantID    string            `json:"tenant_id"`
	UserID      string            `json:"user_id"`
	Created     time.Time         `json:"created"`
	Updated     time.Time         `json:"updated"`
	Flavor      FlavorRef         `json:"flavor"`
	Image       ImageRef          `json:"image"`
	KeyName     string            `json:"key_name"`
	Addresses   map[string][]Address `json:"addresses"`
	Metadata    map[string]string `json:"metadata"`
	SecurityGroups []SecurityGroupRef `json:"security_groups"`
	AvailabilityZone string `json:"OS-EXT-AZ:availability_zone"`
	PowerState  int               `json:"OS-EXT-STS:power_state"`
	VMState     string            `json:"OS-EXT-STS:vm_state"`
	TaskState   string            `json:"OS-EXT-STS:task_state"`
}

type FlavorRef struct {
	ID    string `json:"id"`
	Links []Link `json:"links"`
}

type ImageRef struct {
	ID    string `json:"id"`
	Links []Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type Address struct {
	Addr    string `json:"addr"`
	Version int    `json:"version"`
	Type    string `json:"OS-EXT-IPS:type"`
	MACAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
}

type SecurityGroupRef struct {
	Name string `json:"name"`
}

type InstanceListResponse struct {
	Servers []Instance `json:"servers"`
}

type InstanceResponse struct {
	Server Instance `json:"server"`
}

type InstanceCreateRequest struct {
	Server InstanceCreateBody `json:"server"`
}

type InstanceCreateBody struct {
	Name             string            `json:"name"`
	ImageRef         string            `json:"imageRef"`
	FlavorRef        string            `json:"flavorRef"`
	KeyName          string            `json:"key_name,omitempty"`
	SecurityGroups   []SecurityGroupRef `json:"security_groups,omitempty"`
	Networks         []NetworkRef      `json:"networks,omitempty"`
	AvailabilityZone string            `json:"availability_zone,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
	UserData         string            `json:"user_data,omitempty"`
	BlockDeviceMapping []BlockDeviceMapping `json:"block_device_mapping_v2,omitempty"`
}

type NetworkRef struct {
	UUID    string `json:"uuid,omitempty"`
	Port    string `json:"port,omitempty"`
	FixedIP string `json:"fixed_ip,omitempty"`
}

type BlockDeviceMapping struct {
	BootIndex           int    `json:"boot_index"`
	UUID                string `json:"uuid,omitempty"`
	SourceType          string `json:"source_type"`
	DestinationType     string `json:"destination_type"`
	VolumeSize          int    `json:"volume_size,omitempty"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type InstanceActionRequest struct {
	Start   *struct{}      `json:"os-start,omitempty"`
	Stop    *struct{}      `json:"os-stop,omitempty"`
	Reboot  *RebootAction  `json:"reboot,omitempty"`
}

type RebootAction struct {
	Type string `json:"type"`
}

type Flavor struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	RAM        int     `json:"ram"`
	VCPUs      int     `json:"vcpus"`
	Disk       int     `json:"disk"`
	Ephemeral  int     `json:"OS-FLV-EXT-DATA:ephemeral"`
	RxTxFactor float64 `json:"rxtx_factor"`
	IsPublic   bool    `json:"os-flavor-access:is_public"`
}

type FlavorListResponse struct {
	Flavors []Flavor `json:"flavors"`
}

type FlavorResponse struct {
	Flavor Flavor `json:"flavor"`
}

type Image struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Status    string            `json:"status"`
	Created   time.Time         `json:"created"`
	Updated   time.Time         `json:"updated"`
	MinDisk   int               `json:"minDisk"`
	MinRAM    int               `json:"minRam"`
	Progress  int               `json:"progress"`
	Metadata  map[string]string `json:"metadata"`
}

type ImageListResponse struct {
	Images []Image `json:"images"`
}

type ImageResponse struct {
	Image Image `json:"image"`
}

type Keypair struct {
	Name        string `json:"name"`
	PublicKey   string `json:"public_key"`
	Fingerprint string `json:"fingerprint"`
	UserID      string `json:"user_id"`
}

type KeypairWrapper struct {
	Keypair Keypair `json:"keypair"`
}

type KeypairListResponse struct {
	Keypairs []KeypairWrapper `json:"keypairs"`
}

type KeypairResponse struct {
	Keypair KeypairCreated `json:"keypair"`
}

type KeypairCreated struct {
	Name        string `json:"name"`
	PublicKey   string `json:"public_key"`
	PrivateKey  string `json:"private_key"`
	Fingerprint string `json:"fingerprint"`
	UserID      string `json:"user_id"`
}

type KeypairCreateRequest struct {
	Keypair KeypairCreateBody `json:"keypair"`
}

type KeypairCreateBody struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key,omitempty"`
}

type AvailabilityZone struct {
	ZoneName  string `json:"zoneName"`
	ZoneState struct {
		Available bool `json:"available"`
	} `json:"zoneState"`
}

type AvailabilityZoneListResponse struct {
	AvailabilityZoneInfo []AvailabilityZone `json:"availabilityZoneInfo"`
}
