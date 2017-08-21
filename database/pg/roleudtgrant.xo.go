// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// RoleUdtGrantTable is the database name for the table.
const RoleUdtGrantTable = "information_schema.role_udt_grants"

// RoleUdtGrant represents a row from 'information_schema.role_udt_grants'.
type RoleUdtGrant struct {
	Grantor       SQLIdentifier `json:"grantor"`        // grantor
	Grantee       SQLIdentifier `json:"grantee"`        // grantee
	UdtCatalog    SQLIdentifier `json:"udt_catalog"`    // udt_catalog
	UdtSchema     SQLIdentifier `json:"udt_schema"`     // udt_schema
	UdtName       SQLIdentifier `json:"udt_name"`       // udt_name
	PrivilegeType CharacterData `json:"privilege_type"` // privilege_type
	IsGrantable   YesOrNo       `json:"is_grantable"`   // is_grantable
}

// Constants defining each column in the table.
const (
	RoleUdtGrantGrantorField       = "grantor"
	RoleUdtGrantGranteeField       = "grantee"
	RoleUdtGrantUdtCatalogField    = "udt_catalog"
	RoleUdtGrantUdtSchemaField     = "udt_schema"
	RoleUdtGrantUdtNameField       = "udt_name"
	RoleUdtGrantPrivilegeTypeField = "privilege_type"
	RoleUdtGrantIsGrantableField   = "is_grantable"
)

// WhereClauses for every type in RoleUdtGrant.
var (
	RoleUdtGrantGrantorWhere       SQLIdentifierField = "grantor"
	RoleUdtGrantGranteeWhere       SQLIdentifierField = "grantee"
	RoleUdtGrantUdtCatalogWhere    SQLIdentifierField = "udt_catalog"
	RoleUdtGrantUdtSchemaWhere     SQLIdentifierField = "udt_schema"
	RoleUdtGrantUdtNameWhere       SQLIdentifierField = "udt_name"
	RoleUdtGrantPrivilegeTypeWhere CharacterDataField = "privilege_type"
	RoleUdtGrantIsGrantableWhere   YesOrNoField       = "is_grantable"
)

// QueryOneRoleUdtGrant retrieves a row from 'information_schema.role_udt_grants' as a RoleUdtGrant.
func QueryOneRoleUdtGrant(db XODB, where WhereClause, order OrderBy) (*RoleUdtGrant, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, udt_catalog, udt_schema, udt_name, privilege_type, is_grantable ` +
		`FROM information_schema.role_udt_grants WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	rug := &RoleUdtGrant{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&rug.Grantor, &rug.Grantee, &rug.UdtCatalog, &rug.UdtSchema, &rug.UdtName, &rug.PrivilegeType, &rug.IsGrantable)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rug, nil
}

// QueryRoleUdtGrant retrieves rows from 'information_schema.role_udt_grants' as a slice of RoleUdtGrant.
func QueryRoleUdtGrant(db XODB, where WhereClause, order OrderBy) ([]*RoleUdtGrant, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, udt_catalog, udt_schema, udt_name, privilege_type, is_grantable ` +
		`FROM information_schema.role_udt_grants WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*RoleUdtGrant
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		rug := RoleUdtGrant{}

		err = q.Scan(&rug.Grantor, &rug.Grantee, &rug.UdtCatalog, &rug.UdtSchema, &rug.UdtName, &rug.PrivilegeType, &rug.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &rug)
	}
	return vals, nil
}

// AllRoleUdtGrant retrieves all rows from 'information_schema.role_udt_grants' as a slice of RoleUdtGrant.
func AllRoleUdtGrant(db XODB, order OrderBy) ([]*RoleUdtGrant, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, udt_catalog, udt_schema, udt_name, privilege_type, is_grantable ` +
		`FROM information_schema.role_udt_grants`

	sqlstr := origsqlstr + order.String()

	var vals []*RoleUdtGrant
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		rug := RoleUdtGrant{}

		err = q.Scan(&rug.Grantor, &rug.Grantee, &rug.UdtCatalog, &rug.UdtSchema, &rug.UdtName, &rug.PrivilegeType, &rug.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &rug)
	}
	return vals, nil
}
