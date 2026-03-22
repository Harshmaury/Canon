// Package identity defines canonical platform identity constants.
// All services that make or receive inter-service HTTP calls must
// import these constants — never hardcode header strings locally.
//
// ADR-016: types and constants only. No logic, no HTTP clients.
// ADR-042: IdentityTokenHeader, GateService, DefaultGateAddr added.
// ADR-045: Relay constants added. Canon graduates to v1.0.0.
package identity

// HTTP header constants for inter-service communication.
// Import these — never write header strings as literals.
const (
	// TraceIDHeader is the HTTP header for distributed trace propagation.
	// Set by Nexus, Atlas, and Forge TraceID middleware (ADR-015 / Phase 15-16).
	TraceIDHeader = "X-Trace-ID"

	// ServiceTokenHeader is the HTTP header for inter-service authentication.
	// Required on all non-health inter-service calls (ADR-008).
	// Proves the caller is a known platform service — not an actor identity.
	ServiceTokenHeader = "X-Service-Token"

	// IdentityTokenHeader carries the Gate-issued actor identity token (ADR-042).
	// Present on requests from human actors, agents, and CI pipelines.
	// Independent of ServiceTokenHeader — both may be present simultaneously.
	// Never required on GET /health.
	IdentityTokenHeader = "X-Identity-Token"

	// RelayTokenHeader authenticates an engxa tunnel connection to Relay (ADR-041).
	// Distinct from ServiceTokenHeader — proves tunnel ownership, not service identity.
	RelayTokenHeader = "X-Relay-Token"

	// SubdomainHeader carries the subdomain assigned to an active tunnel (ADR-041).
	// Set by Relay on every forwarded inbound request.
	SubdomainHeader = "X-Engx-Subdomain"

	// OwnerHeader carries the tunnel owner identifier on forwarded requests (ADR-041).
	// Value is the owner string assigned at tunnel registration (e.g. "harsh").
	OwnerHeader = "X-Engx-Owner"
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
	ServiceGate      = "gate"
	ServiceRelay     = "relay" // ADR-041: Relay tunnel service
)

// Platform service default addresses.
// Override with environment variables in production.
const (
	DefaultNexusAddr      = "http://127.0.0.1:8080"
	DefaultAtlasAddr      = "http://127.0.0.1:8081"
	DefaultForgeAddr      = "http://127.0.0.1:8082"
	DefaultMetricsAddr    = "http://127.0.0.1:8083"
	DefaultNavigatorAddr  = "http://127.0.0.1:8084"
	DefaultGuardianAddr   = "http://127.0.0.1:8085"
	DefaultObserverAddr   = "http://127.0.0.1:8086"
	DefaultSentinelAddr   = "http://127.0.0.1:8087"
	DefaultGateAddr       = "http://127.0.0.1:8088"

	// DefaultRelayTunnelAddr is the default Relay tunnel listener address (ADR-041).
	// engxa connects here to open a tunnel. Override with RELAY_TUNNEL_ADDR.
	DefaultRelayTunnelAddr = "relay.engx.dev:9090"
)
