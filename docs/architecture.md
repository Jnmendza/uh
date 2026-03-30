# Union Hub - Data Architecture & Schema

**Database:** MongoDB
**API:** GraphQL (Golang)

## Core Philosophy

Union Hub uses a Multi-Tenant architecture. To support both standalone supporters groups and umbrella organizations with sub-factions, the core entity is an `Organization` with a self-referencing `parentId`.

## Collections

### 1. Organizations (The Tenants)

Represents a supporters group, union, or sub-group.

- `_id`: ObjectId
- `name`: String (e.g., "Frontera SD", "Chavos", "Penya Blaugrana San Diego")
- `type`: String (enum: "umbrella", "sub_group", "standalone")
- `parentId`: ObjectId (Nullable. Ref -> Organizations. Used if type is "sub_group")
- `subdomain`: String (Unique, e.g., "fronterasd")
- `branding`: Object
  - `logoUrl`: String
  - `primaryColor`: String
  - `secondaryColor`: String
- `settings`: Object
  - `followedClubId`: String (For external Sports API)
  - `followedLeagueId`: String (For external Sports API)
- `createdAt`: Timestamp
- `updatedAt`: Timestamp

### 2. Users (Global Identity)

A single identity across the entire Union Hub platform.

- `_id`: ObjectId
- `email`: String (Unique)
- `firstName`: String
- `lastName`: String
- `avatarUrl`: String
- `authProvider`: String (enum: "email", "google")
- `globalRole`: String (enum: "super_admin", "user")
- `createdAt`: Timestamp

### 3. MembershipTiers

The available levels of membership defined by an Organization.

- `_id`: ObjectId
- `organizationId`: ObjectId (Ref -> Organizations)
- `name`: String (e.g., "Free", "Gold", "Founding Member")
- `price`: Number
- `benefits`: Array of Strings
- `isActive`: Boolean

### 4. Memberships (The Tenant Link)

Links a User to an Organization and dictates their permissions for that specific group.

- `_id`: ObjectId
- `userId`: ObjectId (Ref -> Users)
- `organizationId`: ObjectId (Ref -> Organizations)
- `tierId`: ObjectId (Ref -> MembershipTiers)
- `role`: String (enum: "tenant_admin", "moderator", "member")
- `status`: String (enum: "active", "pending", "expired")
- `paymentStatus`: String (enum: "paid", "unpaid")
- `joinedAt`: Timestamp
- `expiresAt`: Timestamp (Nullable)

### 5. Channels

Chat channels within an Organization.

- `_id`: ObjectId
- `organizationId`: ObjectId (Ref -> Organizations)
- `name`: String (e.g., "matchday", "capos-only")
- `allowedTiers`: Array of ObjectIds (Refs -> MembershipTiers. Empty means all members)
- `allowedRoles`: Array of Strings (e.g., ["tenant_admin", "moderator"])

### 6. Messages

- `_id`: ObjectId
- `channelId`: ObjectId (Ref -> Channels)
- `senderId`: ObjectId (Ref -> Users)
- `content`: String
- `isPinned`: Boolean
- `createdAt`: Timestamp

### 7. Events

Matchday meetups, watch parties, or away trips.

- `_id`: ObjectId
- `organizationId`: ObjectId (Ref -> Organizations)
- `title`: String
- `description`: String
- `location`: String
- `date`: Timestamp
- `fixtureId`: String (Nullable. Link to external Sports API)
- `visibilityTiers`: Array of ObjectIds (Refs -> MembershipTiers)
