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

// LevelType is the severity level of a platform event (ADR-037).
// Import from Canon — never hardcode "info"/"warn"/"error" in any service.
type LevelType = string

const (
	LevelInfo  LevelType = "info"  // normal operation
	LevelWarn  LevelType = "warn"  // degraded but not failing
	LevelError LevelType = "error" // failure requiring attention
)

// CLI execution events — emitted by the engx CLI plan executor (ADR-043).
const (
	EventPlanStep EventType = "PLAN_STEP"
)

// ComponentEngx identifies the engx CLI as the event source.
const ComponentEngx ComponentType = "engx"

// ── Workspace topics ──────────────────────────────────────────────────────────
// ADR-002: published by Nexus internal/watcher, consumed by Atlas and Forge.
// Migrated from nexus/pkg/events/topics.go — import from Canon only.
// Never redefine these strings locally in any service.

// TopicType is the type for all platform event bus topic names.
type TopicType = string

const (
	// TopicWorkspaceFileCreated is published when a new file appears in the workspace.
	TopicWorkspaceFileCreated TopicType = "workspace.file.created"

	// TopicWorkspaceFileModified is published when an existing workspace file is written.
	TopicWorkspaceFileModified TopicType = "workspace.file.modified"

	// TopicWorkspaceFileDeleted is published when a workspace file is removed or renamed.
	TopicWorkspaceFileDeleted TopicType = "workspace.file.deleted"

	// TopicWorkspaceUpdated is published after a batch of file events settles (debounce).
	TopicWorkspaceUpdated TopicType = "workspace.updated"

	// TopicWorkspaceProjectDetected is published when the watcher finds a new project manifest.
	TopicWorkspaceProjectDetected TopicType = "workspace.project.detected"
)
