// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// EnabledRoleTable is the database name for the table.
const EnabledRoleTable = "information_schema.enabled_roles"

// EnabledRole represents a row from 'information_schema.enabled_roles'.
type EnabledRole struct {
	RoleName SQLIdentifier `json:"role_name"` // role_name
}

// Constants defining each column in the table.
const (
	EnabledRoleRoleNameField = "role_name"
)

// WhereClauses for every type in EnabledRole.
var (
	EnabledRoleRoleNameWhere SQLIdentifierField = "role_name"
)

// QueryOneEnabledRole retrieves a row from 'information_schema.enabled_roles' as a EnabledRole.
func QueryOneEnabledRole(db XODB, where WhereClause, order OrderBy) (*EnabledRole, error) {
	const origsqlstr = `SELECT ` +
		`role_name ` +
		`FROM information_schema.enabled_roles WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	er := &EnabledRole{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&er.RoleName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return er, nil
}

// QueryEnabledRole retrieves rows from 'information_schema.enabled_roles' as a slice of EnabledRole.
func QueryEnabledRole(db XODB, where WhereClause, order OrderBy) ([]*EnabledRole, error) {
	const origsqlstr = `SELECT ` +
		`role_name ` +
		`FROM information_schema.enabled_roles WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*EnabledRole
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		er := EnabledRole{}

		err = q.Scan(&er.RoleName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &er)
	}
	return vals, nil
}

// AllEnabledRole retrieves all rows from 'information_schema.enabled_roles' as a slice of EnabledRole.
func AllEnabledRole(db XODB, order OrderBy) ([]*EnabledRole, error) {
	const origsqlstr = `SELECT ` +
		`role_name ` +
		`FROM information_schema.enabled_roles`

	sqlstr := origsqlstr + order.String()

	var vals []*EnabledRole
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		er := EnabledRole{}

		err = q.Scan(&er.RoleName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &er)
	}
	return vals, nil
}
