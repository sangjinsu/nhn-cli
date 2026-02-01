package blockstorage

import (
	"fmt"
	"strings"
	"time"
)

// FlexTime handles multiple time formats from NHN Cloud API
type FlexTime struct {
	time.Time
}

func (ft *FlexTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		return nil
	}

	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05.000000",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
	}

	for _, f := range formats {
		t, err := time.Parse(f, s)
		if err == nil {
			ft.Time = t
			return nil
		}
	}
	return fmt.Errorf("unable to parse time: %s", s)
}

type Volume struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	Status           string            `json:"status"`
	Size             int               `json:"size"`
	VolumeType       string            `json:"volume_type"`
	AvailabilityZone string            `json:"availability_zone"`
	CreatedAt        FlexTime          `json:"created_at"`
	UpdatedAt        FlexTime          `json:"updated_at"`
	Attachments      []Attachment      `json:"attachments"`
	Metadata         map[string]string `json:"metadata"`
	SnapshotID       string            `json:"snapshot_id"`
	Bootable         string            `json:"bootable"`
	Encrypted        bool              `json:"encrypted"`
	TenantID         string            `json:"os-vol-tenant-attr:tenant_id"`
}

type Attachment struct {
	ID         string `json:"id"`
	VolumeID   string `json:"volume_id"`
	ServerID   string `json:"server_id"`
	Device     string `json:"device"`
	AttachedAt string `json:"attached_at"`
}

type VolumeListResponse struct {
	Volumes []Volume `json:"volumes"`
}

type VolumeResponse struct {
	Volume Volume `json:"volume"`
}

type VolumeCreateRequest struct {
	Volume VolumeCreateBody `json:"volume"`
}

type VolumeCreateBody struct {
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	Size             int    `json:"size"`
	VolumeType       string `json:"volume_type,omitempty"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
	SnapshotID       string `json:"snapshot_id,omitempty"`
}

type Snapshot struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Status      string            `json:"status"`
	Size        int               `json:"size"`
	VolumeID    string            `json:"volume_id"`
	CreatedAt   FlexTime          `json:"created_at"`
	UpdatedAt   FlexTime          `json:"updated_at"`
	Metadata    map[string]string `json:"metadata"`
}

type SnapshotListResponse struct {
	Snapshots []Snapshot `json:"snapshots"`
}

type SnapshotResponse struct {
	Snapshot Snapshot `json:"snapshot"`
}

type SnapshotCreateRequest struct {
	Snapshot SnapshotCreateBody `json:"snapshot"`
}

type SnapshotCreateBody struct {
	VolumeID    string `json:"volume_id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Force       bool   `json:"force,omitempty"`
}

type VolumeType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type VolumeTypeListResponse struct {
	VolumeTypes []VolumeType `json:"volume_types"`
}
