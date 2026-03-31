package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Strong types for Enums based on GraphQL/Go Architecture
type OrganizationType string

const (
	OrgTypeUmbrella   OrganizationType = "umbrella"
	OrgTypeSubGroup   OrganizationType = "sub_group"
	OrgTypeStandalone OrganizationType = "standalone"
)

type AuthProvider string

const (
	AuthProviderEmail  AuthProvider = "email"
	AuthProviderGoogle AuthProvider = "google"
)

type GlobalRole string

const (
	GlobalRoleSuperAdmin GlobalRole = "super_admin"
	GlobalRoleUser       GlobalRole = "user"
)

type Role string

const (
	RoleTenantAdmin Role = "tenant_admin"
	RoleModerator   Role = "moderator"
	RoleMember      Role = "member"
)

type MembershipStatus string

const (
	MembershipStatusActive  MembershipStatus = "active"
	MembershipStatusPending MembershipStatus = "pending"
	MembershipStatusExpired MembershipStatus = "expired"
)

type PaymentStatus string

const (
	PaymentStatusPaid   PaymentStatus = "paid"
	PaymentStatusUnpaid PaymentStatus = "unpaid"
)

// Embedded Generic Structs
type Branding struct {
	LogoURL        string `bson:"logoUrl" json:"logoUrl"`
	PrimaryColor   string `bson:"primaryColor" json:"primaryColor"`
	SecondaryColor string `bson:"secondaryColor" json:"secondaryColor"`
}

type Settings struct {
	FollowedClubID   string `bson:"followedClubId" json:"followedClubId"`
	FollowedLeagueID string `bson:"followedLeagueId" json:"followedLeagueId"`
}

// Collections

type Organization struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Name      string              `bson:"name" json:"name"`
	Type      OrganizationType    `bson:"type" json:"type"`
	ParentID  *primitive.ObjectID `bson:"parentId,omitempty" json:"parentId,omitempty"`
	Subdomain string              `bson:"subdomain" json:"subdomain"`
	Branding  Branding            `bson:"branding" json:"branding"`
	Settings  Settings            `bson:"settings" json:"settings"`
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email        string             `bson:"email" json:"email"`
	FirstName    string             `bson:"firstName" json:"firstName"`
	LastName     string             `bson:"lastName" json:"lastName"`
	AvatarURL    string             `bson:"avatarUrl" json:"avatarUrl"`
	AuthProvider AuthProvider       `bson:"authProvider" json:"authProvider"`
	GlobalRole   GlobalRole         `bson:"globalRole" json:"globalRole"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}

type MembershipTier struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrganizationID primitive.ObjectID `bson:"organizationId" json:"organizationId"`
	Name           string             `bson:"name" json:"name"`
	Price          float64            `bson:"price" json:"price"`
	Benefits       []string           `bson:"benefits" json:"benefits"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
}

type Membership struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         primitive.ObjectID `bson:"userId" json:"userId"`
	OrganizationID primitive.ObjectID `bson:"organizationId" json:"organizationId"`
	TierID         primitive.ObjectID `bson:"tierId" json:"tierId"`
	Role           Role               `bson:"role" json:"role"`
	Status         MembershipStatus   `bson:"status" json:"status"`
	PaymentStatus  PaymentStatus      `bson:"paymentStatus" json:"paymentStatus"`
	JoinedAt       time.Time          `bson:"joinedAt" json:"joinedAt"`
	ExpiresAt      *time.Time         `bson:"expiresAt,omitempty" json:"expiresAt,omitempty"`
}

type Channel struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	OrganizationID primitive.ObjectID   `bson:"organizationId" json:"organizationId"`
	Name           string               `bson:"name" json:"name"`
	AllowedTiers   []primitive.ObjectID `bson:"allowedTiers" json:"allowedTiers"`
	AllowedRoles   []Role               `bson:"allowedRoles" json:"allowedRoles"`
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChannelID primitive.ObjectID `bson:"channelId" json:"channelId"`
	SenderID  primitive.ObjectID `bson:"senderId" json:"senderId"`
	Content   string             `bson:"content" json:"content"`
	IsPinned  bool               `bson:"isPinned" json:"isPinned"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type Event struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	OrganizationID  primitive.ObjectID   `bson:"organizationId" json:"organizationId"`
	Title           string               `bson:"title" json:"title"`
	Description     string               `bson:"description" json:"description"`
	Location        string               `bson:"location" json:"location"`
	Date            time.Time            `bson:"date" json:"date"`
	FixtureID       *string              `bson:"fixtureId,omitempty" json:"fixtureId,omitempty"`
	VisibilityTiers []primitive.ObjectID `bson:"visibilityTiers" json:"visibilityTiers"`
}
