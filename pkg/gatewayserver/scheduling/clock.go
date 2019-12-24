// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scheduling

import (
	"math"
	"time"
)

// Clock represents an absolute time source.
type Clock interface {
	// IsSynced returns whether the clock is synchronized.
	IsSynced() bool
	// FromServerTime returns an indication of the concentrator time at the given server time.
	FromServerTime(time.Time) (ConcentratorTime, bool)
	// ToServerTime returns an indication of the server time at the given concentrator time.
	ToServerTime(ConcentratorTime) time.Time
	// FromGatewayTime returns an indication of the concentrator time at the given gateway time if available.
	FromGatewayTime(time.Time) (ConcentratorTime, bool)
	// FromTimestampTime returns the concentrator time for the given timestamp.
	FromTimestampTime(timestamp uint32) ConcentratorTime
}

// RolloverClock is a Clock that takes roll-over of a uint32 microsecond concentrator time into account.
type RolloverClock struct {
	synced   bool
	relative uint32
	absolute ConcentratorTime
	server   *time.Time
	gateway  *time.Time
}

// IsSynced implements Clock.
func (c *RolloverClock) IsSynced() bool { return c.synced }

// Sync synchronizes the clock with the given concentrator timestamp and the server time.
func (c *RolloverClock) Sync(timestamp uint32, server time.Time) {
	rollovers := int(c.absolute/ConcentratorTime(time.Microsecond)) / math.MaxUint32
	if passed := int64(timestamp) - int64(c.relative); passed < 0 {
		rollovers++
	}
	if c.server != nil {
		rollovers += int(server.Sub(*c.server)/time.Microsecond) / math.MaxUint32
	}
	c.absolute = (ConcentratorTime(rollovers*math.MaxUint32) + ConcentratorTime(timestamp)) * ConcentratorTime(time.Microsecond)
	c.relative = timestamp
	c.server = &server
	c.gateway = nil
	c.synced = true
}

// SyncWithGatewayAbsolute synchronizes the clock with the given concentrator timestamp, the server time and the
// absolute gateway time that corresponds to the given timestamp.
func (c *RolloverClock) SyncWithGatewayAbsolute(timestamp uint32, server, gateway time.Time) {
	c.Sync(timestamp, server)
	c.gateway = &gateway
}

// SyncWithGatewayConcentrator synchronizes the clock with the given concentrator timestamp, the server time and the
// relative gateway time that corresponds to the given timestamp.
func (c *RolloverClock) SyncWithGatewayConcentrator(timestamp uint32, server time.Time, concentrator ConcentratorTime) {
	c.absolute = concentrator
	c.relative = timestamp
	c.server = &server
	c.gateway = nil
	c.synced = true
}

// FromServerTime implements Clock.
func (c *RolloverClock) FromServerTime(server time.Time) (ConcentratorTime, bool) {
	if c.server == nil {
		return 0, false
	}
	return c.absolute + ConcentratorTime(server.Sub(*c.server)), true
}

// ToServerTime implements Clock.
func (c *RolloverClock) ToServerTime(t ConcentratorTime) time.Time {
	return c.server.Add(time.Duration(t - c.absolute))
}

// FromGatewayTime implements Clock.
func (c *RolloverClock) FromGatewayTime(gateway time.Time) (ConcentratorTime, bool) {
	if c.gateway == nil {
		return 0, false
	}
	return c.absolute + ConcentratorTime(gateway.Sub(*c.gateway)), true
}

// FromTimestampTime implements Clock.
func (c *RolloverClock) FromTimestampTime(timestamp uint32) ConcentratorTime {
	passed := int64(timestamp) - int64(c.relative)
	if passed < -math.MaxUint32/2 {
		passed += 1 << 32
	} else if passed > math.MaxUint32/2 {
		passed -= 1 << 32
	}
	return c.absolute + ConcentratorTime(time.Duration(passed)*time.Microsecond)
}
