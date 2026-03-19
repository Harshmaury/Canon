# WORKFLOW-SESSION.md
# Session: CN-shared-types-v0.3.0
# Date: 2026-03-19

## What changed — Canon shared types module (ADR-016)

Canonical Go module for all platform-wide constants and types.
No logic, no HTTP clients. Three packages: events, identity, descriptor.
Module renamed from Platform → Canon at v0.3.0; descriptor package added.

## New files
- `identity/identity.go`     — TraceIDHeader, ServiceTokenHeader, DefaultXxxAddr, ServiceXxx name constants
- `events/events.go`         — EventType constants, ComponentType constants, OutcomeType constants
- `descriptor/descriptor.go` — Descriptor struct, Runtime struct, ValidTypes map, StatusVerified/Unverified

## Setup

```bash
mkdir -p ~/workspace/projects/apps/canon
cd ~/workspace/projects/apps/canon
unzip -o /mnt/c/Users/harsh/Downloads/engx-drop/canon-v0.3.0-20260319.zip -d .
go build ./...
```

## Usage in other services

Add to go.mod:
  require github.com/Harshmaury/Canon v0.3.0

Replace local constants:
  "X-Trace-ID"       → identity.TraceIDHeader
  "X-Service-Token"  → identity.ServiceTokenHeader
  "SERVICE_CRASHED"  → events.EventServiceCrashed
  "FILE_DROPPED"     → events.EventFileDropped
  "verified"         → descriptor.StatusVerified

## Commit

```bash
cd ~/workspace/projects/apps/canon
git add identity/identity.go events/events.go descriptor/descriptor.go nexus.yaml WORKFLOW-SESSION.md && \
git commit -m "feat: Canon v0.3.0 — descriptor package, renamed from Platform" && \
git tag v0.3.0 && \
git push origin main --tags
```
