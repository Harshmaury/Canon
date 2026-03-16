# WORKFLOW-SESSION.md
# Session: PL-phase1-shared-types
# Date: 2026-03-17

## What changed — Platform shared types module (ADR-016)

New Go module with canonical types and constants.
No logic, no HTTP clients. Three packages: events, identity, descriptor.

## Setup

mkdir -p ~/workspace/projects/apps/platform
cd ~/workspace/projects/apps/platform
unzip -o /mnt/c/Users/harsh/Downloads/engx-drop/platform-types-20260317.zip -d .
go build ./...

## Usage in other services

Add to go.mod:
  require github.com/Harshmaury/Platform v0.1.0

Replace local constants:
  "X-Trace-ID"       → identity.TraceIDHeader
  "X-Service-Token"  → identity.ServiceTokenHeader
  "SERVICE_CRASHED"  → events.EventServiceCrashed
  "verified"         → descriptor.StatusVerified

## Commit

git init && git add . && \
git commit -m "feat: platform shared types module (ADR-016)" && \
git tag v0.1.0 && \
git remote add origin git@github.com:Harshmaury/Platform.git && \
git push -u origin main --tags
