// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// TriggerTable is the database name for the table.
const TriggerTable = "information_schema.triggers"

// Trigger represents a row from 'information_schema.triggers'.
type Trigger struct {
	TriggerCatalog          SQLIdentifier  `json:"trigger_catalog"`            // trigger_catalog
	TriggerSchema           SQLIdentifier  `json:"trigger_schema"`             // trigger_schema
	TriggerName             SQLIdentifier  `json:"trigger_name"`               // trigger_name
	EventManipulation       CharacterData  `json:"event_manipulation"`         // event_manipulation
	EventObjectCatalog      SQLIdentifier  `json:"event_object_catalog"`       // event_object_catalog
	EventObjectSchema       SQLIdentifier  `json:"event_object_schema"`        // event_object_schema
	EventObjectTable        SQLIdentifier  `json:"event_object_table"`         // event_object_table
	ActionOrder             CardinalNumber `json:"action_order"`               // action_order
	ActionCondition         CharacterData  `json:"action_condition"`           // action_condition
	ActionStatement         CharacterData  `json:"action_statement"`           // action_statement
	ActionOrientation       CharacterData  `json:"action_orientation"`         // action_orientation
	ActionTiming            CharacterData  `json:"action_timing"`              // action_timing
	ActionReferenceOldTable SQLIdentifier  `json:"action_reference_old_table"` // action_reference_old_table
	ActionReferenceNewTable SQLIdentifier  `json:"action_reference_new_table"` // action_reference_new_table
	ActionReferenceOldRow   SQLIdentifier  `json:"action_reference_old_row"`   // action_reference_old_row
	ActionReferenceNewRow   SQLIdentifier  `json:"action_reference_new_row"`   // action_reference_new_row
	Created                 TimeStamp      `json:"created"`                    // created
}

// Constants defining each column in the table.
const (
	TriggerTriggerCatalogField          = "trigger_catalog"
	TriggerTriggerSchemaField           = "trigger_schema"
	TriggerTriggerNameField             = "trigger_name"
	TriggerEventManipulationField       = "event_manipulation"
	TriggerEventObjectCatalogField      = "event_object_catalog"
	TriggerEventObjectSchemaField       = "event_object_schema"
	TriggerEventObjectTableField        = "event_object_table"
	TriggerActionOrderField             = "action_order"
	TriggerActionConditionField         = "action_condition"
	TriggerActionStatementField         = "action_statement"
	TriggerActionOrientationField       = "action_orientation"
	TriggerActionTimingField            = "action_timing"
	TriggerActionReferenceOldTableField = "action_reference_old_table"
	TriggerActionReferenceNewTableField = "action_reference_new_table"
	TriggerActionReferenceOldRowField   = "action_reference_old_row"
	TriggerActionReferenceNewRowField   = "action_reference_new_row"
	TriggerCreatedField                 = "created"
)

// WhereClauses for every type in Trigger.
var (
	TriggerTriggerCatalogWhere          SQLIdentifierField  = "trigger_catalog"
	TriggerTriggerSchemaWhere           SQLIdentifierField  = "trigger_schema"
	TriggerTriggerNameWhere             SQLIdentifierField  = "trigger_name"
	TriggerEventManipulationWhere       CharacterDataField  = "event_manipulation"
	TriggerEventObjectCatalogWhere      SQLIdentifierField  = "event_object_catalog"
	TriggerEventObjectSchemaWhere       SQLIdentifierField  = "event_object_schema"
	TriggerEventObjectTableWhere        SQLIdentifierField  = "event_object_table"
	TriggerActionOrderWhere             CardinalNumberField = "action_order"
	TriggerActionConditionWhere         CharacterDataField  = "action_condition"
	TriggerActionStatementWhere         CharacterDataField  = "action_statement"
	TriggerActionOrientationWhere       CharacterDataField  = "action_orientation"
	TriggerActionTimingWhere            CharacterDataField  = "action_timing"
	TriggerActionReferenceOldTableWhere SQLIdentifierField  = "action_reference_old_table"
	TriggerActionReferenceNewTableWhere SQLIdentifierField  = "action_reference_new_table"
	TriggerActionReferenceOldRowWhere   SQLIdentifierField  = "action_reference_old_row"
	TriggerActionReferenceNewRowWhere   SQLIdentifierField  = "action_reference_new_row"
	TriggerCreatedWhere                 TimeStampField      = "created"
)

// QueryOneTrigger retrieves a row from 'information_schema.triggers' as a Trigger.
func QueryOneTrigger(db XODB, where WhereClause, order OrderBy) (*Trigger, error) {
	const origsqlstr = `SELECT ` +
		`trigger_catalog, trigger_schema, trigger_name, event_manipulation, event_object_catalog, event_object_schema, event_object_table, action_order, action_condition, action_statement, action_orientation, action_timing, action_reference_old_table, action_reference_new_table, action_reference_old_row, action_reference_new_row, created ` +
		`FROM information_schema.triggers WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	t := &Trigger{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&t.TriggerCatalog, &t.TriggerSchema, &t.TriggerName, &t.EventManipulation, &t.EventObjectCatalog, &t.EventObjectSchema, &t.EventObjectTable, &t.ActionOrder, &t.ActionCondition, &t.ActionStatement, &t.ActionOrientation, &t.ActionTiming, &t.ActionReferenceOldTable, &t.ActionReferenceNewTable, &t.ActionReferenceOldRow, &t.ActionReferenceNewRow, &t.Created)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return t, nil
}

// QueryTrigger retrieves rows from 'information_schema.triggers' as a slice of Trigger.
func QueryTrigger(db XODB, where WhereClause, order OrderBy) ([]*Trigger, error) {
	const origsqlstr = `SELECT ` +
		`trigger_catalog, trigger_schema, trigger_name, event_manipulation, event_object_catalog, event_object_schema, event_object_table, action_order, action_condition, action_statement, action_orientation, action_timing, action_reference_old_table, action_reference_new_table, action_reference_old_row, action_reference_new_row, created ` +
		`FROM information_schema.triggers WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*Trigger
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		t := Trigger{}

		err = q.Scan(&t.TriggerCatalog, &t.TriggerSchema, &t.TriggerName, &t.EventManipulation, &t.EventObjectCatalog, &t.EventObjectSchema, &t.EventObjectTable, &t.ActionOrder, &t.ActionCondition, &t.ActionStatement, &t.ActionOrientation, &t.ActionTiming, &t.ActionReferenceOldTable, &t.ActionReferenceNewTable, &t.ActionReferenceOldRow, &t.ActionReferenceNewRow, &t.Created)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &t)
	}
	return vals, nil
}

// AllTrigger retrieves all rows from 'information_schema.triggers' as a slice of Trigger.
func AllTrigger(db XODB, order OrderBy) ([]*Trigger, error) {
	const origsqlstr = `SELECT ` +
		`trigger_catalog, trigger_schema, trigger_name, event_manipulation, event_object_catalog, event_object_schema, event_object_table, action_order, action_condition, action_statement, action_orientation, action_timing, action_reference_old_table, action_reference_new_table, action_reference_old_row, action_reference_new_row, created ` +
		`FROM information_schema.triggers`

	sqlstr := origsqlstr + order.String()

	var vals []*Trigger
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		t := Trigger{}

		err = q.Scan(&t.TriggerCatalog, &t.TriggerSchema, &t.TriggerName, &t.EventManipulation, &t.EventObjectCatalog, &t.EventObjectSchema, &t.EventObjectTable, &t.ActionOrder, &t.ActionCondition, &t.ActionStatement, &t.ActionOrientation, &t.ActionTiming, &t.ActionReferenceOldTable, &t.ActionReferenceNewTable, &t.ActionReferenceOldRow, &t.ActionReferenceNewRow, &t.Created)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &t)
	}
	return vals, nil
}
