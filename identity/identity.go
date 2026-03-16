// Package identity defines canonical platform identity constants.
// All services that make or receive inter-service HTTP calls must
// import these constants — never hardcode header strings locally.
//
// ADR-016: types and constants only. No logic, no HTTP clients.
package identity

// HTTP header constants for inter-service communication.
// Import these — never write "X-Service-Token" or "X-Trace-ID" as literals.
const (
	// TraceIDHeader is the HTTP header for distributed trace propagation.
	// Set by Nexus, Atlas, and Forge TraceID middleware (ADR-015 / Phase 15-16).
	TraceIDHeader = "X-Trace-ID"

	// ServiceTokenHeader is the HTTP header for inter-service authentication.
	// Required on all non-health inter-service calls (ADR-008).
	ServiceTokenHeader = "X-Service-Token"
)

// Platform service name constants.
// Used as component identifiers in events, logs, and trace entries.
const (
	ServiceNexus     = "nexus"
	ServiceAtlas     = "atlas"
	ServiceForge     = "forge"
	ServiceMetrics   = "metrics"
	ServiceNavigator = "navigator"
	ServiceGuardian  = "guardian"
	ServiceObserver  = "observer"
	ServiceSentinel  = "sentinel"
)

// Platform service default addresses.
// Override with environment variables in production.
const (
	DefaultNexusAddr     = "http://127.0.0.1:8080"
	DefaultAtlasAddr     = "http://127.0.0.1:8081"
	DefaultForgeAddr     = "http://127.0.0.1:8082"
	DefaultMetricsAddr   = "http://127.0.0.1:8083"
	DefaultNavigatorAddr = "http://127.0.0.1:8084"
	DefaultGuardianAddr  = "http://127.0.0.1:8085"
	DefaultObserverAddr  = "http://127.0.0.1:8086"
	DefaultSentinelAddr  = "http://127.0.0.1:8087"
)
