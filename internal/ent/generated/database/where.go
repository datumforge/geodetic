// Code generated by ent, DO NOT EDIT.

package database

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/geodetic/internal/ent/generated/predicate"
	"github.com/datumforge/geodetic/pkg/enums"

	"github.com/datumforge/geodetic/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldDeletedBy, v))
}

// MappingID applies equality check predicate on the "mapping_id" field. It's identical to MappingIDEQ.
func MappingID(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldMappingID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldOrganizationID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldName, v))
}

// Geo applies equality check predicate on the "geo" field. It's identical to GeoEQ.
func Geo(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldGeo, v))
}

// Dsn applies equality check predicate on the "dsn" field. It's identical to DsnEQ.
func Dsn(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldDsn, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldGroupID, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldToken, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Database {
	return predicate.Database(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Database {
	return predicate.Database(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Database {
	return predicate.Database(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Database {
	return predicate.Database(sql.FieldNotNull(FieldUpdatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Database {
	return predicate.Database(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Database {
	return predicate.Database(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Database {
	return predicate.Database(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Database {
	return predicate.Database(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Database {
	return predicate.Database(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Database {
	return predicate.Database(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.Database {
	return predicate.Database(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.Database {
	return predicate.Database(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldDeletedBy, v))
}

// MappingIDEQ applies the EQ predicate on the "mapping_id" field.
func MappingIDEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldMappingID, v))
}

// MappingIDNEQ applies the NEQ predicate on the "mapping_id" field.
func MappingIDNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldMappingID, v))
}

// MappingIDIn applies the In predicate on the "mapping_id" field.
func MappingIDIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldMappingID, vs...))
}

// MappingIDNotIn applies the NotIn predicate on the "mapping_id" field.
func MappingIDNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldMappingID, vs...))
}

// MappingIDGT applies the GT predicate on the "mapping_id" field.
func MappingIDGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldMappingID, v))
}

// MappingIDGTE applies the GTE predicate on the "mapping_id" field.
func MappingIDGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldMappingID, v))
}

// MappingIDLT applies the LT predicate on the "mapping_id" field.
func MappingIDLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldMappingID, v))
}

// MappingIDLTE applies the LTE predicate on the "mapping_id" field.
func MappingIDLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldMappingID, v))
}

// MappingIDContains applies the Contains predicate on the "mapping_id" field.
func MappingIDContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldMappingID, v))
}

// MappingIDHasPrefix applies the HasPrefix predicate on the "mapping_id" field.
func MappingIDHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldMappingID, v))
}

// MappingIDHasSuffix applies the HasSuffix predicate on the "mapping_id" field.
func MappingIDHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldMappingID, v))
}

// MappingIDEqualFold applies the EqualFold predicate on the "mapping_id" field.
func MappingIDEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldMappingID, v))
}

// MappingIDContainsFold applies the ContainsFold predicate on the "mapping_id" field.
func MappingIDContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldMappingID, v))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// OrganizationIDGT applies the GT predicate on the "organization_id" field.
func OrganizationIDGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldOrganizationID, v))
}

// OrganizationIDGTE applies the GTE predicate on the "organization_id" field.
func OrganizationIDGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldOrganizationID, v))
}

// OrganizationIDLT applies the LT predicate on the "organization_id" field.
func OrganizationIDLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldOrganizationID, v))
}

// OrganizationIDLTE applies the LTE predicate on the "organization_id" field.
func OrganizationIDLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldOrganizationID, v))
}

// OrganizationIDContains applies the Contains predicate on the "organization_id" field.
func OrganizationIDContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldOrganizationID, v))
}

// OrganizationIDHasPrefix applies the HasPrefix predicate on the "organization_id" field.
func OrganizationIDHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldOrganizationID, v))
}

// OrganizationIDHasSuffix applies the HasSuffix predicate on the "organization_id" field.
func OrganizationIDHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldOrganizationID, v))
}

// OrganizationIDEqualFold applies the EqualFold predicate on the "organization_id" field.
func OrganizationIDEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldOrganizationID, v))
}

// OrganizationIDContainsFold applies the ContainsFold predicate on the "organization_id" field.
func OrganizationIDContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldOrganizationID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldName, v))
}

// GeoEQ applies the EQ predicate on the "geo" field.
func GeoEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldGeo, v))
}

// GeoNEQ applies the NEQ predicate on the "geo" field.
func GeoNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldGeo, v))
}

// GeoIn applies the In predicate on the "geo" field.
func GeoIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldGeo, vs...))
}

// GeoNotIn applies the NotIn predicate on the "geo" field.
func GeoNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldGeo, vs...))
}

// GeoGT applies the GT predicate on the "geo" field.
func GeoGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldGeo, v))
}

// GeoGTE applies the GTE predicate on the "geo" field.
func GeoGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldGeo, v))
}

// GeoLT applies the LT predicate on the "geo" field.
func GeoLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldGeo, v))
}

// GeoLTE applies the LTE predicate on the "geo" field.
func GeoLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldGeo, v))
}

// GeoContains applies the Contains predicate on the "geo" field.
func GeoContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldGeo, v))
}

// GeoHasPrefix applies the HasPrefix predicate on the "geo" field.
func GeoHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldGeo, v))
}

// GeoHasSuffix applies the HasSuffix predicate on the "geo" field.
func GeoHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldGeo, v))
}

// GeoIsNil applies the IsNil predicate on the "geo" field.
func GeoIsNil() predicate.Database {
	return predicate.Database(sql.FieldIsNull(FieldGeo))
}

// GeoNotNil applies the NotNil predicate on the "geo" field.
func GeoNotNil() predicate.Database {
	return predicate.Database(sql.FieldNotNull(FieldGeo))
}

// GeoEqualFold applies the EqualFold predicate on the "geo" field.
func GeoEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldGeo, v))
}

// GeoContainsFold applies the ContainsFold predicate on the "geo" field.
func GeoContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldGeo, v))
}

// DsnEQ applies the EQ predicate on the "dsn" field.
func DsnEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldDsn, v))
}

// DsnNEQ applies the NEQ predicate on the "dsn" field.
func DsnNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldDsn, v))
}

// DsnIn applies the In predicate on the "dsn" field.
func DsnIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldDsn, vs...))
}

// DsnNotIn applies the NotIn predicate on the "dsn" field.
func DsnNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldDsn, vs...))
}

// DsnGT applies the GT predicate on the "dsn" field.
func DsnGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldDsn, v))
}

// DsnGTE applies the GTE predicate on the "dsn" field.
func DsnGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldDsn, v))
}

// DsnLT applies the LT predicate on the "dsn" field.
func DsnLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldDsn, v))
}

// DsnLTE applies the LTE predicate on the "dsn" field.
func DsnLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldDsn, v))
}

// DsnContains applies the Contains predicate on the "dsn" field.
func DsnContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldDsn, v))
}

// DsnHasPrefix applies the HasPrefix predicate on the "dsn" field.
func DsnHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldDsn, v))
}

// DsnHasSuffix applies the HasSuffix predicate on the "dsn" field.
func DsnHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldDsn, v))
}

// DsnEqualFold applies the EqualFold predicate on the "dsn" field.
func DsnEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldDsn, v))
}

// DsnContainsFold applies the ContainsFold predicate on the "dsn" field.
func DsnContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldDsn, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldGroupID, v))
}

// GroupIDContains applies the Contains predicate on the "group_id" field.
func GroupIDContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldGroupID, v))
}

// GroupIDHasPrefix applies the HasPrefix predicate on the "group_id" field.
func GroupIDHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldGroupID, v))
}

// GroupIDHasSuffix applies the HasSuffix predicate on the "group_id" field.
func GroupIDHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldGroupID, v))
}

// GroupIDEqualFold applies the EqualFold predicate on the "group_id" field.
func GroupIDEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldGroupID, v))
}

// GroupIDContainsFold applies the ContainsFold predicate on the "group_id" field.
func GroupIDContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldGroupID, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldToken, v))
}

// TokenIsNil applies the IsNil predicate on the "token" field.
func TokenIsNil() predicate.Database {
	return predicate.Database(sql.FieldIsNull(FieldToken))
}

// TokenNotNil applies the NotNil predicate on the "token" field.
func TokenNotNil() predicate.Database {
	return predicate.Database(sql.FieldNotNull(FieldToken))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldToken, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.DatabaseStatus) predicate.Database {
	vc := v
	return predicate.Database(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.DatabaseStatus) predicate.Database {
	vc := v
	return predicate.Database(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.DatabaseStatus) predicate.Database {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Database(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.DatabaseStatus) predicate.Database {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Database(sql.FieldNotIn(FieldStatus, v...))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v enums.DatabaseProvider) predicate.Database {
	vc := v
	return predicate.Database(sql.FieldEQ(FieldProvider, vc))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v enums.DatabaseProvider) predicate.Database {
	vc := v
	return predicate.Database(sql.FieldNEQ(FieldProvider, vc))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...enums.DatabaseProvider) predicate.Database {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Database(sql.FieldIn(FieldProvider, v...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...enums.DatabaseProvider) predicate.Database {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Database(sql.FieldNotIn(FieldProvider, v...))
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.Database {
	return predicate.Database(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Group
		step.Edge.Schema = schemaConfig.Database
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.Database {
	return predicate.Database(func(s *sql.Selector) {
		step := newGroupStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Group
		step.Edge.Schema = schemaConfig.Database
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Database) predicate.Database {
	return predicate.Database(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Database) predicate.Database {
	return predicate.Database(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Database) predicate.Database {
	return predicate.Database(sql.NotPredicates(p))
}
