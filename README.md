# Platform — Shared Types Module

Canonical types and constants for the engx developer platform.

## Packages

- `platform/events`     — EventType, ComponentType, OutcomeType constants
- `platform/identity`   — HTTP header constants, service names, default addresses
- `platform/descriptor` — Descriptor struct (canonical nexus.yaml schema)

## Rules (ADR-016)

- Types and constants only — no logic, no HTTP clients
- No external dependencies beyond Go standard library
- Import these packages — never redefine constants locally
- Changes require ADR-016 amendment

## Usage

```go
import (
    "github.com/Harshmaury/Platform/events"
    "github.com/Harshmaury/Platform/identity"
    "github.com/Harshmaury/Platform/descriptor"
)

// Instead of: req.Header.Set("X-Trace-ID", id)
req.Header.Set(identity.TraceIDHeader, id)

// Instead of: if ev.Type == "SERVICE_CRASHED"
if ev.Type == events.EventServiceCrashed

// Instead of: status = "verified"
status = descriptor.StatusVerified
```
