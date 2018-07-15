// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Resources A container's resources (cgroups config, ulimits, etc)
// swagger:model Resources
type Resources struct {

	// Limit read rate (bytes per second) from a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	//
	BlkioDeviceReadBps []*ThrottleDevice `json:"BlkioDeviceReadBps"`

	// Limit read rate (IO per second) from a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	//
	BlkioDeviceReadIOps []*ThrottleDevice `json:"BlkioDeviceReadIOps"`

	// Limit write rate (bytes per second) to a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	//
	BlkioDeviceWriteBps []*ThrottleDevice `json:"BlkioDeviceWriteBps"`

	// Limit write rate (IO per second) to a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	//
	BlkioDeviceWriteIOps []*ThrottleDevice `json:"BlkioDeviceWriteIOps"`

	// Block IO weight (relative weight), need CFQ IO Scheduler enable.
	// Maximum: 1000
	// Minimum: 0
	BlkioWeight uint16 `json:"BlkioWeight,omitempty"`

	// Block IO weight (relative device weight) in the form `[{"Path": "device_path", "Weight": weight}]`.
	//
	BlkioWeightDevice []*WeightDevice `json:"BlkioWeightDevice"`

	// CPU CFS (Completely Fair Scheduler) period.
	// The length of a CPU period in microseconds.
	//
	// Maximum: 1e+06
	// Minimum: 1000
	CPUPeriod int64 `json:"CPUPeriod,omitempty"`

	// CPU CFS (Completely Fair Scheduler) quota.
	// Microseconds of CPU time that the container can get in a CPU period."
	//
	// Minimum: 1000
	CPUQuota int64 `json:"CPUQuota,omitempty"`

	// An integer value representing this container's relative CPU weight versus other containers.
	CPUShares int64 `json:"CPUShares,omitempty"`

	// Path to `cgroups` under which the container's `cgroup` is created. If the path is not absolute, the path is considered to be relative to the `cgroups` path of the init process. Cgroups are created if they do not already exist.
	CgroupParent string `json:"CgroupParent,omitempty"`

	// The number of usable CPUs (Windows only).
	// On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	//
	CPUCount int64 `json:"CpuCount,omitempty"`

	// The usable percentage of the available CPUs (Windows only).
	// On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	//
	CPUPercent int64 `json:"CpuPercent,omitempty"`

	// The length of a CPU real-time period in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	CPURealtimePeriod int64 `json:"CpuRealtimePeriod,omitempty"`

	// The length of a CPU real-time runtime in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	CPURealtimeRuntime int64 `json:"CpuRealtimeRuntime,omitempty"`

	// CPUs in which to allow execution (e.g., `0-3`, `0,1`)
	CpusetCpus string `json:"CpusetCpus,omitempty"`

	// Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.
	CpusetMems string `json:"CpusetMems,omitempty"`

	// a list of cgroup rules to apply to the container
	DeviceCgroupRules []string `json:"DeviceCgroupRules"`

	// A list of devices to add to the container.
	Devices []*DeviceMapping `json:"Devices"`

	// Maximum IO in bytes per second for the container system drive (Windows only)
	IOMaximumBandwidth uint64 `json:"IOMaximumBandwidth,omitempty"`

	// Maximum IOps for the container system drive (Windows only)
	IOMaximumIOps uint64 `json:"IOMaximumIOps,omitempty"`

	// IntelRdtL3Cbm specifies settings for Intel RDT/CAT group that the container is placed into to limit the resources (e.g., L3 cache) the container has available.
	IntelRdtL3Cbm string `json:"IntelRdtL3Cbm,omitempty"`

	// Kernel memory limit in bytes.
	KernelMemory int64 `json:"KernelMemory,omitempty"`

	// Memory limit in bytes.
	Memory int64 `json:"Memory,omitempty"`

	// MemoryExtra is an integer value representing this container's memory high water mark percentage.
	// The range is in [0, 100].
	//
	// Maximum: 100
	// Minimum: 0
	MemoryExtra *int64 `json:"MemoryExtra,omitempty"`

	// MemoryForceEmptyCtl represents whether to reclaim the page cache when deleting cgroup.
	// Maximum: 1
	// Minimum: 0
	MemoryForceEmptyCtl int64 `json:"MemoryForceEmptyCtl,omitempty"`

	// Memory soft limit in bytes.
	MemoryReservation int64 `json:"MemoryReservation,omitempty"`

	// Total memory limit (memory + swap). Set as `-1` to enable unlimited swap.
	MemorySwap int64 `json:"MemorySwap,omitempty"`

	// Tune a container's memory swappiness behavior. Accepts an integer between 0 and 100.
	// Maximum: 100
	// Minimum: 0
	MemorySwappiness *int64 `json:"MemorySwappiness,omitempty"`

	// MemoryWmarkRatio is an integer value representing this container's memory low water mark percentage.
	// The value of memory low water mark is memory.limit_in_bytes * MemoryWmarkRatio. The range is in [0, 100].
	//
	// Maximum: 100
	// Minimum: 0
	MemoryWmarkRatio *int64 `json:"MemoryWmarkRatio,omitempty"`

	// CPU quota in units of 10<sup>-9</sup> CPUs.
	NanoCpus int64 `json:"NanoCPUs,omitempty"`

	// Disable OOM Killer for the container.
	OomKillDisable *bool `json:"OomKillDisable,omitempty"`

	// Tune a container's pids limit. Set -1 for unlimited. Only on Linux 4.4 does this parameter support.
	//
	PidsLimit int64 `json:"PidsLimit,omitempty"`

	// ScheLatSwitch enables scheduler latency count in cpuacct
	// Maximum: 1
	// Minimum: 0
	ScheLatSwitch int64 `json:"ScheLatSwitch,omitempty"`

	// A list of resource limits to set in the container. For example: `{"Name": "nofile", "Soft": 1024, "Hard": 2048}`"
	//
	Ulimits []*Ulimit `json:"Ulimits"`
}

// Validate validates this resources
func (m *Resources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlkioDeviceReadBps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioDeviceReadIOps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioDeviceWriteBps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioDeviceWriteIOps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioWeightDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryExtra(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryForceEmptyCtl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemorySwappiness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryWmarkRatio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheLatSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUlimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resources) validateBlkioDeviceReadBps(formats strfmt.Registry) error {

	if swag.IsZero(m.BlkioDeviceReadBps) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioDeviceReadBps); i++ {
		if swag.IsZero(m.BlkioDeviceReadBps[i]) { // not required
			continue
		}

		if m.BlkioDeviceReadBps[i] != nil {
			if err := m.BlkioDeviceReadBps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceReadBps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateBlkioDeviceReadIOps(formats strfmt.Registry) error {

	if swag.IsZero(m.BlkioDeviceReadIOps) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioDeviceReadIOps); i++ {
		if swag.IsZero(m.BlkioDeviceReadIOps[i]) { // not required
			continue
		}

		if m.BlkioDeviceReadIOps[i] != nil {
			if err := m.BlkioDeviceReadIOps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceReadIOps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateBlkioDeviceWriteBps(formats strfmt.Registry) error {

	if swag.IsZero(m.BlkioDeviceWriteBps) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioDeviceWriteBps); i++ {
		if swag.IsZero(m.BlkioDeviceWriteBps[i]) { // not required
			continue
		}

		if m.BlkioDeviceWriteBps[i] != nil {
			if err := m.BlkioDeviceWriteBps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceWriteBps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateBlkioDeviceWriteIOps(formats strfmt.Registry) error {

	if swag.IsZero(m.BlkioDeviceWriteIOps) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioDeviceWriteIOps); i++ {
		if swag.IsZero(m.BlkioDeviceWriteIOps[i]) { // not required
			continue
		}

		if m.BlkioDeviceWriteIOps[i] != nil {
			if err := m.BlkioDeviceWriteIOps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceWriteIOps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateBlkioWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.BlkioWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("BlkioWeight", "body", int64(m.BlkioWeight), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("BlkioWeight", "body", int64(m.BlkioWeight), 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateBlkioWeightDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.BlkioWeightDevice) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioWeightDevice); i++ {
		if swag.IsZero(m.BlkioWeightDevice[i]) { // not required
			continue
		}

		if m.BlkioWeightDevice[i] != nil {
			if err := m.BlkioWeightDevice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioWeightDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateCPUPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUPeriod) { // not required
		return nil
	}

	if err := validate.MinimumInt("CPUPeriod", "body", int64(m.CPUPeriod), 1000, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("CPUPeriod", "body", int64(m.CPUPeriod), 1e+06, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateCPUQuota(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUQuota) { // not required
		return nil
	}

	if err := validate.MinimumInt("CPUQuota", "body", int64(m.CPUQuota), 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateMemoryExtra(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryExtra) { // not required
		return nil
	}

	if err := validate.MinimumInt("MemoryExtra", "body", int64(*m.MemoryExtra), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("MemoryExtra", "body", int64(*m.MemoryExtra), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateMemoryForceEmptyCtl(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryForceEmptyCtl) { // not required
		return nil
	}

	if err := validate.MinimumInt("MemoryForceEmptyCtl", "body", int64(m.MemoryForceEmptyCtl), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("MemoryForceEmptyCtl", "body", int64(m.MemoryForceEmptyCtl), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateMemorySwappiness(formats strfmt.Registry) error {

	if swag.IsZero(m.MemorySwappiness) { // not required
		return nil
	}

	if err := validate.MinimumInt("MemorySwappiness", "body", int64(*m.MemorySwappiness), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("MemorySwappiness", "body", int64(*m.MemorySwappiness), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateMemoryWmarkRatio(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryWmarkRatio) { // not required
		return nil
	}

	if err := validate.MinimumInt("MemoryWmarkRatio", "body", int64(*m.MemoryWmarkRatio), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("MemoryWmarkRatio", "body", int64(*m.MemoryWmarkRatio), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateScheLatSwitch(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheLatSwitch) { // not required
		return nil
	}

	if err := validate.MinimumInt("ScheLatSwitch", "body", int64(m.ScheLatSwitch), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("ScheLatSwitch", "body", int64(m.ScheLatSwitch), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateUlimits(formats strfmt.Registry) error {

	if swag.IsZero(m.Ulimits) { // not required
		return nil
	}

	for i := 0; i < len(m.Ulimits); i++ {
		if swag.IsZero(m.Ulimits[i]) { // not required
			continue
		}

		if m.Ulimits[i] != nil {
			if err := m.Ulimits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Ulimits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resources) UnmarshalBinary(b []byte) error {
	var res Resources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
