// Package events defines canonical platform event type constants.
// All services that emit, consume, or filter platform events must
// import these constants — never redefine event type strings locally.
//
// ADR-016: this module is types and constants only. No logic, no HTTP clients.
package events

// EventType is the type for all platform event type strings.
type EventType = string

// Service lifecycle events — emitted by Nexus daemon.
const (
	EventServiceStarted EventType = "SERVICE_STARTED"
	EventServiceStopped EventType = "SERVICE_STOPPED"
	EventServiceCrashed EventType = "SERVICE_CRASHED"
	EventServiceHealed  EventType = "SERVICE_HEALED"
	EventStateChanged   EventType = "STATE_CHANGED"
	EventSystemAlert    EventType = "SYSTEM_ALERT"
)

// Drop intelligence events — emitted by Nexus drop system.
const (
	EventFileDropped EventType = "FILE_DROPPED"
	EventFileRouted  EventType = "FILE_ROUTED"
)

// ComponentType identifies which platform domain emitted an event.
type ComponentType = string

const (
	ComponentNexus  ComponentType = "nexus"
	ComponentDrop   ComponentType = "drop"
	ComponentSystem ComponentType = "system"
	ComponentAtlas  ComponentType = "atlas"
	ComponentForge  ComponentType = "forge"
)

// OutcomeType is the result of a platform action.
type OutcomeType = string

const (
	OutcomeSuccess  OutcomeType = "success"
	OutcomeFailure  OutcomeType = "failure"
	OutcomeDeferred OutcomeType = "deferred"
	OutcomeInfo     OutcomeType = "" // informational — no actionable result
)
