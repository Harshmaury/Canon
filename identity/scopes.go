// Package identity defines canonical platform identity constants.
//
// ADR-042: scope constants for Gate-issued tokens.
// Import from Canon — never hardcode scope strings in any service.
// Scope values are permanent API surface. Never rename or remove.
package identity

// Scope constants for Gate-issued identity tokens.
// A token carries zero or more scopes. Services check for the required
// scope before permitting an operation. Missing scope = HTTP 403 ErrForbidden.
const (
	// ScopeExecute permits Forge command submission and project start/stop.
	ScopeExecute = "execute"

	// ScopeObserve permits read access to events, metrics, history, and topology.
	ScopeObserve = "observe"

	// ScopeRegister permits project and service registration (engx register, engx init).
	ScopeRegister = "register"

	// ScopeAdmin permits token revocation, Gate key rotation, and Guardian rule management.
	ScopeAdmin = "admin"
)
