# SERVICE-CONTRACT.md — Canon

**Service:** canon
**Domain:** Library (shared types)
**Port:** none
**ADRs:** ADR-016 (platform shared types)
**Version:** 0.1.0
**Module:** github.com/Harshmaury/Canon
**Updated:** 2026-03-18

---

## Role

Canonical types and constants module. Provides the single authoritative
definition of platform-wide types, header constants, service names, and
the nexus.yaml descriptor schema. Types and constants only — no logic,
no HTTP clients, no external dependencies.

---

## Packages

| Package              | Contents |
|----------------------|----------|
| `canon/events`       | `EventType`, `ComponentType`, `OutcomeType` constants |
| `canon/identity`     | `TraceIDHeader`, `ServiceTokenHeader`, service names, default addresses |
| `canon/descriptor`   | `Descriptor` struct, `ValidTypes` map, `StatusVerified`/`Unverified` |

---

## Key constants

```go
// identity
identity.TraceIDHeader      = "X-Trace-ID"
identity.ServiceTokenHeader = "X-Service-Token"
identity.DefaultNexusAddr   = "http://127.0.0.1:8080"
// ... through identity.DefaultSentinelAddr = "http://127.0.0.1:8087"

// events
events.EventServiceStarted / Stopped / Crashed / Healed
events.ComponentNexus / Atlas / Forge / Metrics / Navigator / Guardian / Observer / Sentinel
events.OutcomeSuccess / Failure / Deferred / Info

// descriptor
descriptor.StatusVerified / StatusUnverified
descriptor.ValidTypes  // "platform-daemon","web-api","worker","cli",...
```

---

## Dependencies

None. Canon has no external dependencies beyond the Go standard library.

---

## Guarantees

- Types and constants only — no logic, no side effects, no init() functions.
- Import Canon — never redefine these constants locally in any service.
- Changes to Canon require ADR-016 amendment.
- Canon is imported as v0.3.0 in all service go.mod files.

## Non-Responsibilities

- Canon does not make HTTP calls.
- Canon does not connect to any platform service.
- Canon does not parse nexus.yaml — that is Atlas's responsibility
  (`internal/validator/nexus_yaml.go`). Canon only defines the `Descriptor` type.
- Canon does not implement any business logic.

## Data Authority

None. Canon is a type definition library, not a data store.

## Concurrency Model

Not applicable. Compile-time constants and types — no runtime state.
