// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/geodetic/internal/ent/generated/database"
)

// Database is the model entity for the Database schema.
type Database struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// the ID of the organization
	OrganizationID string `json:"organization_id,omitempty"`
	// the name to the database
	Name string `json:"name,omitempty"`
	// the geo location of the database
	Geo string `json:"geo,omitempty"`
	// the DSN to the database
	Dsn          string `json:"dsn,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Database) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case database.FieldID, database.FieldCreatedBy, database.FieldUpdatedBy, database.FieldDeletedBy, database.FieldOrganizationID, database.FieldName, database.FieldGeo, database.FieldDsn:
			values[i] = new(sql.NullString)
		case database.FieldCreatedAt, database.FieldUpdatedAt, database.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Database fields.
func (d *Database) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case database.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				d.ID = value.String
			}
		case database.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case database.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case database.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				d.CreatedBy = value.String
			}
		case database.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				d.UpdatedBy = value.String
			}
		case database.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				d.DeletedAt = value.Time
			}
		case database.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				d.DeletedBy = value.String
			}
		case database.FieldOrganizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				d.OrganizationID = value.String
			}
		case database.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case database.FieldGeo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field geo", values[i])
			} else if value.Valid {
				d.Geo = value.String
			}
		case database.FieldDsn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dsn", values[i])
			} else if value.Valid {
				d.Dsn = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Database.
// This includes values selected through modifiers, order, etc.
func (d *Database) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// Update returns a builder for updating this Database.
// Note that you need to call Database.Unwrap() before calling this method if this Database
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Database) Update() *DatabaseUpdateOne {
	return NewDatabaseClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Database entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Database) Unwrap() *Database {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("generated: Database is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Database) String() string {
	var builder strings.Builder
	builder.WriteString("Database(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(d.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(d.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(d.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(d.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(d.OrganizationID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("geo=")
	builder.WriteString(d.Geo)
	builder.WriteString(", ")
	builder.WriteString("dsn=")
	builder.WriteString(d.Dsn)
	builder.WriteByte(')')
	return builder.String()
}

// Databases is a parsable slice of Database.
type Databases []*Database
