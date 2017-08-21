// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// UsagePrivilegeTable is the database name for the table.
const UsagePrivilegeTable = "information_schema.usage_privileges"

// UsagePrivilege represents a row from 'information_schema.usage_privileges'.
type UsagePrivilege struct {
	Grantor       SQLIdentifier `json:"grantor"`        // grantor
	Grantee       SQLIdentifier `json:"grantee"`        // grantee
	ObjectCatalog SQLIdentifier `json:"object_catalog"` // object_catalog
	ObjectSchema  SQLIdentifier `json:"object_schema"`  // object_schema
	ObjectName    SQLIdentifier `json:"object_name"`    // object_name
	ObjectType    CharacterData `json:"object_type"`    // object_type
	PrivilegeType CharacterData `json:"privilege_type"` // privilege_type
	IsGrantable   YesOrNo       `json:"is_grantable"`   // is_grantable
}

// Constants defining each column in the table.
const (
	UsagePrivilegeGrantorField       = "grantor"
	UsagePrivilegeGranteeField       = "grantee"
	UsagePrivilegeObjectCatalogField = "object_catalog"
	UsagePrivilegeObjectSchemaField  = "object_schema"
	UsagePrivilegeObjectNameField    = "object_name"
	UsagePrivilegeObjectTypeField    = "object_type"
	UsagePrivilegePrivilegeTypeField = "privilege_type"
	UsagePrivilegeIsGrantableField   = "is_grantable"
)

// WhereClauses for every type in UsagePrivilege.
var (
	UsagePrivilegeGrantorWhere       SQLIdentifierField = "grantor"
	UsagePrivilegeGranteeWhere       SQLIdentifierField = "grantee"
	UsagePrivilegeObjectCatalogWhere SQLIdentifierField = "object_catalog"
	UsagePrivilegeObjectSchemaWhere  SQLIdentifierField = "object_schema"
	UsagePrivilegeObjectNameWhere    SQLIdentifierField = "object_name"
	UsagePrivilegeObjectTypeWhere    CharacterDataField = "object_type"
	UsagePrivilegePrivilegeTypeWhere CharacterDataField = "privilege_type"
	UsagePrivilegeIsGrantableWhere   YesOrNoField       = "is_grantable"
)

// QueryOneUsagePrivilege retrieves a row from 'information_schema.usage_privileges' as a UsagePrivilege.
func QueryOneUsagePrivilege(db XODB, where WhereClause, order OrderBy) (*UsagePrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, object_catalog, object_schema, object_name, object_type, privilege_type, is_grantable ` +
		`FROM information_schema.usage_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	up := &UsagePrivilege{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&up.Grantor, &up.Grantee, &up.ObjectCatalog, &up.ObjectSchema, &up.ObjectName, &up.ObjectType, &up.PrivilegeType, &up.IsGrantable)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return up, nil
}

// QueryUsagePrivilege retrieves rows from 'information_schema.usage_privileges' as a slice of UsagePrivilege.
func QueryUsagePrivilege(db XODB, where WhereClause, order OrderBy) ([]*UsagePrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, object_catalog, object_schema, object_name, object_type, privilege_type, is_grantable ` +
		`FROM information_schema.usage_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*UsagePrivilege
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		up := UsagePrivilege{}

		err = q.Scan(&up.Grantor, &up.Grantee, &up.ObjectCatalog, &up.ObjectSchema, &up.ObjectName, &up.ObjectType, &up.PrivilegeType, &up.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &up)
	}
	return vals, nil
}

// AllUsagePrivilege retrieves all rows from 'information_schema.usage_privileges' as a slice of UsagePrivilege.
func AllUsagePrivilege(db XODB, order OrderBy) ([]*UsagePrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, object_catalog, object_schema, object_name, object_type, privilege_type, is_grantable ` +
		`FROM information_schema.usage_privileges`

	sqlstr := origsqlstr + order.String()

	var vals []*UsagePrivilege
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		up := UsagePrivilege{}

		err = q.Scan(&up.Grantor, &up.Grantee, &up.ObjectCatalog, &up.ObjectSchema, &up.ObjectName, &up.ObjectType, &up.PrivilegeType, &up.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &up)
	}
	return vals, nil
}
