package server

import "github.com/kongweiguo/spire-tenant/pkg/common/telemetry"

// Call Counters (timing and success metrics)
// Allows adding labels in-code

// StartRotateServerSVIDCall return metric for
// Server's SVID Rotation.
func StartRotateServerSVIDCall(m telemetry.Metrics) *telemetry.CallCounter {
	return telemetry.StartCall(m, telemetry.SVID, telemetry.Rotate)
}

// End Call Counters
