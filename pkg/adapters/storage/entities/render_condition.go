package entities

import (
	"time"
)

type RenderCondition struct {
	ID                uint     `gorm:"primaryKey"`
	TriggerQuestionID uint     `gorm:"not null"`
	TriggerOptionID   uint     `gorm:"not null"`
	TargetQuestionID  uint     `gorm:"not null"`
	TriggerQuestion   Question `gorm:"foreignKey:TriggerQuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TriggerOption     Option   `gorm:"foreignKey:TriggerOptionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TargetQuestion    Question `gorm:"foreignKey:TargetQuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
