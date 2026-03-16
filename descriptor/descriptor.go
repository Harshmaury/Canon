// Package descriptor defines the canonical nexus.yaml project descriptor schema.
// This is the authoritative Go representation of the nexus.yaml contract (ADR-009).
//
// Atlas validator, Sentinel, Guardian, and any future service that reads
// nexus.yaml must import this package rather than defining their own struct.
//
// ADR-016: types only. Parsing logic lives in Atlas internal/validator.
package descriptor

// Descriptor is the canonical Go representation of nexus.yaml.
// Required fields: Name, ID, Type, Language, Version.
// Optional fields: Keywords, Capabilities, DependsOn, Runtime.
type Descriptor struct {
	Name         string   `yaml:"name"         json:"name"`
	ID           string   `yaml:"id"           json:"id"`
	Type         string   `yaml:"type"         json:"type"`
	Language     string   `yaml:"language"     json:"language"`
	Version      string   `yaml:"version"      json:"version"`
	Keywords     []string `yaml:"keywords"     json:"keywords"`
	Capabilities []string `yaml:"capabilities" json:"capabilities"`
	DependsOn    []string `yaml:"depends_on"   json:"depends_on"`
	Runtime      Runtime  `yaml:"runtime"      json:"runtime"`
}

// Runtime describes how the project is executed.
type Runtime struct {
	Provider string `yaml:"provider" json:"provider"` // process | docker | k8s
	Port     int    `yaml:"port"     json:"port"`
}

// ValidTypes is the set of allowed project type values (ADR-009).
// All services that validate nexus.yaml must use this set.
var ValidTypes = map[string]bool{
	"platform-daemon": true,
	"web-api":         true,
	"worker":          true,
	"cli":             true,
	"database":        true,
	"message-broker":  true,
	"gateway":         true,
	"library":         true,
	"automation":      true,
	"ml-service":      true,
}

// StatusVerified and StatusUnverified are the canonical project status values.
const (
	StatusVerified   = "verified"
	StatusUnverified = "unverified"
)
