package service

import (
	"reflect"
	"traveldeck/domain"
)

type NullCapability struct{ Enabled bool }

func (c *NullCapability) Ready() bool {
	if c == nil {
		return false
	}
	return c.Enabled
}

// capabilityReady reports whether the dependency can actually serve requests.
// A nil interface, or a non-nil interface wrapping a nil pointer (an "empty
// implementation" produced, for example, by a typed-nil *NullCapability),
// must be treated as unavailable instead of dereferenced: calling Ready on
// such a value would panic. The reflect check closes the typed-nil hole that a
// plain `cap == nil` comparison misses.
func capabilityReady(cap domain.Capability) bool {
	if cap == nil {
		return false
	}
	v := reflect.ValueOf(cap)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return false
	}
	return cap.Ready()
}

// CapabilityAvailable reports whether cap is a usable dependency. It never
// panics, even for empty/typed-nil implementations.
func CapabilityAvailable(cap domain.Capability) bool {
	return capabilityReady(cap)
}
