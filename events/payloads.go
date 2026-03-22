// Package events — workspace event payload types.
// ADR-045: migrated from github.com/Harshmaury/Nexus/pkg/events.
// All consumers of workspace event payloads must import from Canon.
// Never redefine these types locally in any service.
//
// Migration guide:
//   Before: import nexusevents "github.com/Harshmaury/Nexus/pkg/events"
//   After:  import canonevents "github.com/Harshmaury/Canon/events"
//
//   nexusevents.WorkspaceFilePayload    → canonevents.WorkspaceFilePayload
//   nexusevents.WorkspaceUpdatedPayload → canonevents.WorkspaceUpdatedPayload
//   nexusevents.WorkspaceProjectPayload → canonevents.WorkspaceProjectPayload
package events

import "time"

// WorkspaceFilePayload is the payload for file created/modified/deleted events.
// Published by Nexus internal/watcher on TopicWorkspaceFileCreated,
// TopicWorkspaceFileModified, and TopicWorkspaceFileDeleted (ADR-002).
type WorkspaceFilePayload struct {
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	SizeBytes int64     `json:"size_bytes"`
	EventAt   time.Time `json:"event_at"`
}

// WorkspaceUpdatedPayload is the payload for TopicWorkspaceUpdated batch events.
// Published after a debounce window settles following a burst of file events.
type WorkspaceUpdatedPayload struct {
	WatchDir string    `json:"watch_dir"`
	EventAt  time.Time `json:"event_at"`
}

// WorkspaceProjectPayload is the payload for TopicWorkspaceProjectDetected.
// Published when the Nexus watcher discovers a new nexus.yaml manifest.
type WorkspaceProjectPayload struct {
	Path       string    `json:"path"`
	Name       string    `json:"name"`
	DetectedBy string    `json:"detected_by"`
	DetectedAt time.Time `json:"detected_at"`
}
