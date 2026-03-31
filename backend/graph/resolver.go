package graph

import (
	"github.com/Jnmendza/uh/backend/internal/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct{
	OrgRepo  *repository.OrganizationRepo
	UserRepo *repository.UserRepo
	MembershipRepo *repository.MembershipRepo
	MembershipTierRepo *repository.MembershipTierRepo
}
