// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// TableConstraintTable is the database name for the table.
const TableConstraintTable = "information_schema.table_constraints"

// TableConstraint represents a row from 'information_schema.table_constraints'.
type TableConstraint struct {
	ConstraintCatalog SQLIdentifier `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  SQLIdentifier `json:"constraint_schema"`  // constraint_schema
	ConstraintName    SQLIdentifier `json:"constraint_name"`    // constraint_name
	TableCatalog      SQLIdentifier `json:"table_catalog"`      // table_catalog
	TableSchema       SQLIdentifier `json:"table_schema"`       // table_schema
	TableName         SQLIdentifier `json:"table_name"`         // table_name
	ConstraintType    CharacterData `json:"constraint_type"`    // constraint_type
	IsDeferrable      YesOrNo       `json:"is_deferrable"`      // is_deferrable
	InitiallyDeferred YesOrNo       `json:"initially_deferred"` // initially_deferred
}

// Constants defining each column in the table.
const (
	TableConstraintConstraintCatalogField = "constraint_catalog"
	TableConstraintConstraintSchemaField  = "constraint_schema"
	TableConstraintConstraintNameField    = "constraint_name"
	TableConstraintTableCatalogField      = "table_catalog"
	TableConstraintTableSchemaField       = "table_schema"
	TableConstraintTableNameField         = "table_name"
	TableConstraintConstraintTypeField    = "constraint_type"
	TableConstraintIsDeferrableField      = "is_deferrable"
	TableConstraintInitiallyDeferredField = "initially_deferred"
)

// WhereClauses for every type in TableConstraint.
var (
	TableConstraintConstraintCatalogWhere SQLIdentifierField = "constraint_catalog"
	TableConstraintConstraintSchemaWhere  SQLIdentifierField = "constraint_schema"
	TableConstraintConstraintNameWhere    SQLIdentifierField = "constraint_name"
	TableConstraintTableCatalogWhere      SQLIdentifierField = "table_catalog"
	TableConstraintTableSchemaWhere       SQLIdentifierField = "table_schema"
	TableConstraintTableNameWhere         SQLIdentifierField = "table_name"
	TableConstraintConstraintTypeWhere    CharacterDataField = "constraint_type"
	TableConstraintIsDeferrableWhere      YesOrNoField       = "is_deferrable"
	TableConstraintInitiallyDeferredWhere YesOrNoField       = "initially_deferred"
)

// QueryOneTableConstraint retrieves a row from 'information_schema.table_constraints' as a TableConstraint.
func QueryOneTableConstraint(db XODB, where WhereClause, order OrderBy) (*TableConstraint, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, table_catalog, table_schema, table_name, constraint_type, is_deferrable, initially_deferred ` +
		`FROM information_schema.table_constraints WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	tc := &TableConstraint{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&tc.ConstraintCatalog, &tc.ConstraintSchema, &tc.ConstraintName, &tc.TableCatalog, &tc.TableSchema, &tc.TableName, &tc.ConstraintType, &tc.IsDeferrable, &tc.InitiallyDeferred)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return tc, nil
}

// QueryTableConstraint retrieves rows from 'information_schema.table_constraints' as a slice of TableConstraint.
func QueryTableConstraint(db XODB, where WhereClause, order OrderBy) ([]*TableConstraint, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, table_catalog, table_schema, table_name, constraint_type, is_deferrable, initially_deferred ` +
		`FROM information_schema.table_constraints WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*TableConstraint
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		tc := TableConstraint{}

		err = q.Scan(&tc.ConstraintCatalog, &tc.ConstraintSchema, &tc.ConstraintName, &tc.TableCatalog, &tc.TableSchema, &tc.TableName, &tc.ConstraintType, &tc.IsDeferrable, &tc.InitiallyDeferred)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &tc)
	}
	return vals, nil
}

// AllTableConstraint retrieves all rows from 'information_schema.table_constraints' as a slice of TableConstraint.
func AllTableConstraint(db XODB, order OrderBy) ([]*TableConstraint, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, table_catalog, table_schema, table_name, constraint_type, is_deferrable, initially_deferred ` +
		`FROM information_schema.table_constraints`

	sqlstr := origsqlstr + order.String()

	var vals []*TableConstraint
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		tc := TableConstraint{}

		err = q.Scan(&tc.ConstraintCatalog, &tc.ConstraintSchema, &tc.ConstraintName, &tc.TableCatalog, &tc.TableSchema, &tc.TableName, &tc.ConstraintType, &tc.IsDeferrable, &tc.InitiallyDeferred)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &tc)
	}
	return vals, nil
}
