package api

import (
	"context"

	"github.com/zcalusic/sysinfo"
)

func RetreiveHardwareStats(ctx context.Context) *HardwareStat {
	var si sysinfo.SysInfo

	si.GetSysInfo()

	h := &HardwareStat{
		Product: &HWProduct{
			Name:    si.Product.Name,
			Vendor:  si.Product.Vendor,
			Version: si.Product.Version,
			Serial:  si.Product.Serial,
		},
		Board: &HWBoard{
			Name:     si.Board.Name,
			Vendor:   si.Board.Vendor,
			Version:  si.Board.Version,
			Serial:   si.Board.Serial,
			AssetTag: si.Board.AssetTag,
		},
		Chassis: &HWChassis{
			Type:     uint64(si.Chassis.Type),
			Vendor:   si.Chassis.Vendor,
			Version:  si.Chassis.Version,
			Serial:   si.Chassis.Serial,
			AssetTag: si.Chassis.AssetTag,
		},
		Bios: &HWBIOS{
			Vendor:  si.BIOS.Vendor,
			Version: si.BIOS.Version,
			Date:    si.BIOS.Date,
		},
		Cpu: &HWCPU{
			Vendor:  si.CPU.Vendor,
			Model:   si.CPU.Model,
			Speed:   uint64(si.CPU.Speed),
			Cache:   uint64(si.CPU.Cache),
			Cpus:    uint64(si.CPU.Cpus),
			Cores:   uint64(si.CPU.Cores),
			Threads: uint64(si.CPU.Threads),
		},
		Memory: &HWMemory{
			Type:  si.Memory.Type,
			Speed: uint64(si.Memory.Speed),
			Size:  uint64(si.Memory.Size),
		},
		Storage: []*HWStorageDevice{},
		Network: []*HWNetwork{},
	}

	for _, s := range si.Storage {
		h.Storage = append(h.Storage, convertStorageDeviceToHWStorage(s))
	}

	for _, s := range si.Network {
		h.Network = append(h.Network, convertNetworkDeviceToHWNetwork(s))
	}

	return h
}

func convertStorageDeviceToHWStorage(s sysinfo.StorageDevice) *HWStorageDevice {
	return &HWStorageDevice{
		Name:   s.Name,
		Driver: s.Driver,
		Vendor: s.Vendor,
		Model:  s.Model,
		Serial: s.Serial,
		Size:   uint64(s.Size),
	}
}

func convertNetworkDeviceToHWNetwork(s sysinfo.NetworkDevice) *HWNetwork {
	return &HWNetwork{
		Name:       s.Name,
		Driver:     s.Driver,
		MacAddress: s.MACAddress,
		Port:       s.Port,
		Speed:      uint64(s.Speed),
	}
}
