// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package geodeticclient

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/datumforge/geodetic/pkg/enums"
)

// CreateDatabaseInput is used for create Database object.
// Input was generated by ent.
type CreateDatabaseInput struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	// the ID of the organization
	OrganizationID string `json:"organizationID"`
	// the name to the database
	Name string `json:"name"`
	// the geo location of the database
	Geo *string `json:"geo,omitempty"`
	// the DSN to the database
	Dsn string `json:"dsn"`
	// the auth token used to connect to the database
	Token *string `json:"token,omitempty"`
	// status of the database
	Status *enums.DatabaseStatus `json:"status,omitempty"`
	// provider of the database
	Provider *enums.DatabaseProvider `json:"provider,omitempty"`
	GroupID  string                  `json:"groupID"`
}

// CreateGroupInput is used for create Group object.
// Input was generated by ent.
type CreateGroupInput struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	// the name of the group in turso
	Name string `json:"name"`
	// the description of the group
	Description *string `json:"description,omitempty"`
	// the primary of the group
	PrimaryLocation string `json:"primaryLocation"`
	// the replica locations of the group
	Locations []string `json:"locations,omitempty"`
	// the auth token used to connect to the group
	Token *string `json:"token,omitempty"`
	// region the group
	Region      *enums.Region `json:"region,omitempty"`
	DatabaseIDs []string      `json:"databaseIDs,omitempty"`
}

type Database struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	DeletedBy *string    `json:"deletedBy,omitempty"`
	// the ID of the organization
	OrganizationID string `json:"organizationID"`
	// the name to the database
	Name string `json:"name"`
	// the geo location of the database
	Geo *string `json:"geo,omitempty"`
	// the DSN to the database
	Dsn string `json:"dsn"`
	// the ID of the group
	GroupID string `json:"groupID"`
	// status of the database
	Status enums.DatabaseStatus `json:"status"`
	// provider of the database
	Provider enums.DatabaseProvider `json:"provider"`
	Group    *Group                 `json:"group"`
}

func (Database) IsNode() {}

// A connection to a list of items.
type DatabaseConnection struct {
	// A list of edges.
	Edges []*DatabaseEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// Return response for createDatabase mutation
type DatabaseCreatePayload struct {
	// Created database
	Database *Database `json:"database"`
}

// Return response for deleteDatabase mutation
type DatabaseDeletePayload struct {
	// Deleted database ID
	DeletedID string `json:"deletedID"`
}

// An edge in a connection.
type DatabaseEdge struct {
	// The item at the end of the edge.
	Node *Database `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Return response for updateDatabase mutation
type DatabaseUpdatePayload struct {
	// Updated database
	Database *Database `json:"database"`
}

// DatabaseWhereInput is used for filtering Database objects.
// Input was generated by ent.
type DatabaseWhereInput struct {
	Not *DatabaseWhereInput   `json:"not,omitempty"`
	And []*DatabaseWhereInput `json:"and,omitempty"`
	Or  []*DatabaseWhereInput `json:"or,omitempty"`
	// id field predicates
	ID             *string  `json:"id,omitempty"`
	IDNeq          *string  `json:"idNEQ,omitempty"`
	IDIn           []string `json:"idIn,omitempty"`
	IDNotIn        []string `json:"idNotIn,omitempty"`
	IDGt           *string  `json:"idGT,omitempty"`
	IDGte          *string  `json:"idGTE,omitempty"`
	IDLt           *string  `json:"idLT,omitempty"`
	IDLte          *string  `json:"idLTE,omitempty"`
	IDEqualFold    *string  `json:"idEqualFold,omitempty"`
	IDContainsFold *string  `json:"idContainsFold,omitempty"`
	// created_at field predicates
	CreatedAt       *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq    *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn     []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn  []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt     *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte    *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt     *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte    *time.Time   `json:"createdAtLTE,omitempty"`
	CreatedAtIsNil  *bool        `json:"createdAtIsNil,omitempty"`
	CreatedAtNotNil *bool        `json:"createdAtNotNil,omitempty"`
	// updated_at field predicates
	UpdatedAt       *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq    *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn     []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn  []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt     *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte    *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt     *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte    *time.Time   `json:"updatedAtLTE,omitempty"`
	UpdatedAtIsNil  *bool        `json:"updatedAtIsNil,omitempty"`
	UpdatedAtNotNil *bool        `json:"updatedAtNotNil,omitempty"`
	// created_by field predicates
	CreatedBy             *string  `json:"createdBy,omitempty"`
	CreatedByNeq          *string  `json:"createdByNEQ,omitempty"`
	CreatedByIn           []string `json:"createdByIn,omitempty"`
	CreatedByNotIn        []string `json:"createdByNotIn,omitempty"`
	CreatedByGt           *string  `json:"createdByGT,omitempty"`
	CreatedByGte          *string  `json:"createdByGTE,omitempty"`
	CreatedByLt           *string  `json:"createdByLT,omitempty"`
	CreatedByLte          *string  `json:"createdByLTE,omitempty"`
	CreatedByContains     *string  `json:"createdByContains,omitempty"`
	CreatedByHasPrefix    *string  `json:"createdByHasPrefix,omitempty"`
	CreatedByHasSuffix    *string  `json:"createdByHasSuffix,omitempty"`
	CreatedByIsNil        *bool    `json:"createdByIsNil,omitempty"`
	CreatedByNotNil       *bool    `json:"createdByNotNil,omitempty"`
	CreatedByEqualFold    *string  `json:"createdByEqualFold,omitempty"`
	CreatedByContainsFold *string  `json:"createdByContainsFold,omitempty"`
	// updated_by field predicates
	UpdatedBy             *string  `json:"updatedBy,omitempty"`
	UpdatedByNeq          *string  `json:"updatedByNEQ,omitempty"`
	UpdatedByIn           []string `json:"updatedByIn,omitempty"`
	UpdatedByNotIn        []string `json:"updatedByNotIn,omitempty"`
	UpdatedByGt           *string  `json:"updatedByGT,omitempty"`
	UpdatedByGte          *string  `json:"updatedByGTE,omitempty"`
	UpdatedByLt           *string  `json:"updatedByLT,omitempty"`
	UpdatedByLte          *string  `json:"updatedByLTE,omitempty"`
	UpdatedByContains     *string  `json:"updatedByContains,omitempty"`
	UpdatedByHasPrefix    *string  `json:"updatedByHasPrefix,omitempty"`
	UpdatedByHasSuffix    *string  `json:"updatedByHasSuffix,omitempty"`
	UpdatedByIsNil        *bool    `json:"updatedByIsNil,omitempty"`
	UpdatedByNotNil       *bool    `json:"updatedByNotNil,omitempty"`
	UpdatedByEqualFold    *string  `json:"updatedByEqualFold,omitempty"`
	UpdatedByContainsFold *string  `json:"updatedByContainsFold,omitempty"`
	// deleted_at field predicates
	DeletedAt       *time.Time   `json:"deletedAt,omitempty"`
	DeletedAtNeq    *time.Time   `json:"deletedAtNEQ,omitempty"`
	DeletedAtIn     []*time.Time `json:"deletedAtIn,omitempty"`
	DeletedAtNotIn  []*time.Time `json:"deletedAtNotIn,omitempty"`
	DeletedAtGt     *time.Time   `json:"deletedAtGT,omitempty"`
	DeletedAtGte    *time.Time   `json:"deletedAtGTE,omitempty"`
	DeletedAtLt     *time.Time   `json:"deletedAtLT,omitempty"`
	DeletedAtLte    *time.Time   `json:"deletedAtLTE,omitempty"`
	DeletedAtIsNil  *bool        `json:"deletedAtIsNil,omitempty"`
	DeletedAtNotNil *bool        `json:"deletedAtNotNil,omitempty"`
	// deleted_by field predicates
	DeletedBy             *string  `json:"deletedBy,omitempty"`
	DeletedByNeq          *string  `json:"deletedByNEQ,omitempty"`
	DeletedByIn           []string `json:"deletedByIn,omitempty"`
	DeletedByNotIn        []string `json:"deletedByNotIn,omitempty"`
	DeletedByGt           *string  `json:"deletedByGT,omitempty"`
	DeletedByGte          *string  `json:"deletedByGTE,omitempty"`
	DeletedByLt           *string  `json:"deletedByLT,omitempty"`
	DeletedByLte          *string  `json:"deletedByLTE,omitempty"`
	DeletedByContains     *string  `json:"deletedByContains,omitempty"`
	DeletedByHasPrefix    *string  `json:"deletedByHasPrefix,omitempty"`
	DeletedByHasSuffix    *string  `json:"deletedByHasSuffix,omitempty"`
	DeletedByIsNil        *bool    `json:"deletedByIsNil,omitempty"`
	DeletedByNotNil       *bool    `json:"deletedByNotNil,omitempty"`
	DeletedByEqualFold    *string  `json:"deletedByEqualFold,omitempty"`
	DeletedByContainsFold *string  `json:"deletedByContainsFold,omitempty"`
	// organization_id field predicates
	OrganizationID             *string  `json:"organizationID,omitempty"`
	OrganizationIDNeq          *string  `json:"organizationIDNEQ,omitempty"`
	OrganizationIDIn           []string `json:"organizationIDIn,omitempty"`
	OrganizationIDNotIn        []string `json:"organizationIDNotIn,omitempty"`
	OrganizationIDGt           *string  `json:"organizationIDGT,omitempty"`
	OrganizationIDGte          *string  `json:"organizationIDGTE,omitempty"`
	OrganizationIDLt           *string  `json:"organizationIDLT,omitempty"`
	OrganizationIDLte          *string  `json:"organizationIDLTE,omitempty"`
	OrganizationIDContains     *string  `json:"organizationIDContains,omitempty"`
	OrganizationIDHasPrefix    *string  `json:"organizationIDHasPrefix,omitempty"`
	OrganizationIDHasSuffix    *string  `json:"organizationIDHasSuffix,omitempty"`
	OrganizationIDEqualFold    *string  `json:"organizationIDEqualFold,omitempty"`
	OrganizationIDContainsFold *string  `json:"organizationIDContainsFold,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	// geo field predicates
	Geo             *string  `json:"geo,omitempty"`
	GeoNeq          *string  `json:"geoNEQ,omitempty"`
	GeoIn           []string `json:"geoIn,omitempty"`
	GeoNotIn        []string `json:"geoNotIn,omitempty"`
	GeoGt           *string  `json:"geoGT,omitempty"`
	GeoGte          *string  `json:"geoGTE,omitempty"`
	GeoLt           *string  `json:"geoLT,omitempty"`
	GeoLte          *string  `json:"geoLTE,omitempty"`
	GeoContains     *string  `json:"geoContains,omitempty"`
	GeoHasPrefix    *string  `json:"geoHasPrefix,omitempty"`
	GeoHasSuffix    *string  `json:"geoHasSuffix,omitempty"`
	GeoIsNil        *bool    `json:"geoIsNil,omitempty"`
	GeoNotNil       *bool    `json:"geoNotNil,omitempty"`
	GeoEqualFold    *string  `json:"geoEqualFold,omitempty"`
	GeoContainsFold *string  `json:"geoContainsFold,omitempty"`
	// dsn field predicates
	Dsn             *string  `json:"dsn,omitempty"`
	DsnNeq          *string  `json:"dsnNEQ,omitempty"`
	DsnIn           []string `json:"dsnIn,omitempty"`
	DsnNotIn        []string `json:"dsnNotIn,omitempty"`
	DsnGt           *string  `json:"dsnGT,omitempty"`
	DsnGte          *string  `json:"dsnGTE,omitempty"`
	DsnLt           *string  `json:"dsnLT,omitempty"`
	DsnLte          *string  `json:"dsnLTE,omitempty"`
	DsnContains     *string  `json:"dsnContains,omitempty"`
	DsnHasPrefix    *string  `json:"dsnHasPrefix,omitempty"`
	DsnHasSuffix    *string  `json:"dsnHasSuffix,omitempty"`
	DsnEqualFold    *string  `json:"dsnEqualFold,omitempty"`
	DsnContainsFold *string  `json:"dsnContainsFold,omitempty"`
	// group_id field predicates
	GroupID             *string  `json:"groupID,omitempty"`
	GroupIDNeq          *string  `json:"groupIDNEQ,omitempty"`
	GroupIDIn           []string `json:"groupIDIn,omitempty"`
	GroupIDNotIn        []string `json:"groupIDNotIn,omitempty"`
	GroupIDGt           *string  `json:"groupIDGT,omitempty"`
	GroupIDGte          *string  `json:"groupIDGTE,omitempty"`
	GroupIDLt           *string  `json:"groupIDLT,omitempty"`
	GroupIDLte          *string  `json:"groupIDLTE,omitempty"`
	GroupIDContains     *string  `json:"groupIDContains,omitempty"`
	GroupIDHasPrefix    *string  `json:"groupIDHasPrefix,omitempty"`
	GroupIDHasSuffix    *string  `json:"groupIDHasSuffix,omitempty"`
	GroupIDEqualFold    *string  `json:"groupIDEqualFold,omitempty"`
	GroupIDContainsFold *string  `json:"groupIDContainsFold,omitempty"`
	// status field predicates
	Status      *enums.DatabaseStatus  `json:"status,omitempty"`
	StatusNeq   *enums.DatabaseStatus  `json:"statusNEQ,omitempty"`
	StatusIn    []enums.DatabaseStatus `json:"statusIn,omitempty"`
	StatusNotIn []enums.DatabaseStatus `json:"statusNotIn,omitempty"`
	// provider field predicates
	Provider      *enums.DatabaseProvider  `json:"provider,omitempty"`
	ProviderNeq   *enums.DatabaseProvider  `json:"providerNEQ,omitempty"`
	ProviderIn    []enums.DatabaseProvider `json:"providerIn,omitempty"`
	ProviderNotIn []enums.DatabaseProvider `json:"providerNotIn,omitempty"`
	// group edge predicates
	HasGroup     *bool              `json:"hasGroup,omitempty"`
	HasGroupWith []*GroupWhereInput `json:"hasGroupWith,omitempty"`
}

type Group struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	DeletedBy *string    `json:"deletedBy,omitempty"`
	// the name of the group in turso
	Name string `json:"name"`
	// the description of the group
	Description *string `json:"description,omitempty"`
	// the primary of the group
	PrimaryLocation string `json:"primaryLocation"`
	// the replica locations of the group
	Locations []string `json:"locations,omitempty"`
	// region the group
	Region    enums.Region `json:"region"`
	Databases []*Database  `json:"databases,omitempty"`
}

func (Group) IsNode() {}

// A connection to a list of items.
type GroupConnection struct {
	// A list of edges.
	Edges []*GroupEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// Return response for createGroup mutation
type GroupCreatePayload struct {
	// Created group
	Group *Group `json:"group"`
}

// Return response for deleteGroup mutation
type GroupDeletePayload struct {
	// Deleted group ID
	DeletedID string `json:"deletedID"`
}

// An edge in a connection.
type GroupEdge struct {
	// The item at the end of the edge.
	Node *Group `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Return response for updateGroup mutation
type GroupUpdatePayload struct {
	// Updated group
	Group *Group `json:"group"`
}

// GroupWhereInput is used for filtering Group objects.
// Input was generated by ent.
type GroupWhereInput struct {
	Not *GroupWhereInput   `json:"not,omitempty"`
	And []*GroupWhereInput `json:"and,omitempty"`
	Or  []*GroupWhereInput `json:"or,omitempty"`
	// id field predicates
	ID             *string  `json:"id,omitempty"`
	IDNeq          *string  `json:"idNEQ,omitempty"`
	IDIn           []string `json:"idIn,omitempty"`
	IDNotIn        []string `json:"idNotIn,omitempty"`
	IDGt           *string  `json:"idGT,omitempty"`
	IDGte          *string  `json:"idGTE,omitempty"`
	IDLt           *string  `json:"idLT,omitempty"`
	IDLte          *string  `json:"idLTE,omitempty"`
	IDEqualFold    *string  `json:"idEqualFold,omitempty"`
	IDContainsFold *string  `json:"idContainsFold,omitempty"`
	// created_at field predicates
	CreatedAt       *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq    *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn     []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn  []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt     *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte    *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt     *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte    *time.Time   `json:"createdAtLTE,omitempty"`
	CreatedAtIsNil  *bool        `json:"createdAtIsNil,omitempty"`
	CreatedAtNotNil *bool        `json:"createdAtNotNil,omitempty"`
	// updated_at field predicates
	UpdatedAt       *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq    *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn     []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn  []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt     *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte    *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt     *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte    *time.Time   `json:"updatedAtLTE,omitempty"`
	UpdatedAtIsNil  *bool        `json:"updatedAtIsNil,omitempty"`
	UpdatedAtNotNil *bool        `json:"updatedAtNotNil,omitempty"`
	// created_by field predicates
	CreatedBy             *string  `json:"createdBy,omitempty"`
	CreatedByNeq          *string  `json:"createdByNEQ,omitempty"`
	CreatedByIn           []string `json:"createdByIn,omitempty"`
	CreatedByNotIn        []string `json:"createdByNotIn,omitempty"`
	CreatedByGt           *string  `json:"createdByGT,omitempty"`
	CreatedByGte          *string  `json:"createdByGTE,omitempty"`
	CreatedByLt           *string  `json:"createdByLT,omitempty"`
	CreatedByLte          *string  `json:"createdByLTE,omitempty"`
	CreatedByContains     *string  `json:"createdByContains,omitempty"`
	CreatedByHasPrefix    *string  `json:"createdByHasPrefix,omitempty"`
	CreatedByHasSuffix    *string  `json:"createdByHasSuffix,omitempty"`
	CreatedByIsNil        *bool    `json:"createdByIsNil,omitempty"`
	CreatedByNotNil       *bool    `json:"createdByNotNil,omitempty"`
	CreatedByEqualFold    *string  `json:"createdByEqualFold,omitempty"`
	CreatedByContainsFold *string  `json:"createdByContainsFold,omitempty"`
	// updated_by field predicates
	UpdatedBy             *string  `json:"updatedBy,omitempty"`
	UpdatedByNeq          *string  `json:"updatedByNEQ,omitempty"`
	UpdatedByIn           []string `json:"updatedByIn,omitempty"`
	UpdatedByNotIn        []string `json:"updatedByNotIn,omitempty"`
	UpdatedByGt           *string  `json:"updatedByGT,omitempty"`
	UpdatedByGte          *string  `json:"updatedByGTE,omitempty"`
	UpdatedByLt           *string  `json:"updatedByLT,omitempty"`
	UpdatedByLte          *string  `json:"updatedByLTE,omitempty"`
	UpdatedByContains     *string  `json:"updatedByContains,omitempty"`
	UpdatedByHasPrefix    *string  `json:"updatedByHasPrefix,omitempty"`
	UpdatedByHasSuffix    *string  `json:"updatedByHasSuffix,omitempty"`
	UpdatedByIsNil        *bool    `json:"updatedByIsNil,omitempty"`
	UpdatedByNotNil       *bool    `json:"updatedByNotNil,omitempty"`
	UpdatedByEqualFold    *string  `json:"updatedByEqualFold,omitempty"`
	UpdatedByContainsFold *string  `json:"updatedByContainsFold,omitempty"`
	// deleted_at field predicates
	DeletedAt       *time.Time   `json:"deletedAt,omitempty"`
	DeletedAtNeq    *time.Time   `json:"deletedAtNEQ,omitempty"`
	DeletedAtIn     []*time.Time `json:"deletedAtIn,omitempty"`
	DeletedAtNotIn  []*time.Time `json:"deletedAtNotIn,omitempty"`
	DeletedAtGt     *time.Time   `json:"deletedAtGT,omitempty"`
	DeletedAtGte    *time.Time   `json:"deletedAtGTE,omitempty"`
	DeletedAtLt     *time.Time   `json:"deletedAtLT,omitempty"`
	DeletedAtLte    *time.Time   `json:"deletedAtLTE,omitempty"`
	DeletedAtIsNil  *bool        `json:"deletedAtIsNil,omitempty"`
	DeletedAtNotNil *bool        `json:"deletedAtNotNil,omitempty"`
	// deleted_by field predicates
	DeletedBy             *string  `json:"deletedBy,omitempty"`
	DeletedByNeq          *string  `json:"deletedByNEQ,omitempty"`
	DeletedByIn           []string `json:"deletedByIn,omitempty"`
	DeletedByNotIn        []string `json:"deletedByNotIn,omitempty"`
	DeletedByGt           *string  `json:"deletedByGT,omitempty"`
	DeletedByGte          *string  `json:"deletedByGTE,omitempty"`
	DeletedByLt           *string  `json:"deletedByLT,omitempty"`
	DeletedByLte          *string  `json:"deletedByLTE,omitempty"`
	DeletedByContains     *string  `json:"deletedByContains,omitempty"`
	DeletedByHasPrefix    *string  `json:"deletedByHasPrefix,omitempty"`
	DeletedByHasSuffix    *string  `json:"deletedByHasSuffix,omitempty"`
	DeletedByIsNil        *bool    `json:"deletedByIsNil,omitempty"`
	DeletedByNotNil       *bool    `json:"deletedByNotNil,omitempty"`
	DeletedByEqualFold    *string  `json:"deletedByEqualFold,omitempty"`
	DeletedByContainsFold *string  `json:"deletedByContainsFold,omitempty"`
	// name field predicates
	Name             *string  `json:"name,omitempty"`
	NameNeq          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGt           *string  `json:"nameGT,omitempty"`
	NameGte          *string  `json:"nameGTE,omitempty"`
	NameLt           *string  `json:"nameLT,omitempty"`
	NameLte          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
	// description field predicates
	Description             *string  `json:"description,omitempty"`
	DescriptionNeq          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGt           *string  `json:"descriptionGT,omitempty"`
	DescriptionGte          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLt           *string  `json:"descriptionLT,omitempty"`
	DescriptionLte          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionIsNil        *bool    `json:"descriptionIsNil,omitempty"`
	DescriptionNotNil       *bool    `json:"descriptionNotNil,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`
	// primary_location field predicates
	PrimaryLocation             *string  `json:"primaryLocation,omitempty"`
	PrimaryLocationNeq          *string  `json:"primaryLocationNEQ,omitempty"`
	PrimaryLocationIn           []string `json:"primaryLocationIn,omitempty"`
	PrimaryLocationNotIn        []string `json:"primaryLocationNotIn,omitempty"`
	PrimaryLocationGt           *string  `json:"primaryLocationGT,omitempty"`
	PrimaryLocationGte          *string  `json:"primaryLocationGTE,omitempty"`
	PrimaryLocationLt           *string  `json:"primaryLocationLT,omitempty"`
	PrimaryLocationLte          *string  `json:"primaryLocationLTE,omitempty"`
	PrimaryLocationContains     *string  `json:"primaryLocationContains,omitempty"`
	PrimaryLocationHasPrefix    *string  `json:"primaryLocationHasPrefix,omitempty"`
	PrimaryLocationHasSuffix    *string  `json:"primaryLocationHasSuffix,omitempty"`
	PrimaryLocationEqualFold    *string  `json:"primaryLocationEqualFold,omitempty"`
	PrimaryLocationContainsFold *string  `json:"primaryLocationContainsFold,omitempty"`
	// region field predicates
	Region      *enums.Region  `json:"region,omitempty"`
	RegionNeq   *enums.Region  `json:"regionNEQ,omitempty"`
	RegionIn    []enums.Region `json:"regionIn,omitempty"`
	RegionNotIn []enums.Region `json:"regionNotIn,omitempty"`
	// databases edge predicates
	HasDatabases     *bool                 `json:"hasDatabases,omitempty"`
	HasDatabasesWith []*DatabaseWhereInput `json:"hasDatabasesWith,omitempty"`
}

type Mutation struct {
}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor,omitempty"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor,omitempty"`
}

type Query struct {
}

// UpdateDatabaseInput is used for update Database object.
// Input was generated by ent.
type UpdateDatabaseInput struct {
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	ClearUpdatedAt *bool      `json:"clearUpdatedAt,omitempty"`
	UpdatedBy      *string    `json:"updatedBy,omitempty"`
	ClearUpdatedBy *bool      `json:"clearUpdatedBy,omitempty"`
	// the ID of the organization
	OrganizationID *string `json:"organizationID,omitempty"`
	// the name to the database
	Name *string `json:"name,omitempty"`
	// the geo location of the database
	Geo      *string `json:"geo,omitempty"`
	ClearGeo *bool   `json:"clearGeo,omitempty"`
	// the DSN to the database
	Dsn *string `json:"dsn,omitempty"`
	// the auth token used to connect to the database
	Token      *string `json:"token,omitempty"`
	ClearToken *bool   `json:"clearToken,omitempty"`
	// status of the database
	Status *enums.DatabaseStatus `json:"status,omitempty"`
	// provider of the database
	Provider *enums.DatabaseProvider `json:"provider,omitempty"`
	GroupID  *string                 `json:"groupID,omitempty"`
}

// UpdateGroupInput is used for update Group object.
// Input was generated by ent.
type UpdateGroupInput struct {
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	ClearUpdatedAt *bool      `json:"clearUpdatedAt,omitempty"`
	UpdatedBy      *string    `json:"updatedBy,omitempty"`
	ClearUpdatedBy *bool      `json:"clearUpdatedBy,omitempty"`
	// the name of the group in turso
	Name *string `json:"name,omitempty"`
	// the description of the group
	Description      *string `json:"description,omitempty"`
	ClearDescription *bool   `json:"clearDescription,omitempty"`
	// the primary of the group
	PrimaryLocation *string `json:"primaryLocation,omitempty"`
	// the replica locations of the group
	Locations       []string `json:"locations,omitempty"`
	AppendLocations []string `json:"appendLocations,omitempty"`
	ClearLocations  *bool    `json:"clearLocations,omitempty"`
	// the auth token used to connect to the group
	Token      *string `json:"token,omitempty"`
	ClearToken *bool   `json:"clearToken,omitempty"`
	// region the group
	Region            *enums.Region `json:"region,omitempty"`
	AddDatabaseIDs    []string      `json:"addDatabaseIDs,omitempty"`
	RemoveDatabaseIDs []string      `json:"removeDatabaseIDs,omitempty"`
	ClearDatabases    *bool         `json:"clearDatabases,omitempty"`
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
