// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/geodetic/internal/ent/generated/group"
	"github.com/datumforge/geodetic/pkg/enums"
)

// Group is the model entity for the Group schema.
type Group struct {
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
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// the name of the group in turso
	Name string `json:"name,omitempty"`
	// the description of the group
	Description string `json:"description,omitempty"`
	// the primary of the group
	PrimaryLocation string `json:"primary_location,omitempty"`
	// the replica locations of the group
	Locations []string `json:"locations,omitempty"`
	// the auth token used to connect to the group
	Token string `json:"-"`
	// region the group
	Region enums.Region `json:"region,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges        GroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Databases holds the value of the databases edge.
	Databases []*Database `json:"databases,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedDatabases map[string][]*Database
}

// DatabasesOrErr returns the Databases value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) DatabasesOrErr() ([]*Database, error) {
	if e.loadedTypes[0] {
		return e.Databases, nil
	}
	return nil, &NotLoadedError{edge: "databases"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldLocations:
			values[i] = new([]byte)
		case group.FieldID, group.FieldCreatedBy, group.FieldUpdatedBy, group.FieldMappingID, group.FieldDeletedBy, group.FieldName, group.FieldDescription, group.FieldPrimaryLocation, group.FieldToken, group.FieldRegion:
			values[i] = new(sql.NullString)
		case group.FieldCreatedAt, group.FieldUpdatedAt, group.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gr.ID = value.String
			}
		case group.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case group.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gr.UpdatedAt = value.Time
			}
		case group.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gr.CreatedBy = value.String
			}
		case group.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gr.UpdatedBy = value.String
			}
		case group.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				gr.MappingID = value.String
			}
		case group.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gr.DeletedAt = value.Time
			}
		case group.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				gr.DeletedBy = value.String
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gr.Description = value.String
			}
		case group.FieldPrimaryLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_location", values[i])
			} else if value.Valid {
				gr.PrimaryLocation = value.String
			}
		case group.FieldLocations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field locations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gr.Locations); err != nil {
					return fmt.Errorf("unmarshal field locations: %w", err)
				}
			}
		case group.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				gr.Token = value.String
			}
		case group.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				gr.Region = enums.Region(value.String)
			}
		default:
			gr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Group.
// This includes values selected through modifiers, order, etc.
func (gr *Group) Value(name string) (ent.Value, error) {
	return gr.selectValues.Get(name)
}

// QueryDatabases queries the "databases" edge of the Group entity.
func (gr *Group) QueryDatabases() *DatabaseQuery {
	return NewGroupClient(gr.config).QueryDatabases(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return NewGroupClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("generated: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(gr.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(gr.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(gr.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(gr.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(gr.Description)
	builder.WriteString(", ")
	builder.WriteString("primary_location=")
	builder.WriteString(gr.PrimaryLocation)
	builder.WriteString(", ")
	builder.WriteString("locations=")
	builder.WriteString(fmt.Sprintf("%v", gr.Locations))
	builder.WriteString(", ")
	builder.WriteString("token=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(fmt.Sprintf("%v", gr.Region))
	builder.WriteByte(')')
	return builder.String()
}

// NamedDatabases returns the Databases named value or an error if the edge was not
// loaded in eager-loading with this name.
func (gr *Group) NamedDatabases(name string) ([]*Database, error) {
	if gr.Edges.namedDatabases == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := gr.Edges.namedDatabases[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (gr *Group) appendNamedDatabases(name string, edges ...*Database) {
	if gr.Edges.namedDatabases == nil {
		gr.Edges.namedDatabases = make(map[string][]*Database)
	}
	if len(edges) == 0 {
		gr.Edges.namedDatabases[name] = []*Database{}
	} else {
		gr.Edges.namedDatabases[name] = append(gr.Edges.namedDatabases[name], edges...)
	}
}

// Groups is a parsable slice of Group.
type Groups []*Group
