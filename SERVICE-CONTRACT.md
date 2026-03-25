// @canon-project: canon
// @canon-path: SERVICE-CONTRACT.md
# SERVICE-CONTRACT.md — Canon
# @version: 1.0.1
# @updated: 2026-03-25

**Type:** Library · **Module:** `github.com/Harshmaury/Canon` · **Domain:** Shared protocol

---

## Code

```
identity/identity.go    HTTP header constants, service names, default addresses, relay constants
events/events.go        EventType, ComponentType, OutcomeType, TopicType constants
events/payloads.go      WorkspaceFilePayload, WorkspaceUpdatedPayload, WorkspaceProjectPayload, SystemAlertPayload
descriptor/descriptor.go  Descriptor struct, ValidTypes map, StatusVerified/Unverified
```

---

## Contract

No HTTP surface. Compile-time only.

**`identity` package — key constants:**
```
TraceIDHeader       = "X-Trace-ID"
ServiceTokenHeader  = "X-Service-Token"
IdentityTokenHeader = "X-Identity-Token"
RelayTokenHeader    = "X-Relay-Token"
DefaultNexusAddr    = "http://127.0.0.1:8080"   (through DefaultGateAddr :8088)
```

**`descriptor.ValidTypes`:** `platform-daemon`, `web-api`, `worker`, `cli`, `database`, `message-broker`, `gateway`, `library`, `automation`, `ml-service`, `governance`, `tool`.

**Versioning:** any additive constant → patch. Any rename or removal → major, requires ADR-016 amendment.

---

## Control

No runtime behavior. Compile-time constants and types. No `init()` functions, no side effects.

---

## Context

Single source of truth for all protocol constants on the platform. Arbiter A-C-001/002 blocks any service that redefines these locally. No external dependencies beyond Go stdlib.
