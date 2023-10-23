package models

import "gorm.io/gorm"

type Scripts struct {
	gorm.Model
	ScriptID uint
	Trigger  string
}

func (s Scripts) IsValid() bool {

	// Validates the State
	// Additional validation and hooks for the State validation can be added here
	// WARNING: Validation should be scoped to the State

	return true
}
