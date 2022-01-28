package resolver

import sa "github.com/Nusantara-Muda/scholarship-api"

// RequirementResolver ...
type RequirementResolver struct {
	Requirement sa.Requirement
}

// ID ...
func (r RequirementResolver) ID() *int32 {
	ID := int32(r.Requirement.ID)
	return &ID
}

// ScholarshipID ...
func (r RequirementResolver) ScholarshipID() *int32 {
	scholarshipID := int32(r.Requirement.ScholarshipID)
	return &scholarshipID
}

//Type ...
func (r RequirementResolver) Type() *string {
	return &r.Requirement.Type
}

// Name ...
func (r RequirementResolver) Name() *string {
	return &r.Requirement.Name
}

// Value ...
func (r RequirementResolver) Value() *string {
	return &r.Requirement.Value
}
