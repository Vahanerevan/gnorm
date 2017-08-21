// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// RoutinePrivilegeTable is the database name for the table.
const RoutinePrivilegeTable = "information_schema.routine_privileges"

// RoutinePrivilege represents a row from 'information_schema.routine_privileges'.
type RoutinePrivilege struct {
	Grantor         SQLIdentifier `json:"grantor"`          // grantor
	Grantee         SQLIdentifier `json:"grantee"`          // grantee
	SpecificCatalog SQLIdentifier `json:"specific_catalog"` // specific_catalog
	SpecificSchema  SQLIdentifier `json:"specific_schema"`  // specific_schema
	SpecificName    SQLIdentifier `json:"specific_name"`    // specific_name
	RoutineCatalog  SQLIdentifier `json:"routine_catalog"`  // routine_catalog
	RoutineSchema   SQLIdentifier `json:"routine_schema"`   // routine_schema
	RoutineName     SQLIdentifier `json:"routine_name"`     // routine_name
	PrivilegeType   CharacterData `json:"privilege_type"`   // privilege_type
	IsGrantable     YesOrNo       `json:"is_grantable"`     // is_grantable
}

// Constants defining each column in the table.
const (
	RoutinePrivilegeGrantorField         = "grantor"
	RoutinePrivilegeGranteeField         = "grantee"
	RoutinePrivilegeSpecificCatalogField = "specific_catalog"
	RoutinePrivilegeSpecificSchemaField  = "specific_schema"
	RoutinePrivilegeSpecificNameField    = "specific_name"
	RoutinePrivilegeRoutineCatalogField  = "routine_catalog"
	RoutinePrivilegeRoutineSchemaField   = "routine_schema"
	RoutinePrivilegeRoutineNameField     = "routine_name"
	RoutinePrivilegePrivilegeTypeField   = "privilege_type"
	RoutinePrivilegeIsGrantableField     = "is_grantable"
)

// WhereClauses for every type in RoutinePrivilege.
var (
	RoutinePrivilegeGrantorWhere         SQLIdentifierField = "grantor"
	RoutinePrivilegeGranteeWhere         SQLIdentifierField = "grantee"
	RoutinePrivilegeSpecificCatalogWhere SQLIdentifierField = "specific_catalog"
	RoutinePrivilegeSpecificSchemaWhere  SQLIdentifierField = "specific_schema"
	RoutinePrivilegeSpecificNameWhere    SQLIdentifierField = "specific_name"
	RoutinePrivilegeRoutineCatalogWhere  SQLIdentifierField = "routine_catalog"
	RoutinePrivilegeRoutineSchemaWhere   SQLIdentifierField = "routine_schema"
	RoutinePrivilegeRoutineNameWhere     SQLIdentifierField = "routine_name"
	RoutinePrivilegePrivilegeTypeWhere   CharacterDataField = "privilege_type"
	RoutinePrivilegeIsGrantableWhere     YesOrNoField       = "is_grantable"
)

// QueryOneRoutinePrivilege retrieves a row from 'information_schema.routine_privileges' as a RoutinePrivilege.
func QueryOneRoutinePrivilege(db XODB, where WhereClause, order OrderBy) (*RoutinePrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, specific_catalog, specific_schema, specific_name, routine_catalog, routine_schema, routine_name, privilege_type, is_grantable ` +
		`FROM information_schema.routine_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	rp := &RoutinePrivilege{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&rp.Grantor, &rp.Grantee, &rp.SpecificCatalog, &rp.SpecificSchema, &rp.SpecificName, &rp.RoutineCatalog, &rp.RoutineSchema, &rp.RoutineName, &rp.PrivilegeType, &rp.IsGrantable)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rp, nil
}

// QueryRoutinePrivilege retrieves rows from 'information_schema.routine_privileges' as a slice of RoutinePrivilege.
func QueryRoutinePrivilege(db XODB, where WhereClause, order OrderBy) ([]*RoutinePrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, specific_catalog, specific_schema, specific_name, routine_catalog, routine_schema, routine_name, privilege_type, is_grantable ` +
		`FROM information_schema.routine_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*RoutinePrivilege
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		rp := RoutinePrivilege{}

		err = q.Scan(&rp.Grantor, &rp.Grantee, &rp.SpecificCatalog, &rp.SpecificSchema, &rp.SpecificName, &rp.RoutineCatalog, &rp.RoutineSchema, &rp.RoutineName, &rp.PrivilegeType, &rp.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &rp)
	}
	return vals, nil
}

// AllRoutinePrivilege retrieves all rows from 'information_schema.routine_privileges' as a slice of RoutinePrivilege.
func AllRoutinePrivilege(db XODB, order OrderBy) ([]*RoutinePrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, specific_catalog, specific_schema, specific_name, routine_catalog, routine_schema, routine_name, privilege_type, is_grantable ` +
		`FROM information_schema.routine_privileges`

	sqlstr := origsqlstr + order.String()

	var vals []*RoutinePrivilege
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		rp := RoutinePrivilege{}

		err = q.Scan(&rp.Grantor, &rp.Grantee, &rp.SpecificCatalog, &rp.SpecificSchema, &rp.SpecificName, &rp.RoutineCatalog, &rp.RoutineSchema, &rp.RoutineName, &rp.PrivilegeType, &rp.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &rp)
	}
	return vals, nil
}
