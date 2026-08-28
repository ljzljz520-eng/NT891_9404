package service

import "traveldeck/domain"

type NullCapability struct{ Enabled bool }

func (c *NullCapability) Ready() bool { return c.Enabled }

func CapabilityAvailable(cap domain.Capability) bool {
	if cap == nil {
		return false
	}
	return cap.Ready()
}
