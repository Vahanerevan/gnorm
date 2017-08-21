// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// CharacterSetTable is the database name for the table.
const CharacterSetTable = "information_schema.character_sets"

// CharacterSet represents a row from 'information_schema.character_sets'.
type CharacterSet struct {
	CharacterSetCatalog   SQLIdentifier `json:"character_set_catalog"`   // character_set_catalog
	CharacterSetSchema    SQLIdentifier `json:"character_set_schema"`    // character_set_schema
	CharacterSetName      SQLIdentifier `json:"character_set_name"`      // character_set_name
	CharacterRepertoire   SQLIdentifier `json:"character_repertoire"`    // character_repertoire
	FormOfUse             SQLIdentifier `json:"form_of_use"`             // form_of_use
	DefaultCollateCatalog SQLIdentifier `json:"default_collate_catalog"` // default_collate_catalog
	DefaultCollateSchema  SQLIdentifier `json:"default_collate_schema"`  // default_collate_schema
	DefaultCollateName    SQLIdentifier `json:"default_collate_name"`    // default_collate_name
}

// Constants defining each column in the table.
const (
	CharacterSetCharacterSetCatalogField   = "character_set_catalog"
	CharacterSetCharacterSetSchemaField    = "character_set_schema"
	CharacterSetCharacterSetNameField      = "character_set_name"
	CharacterSetCharacterRepertoireField   = "character_repertoire"
	CharacterSetFormOfUseField             = "form_of_use"
	CharacterSetDefaultCollateCatalogField = "default_collate_catalog"
	CharacterSetDefaultCollateSchemaField  = "default_collate_schema"
	CharacterSetDefaultCollateNameField    = "default_collate_name"
)

// WhereClauses for every type in CharacterSet.
var (
	CharacterSetCharacterSetCatalogWhere   SQLIdentifierField = "character_set_catalog"
	CharacterSetCharacterSetSchemaWhere    SQLIdentifierField = "character_set_schema"
	CharacterSetCharacterSetNameWhere      SQLIdentifierField = "character_set_name"
	CharacterSetCharacterRepertoireWhere   SQLIdentifierField = "character_repertoire"
	CharacterSetFormOfUseWhere             SQLIdentifierField = "form_of_use"
	CharacterSetDefaultCollateCatalogWhere SQLIdentifierField = "default_collate_catalog"
	CharacterSetDefaultCollateSchemaWhere  SQLIdentifierField = "default_collate_schema"
	CharacterSetDefaultCollateNameWhere    SQLIdentifierField = "default_collate_name"
)

// QueryOneCharacterSet retrieves a row from 'information_schema.character_sets' as a CharacterSet.
func QueryOneCharacterSet(db XODB, where WhereClause, order OrderBy) (*CharacterSet, error) {
	const origsqlstr = `SELECT ` +
		`character_set_catalog, character_set_schema, character_set_name, character_repertoire, form_of_use, default_collate_catalog, default_collate_schema, default_collate_name ` +
		`FROM information_schema.character_sets WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	cs := &CharacterSet{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&cs.CharacterSetCatalog, &cs.CharacterSetSchema, &cs.CharacterSetName, &cs.CharacterRepertoire, &cs.FormOfUse, &cs.DefaultCollateCatalog, &cs.DefaultCollateSchema, &cs.DefaultCollateName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return cs, nil
}

// QueryCharacterSet retrieves rows from 'information_schema.character_sets' as a slice of CharacterSet.
func QueryCharacterSet(db XODB, where WhereClause, order OrderBy) ([]*CharacterSet, error) {
	const origsqlstr = `SELECT ` +
		`character_set_catalog, character_set_schema, character_set_name, character_repertoire, form_of_use, default_collate_catalog, default_collate_schema, default_collate_name ` +
		`FROM information_schema.character_sets WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*CharacterSet
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		cs := CharacterSet{}

		err = q.Scan(&cs.CharacterSetCatalog, &cs.CharacterSetSchema, &cs.CharacterSetName, &cs.CharacterRepertoire, &cs.FormOfUse, &cs.DefaultCollateCatalog, &cs.DefaultCollateSchema, &cs.DefaultCollateName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &cs)
	}
	return vals, nil
}

// AllCharacterSet retrieves all rows from 'information_schema.character_sets' as a slice of CharacterSet.
func AllCharacterSet(db XODB, order OrderBy) ([]*CharacterSet, error) {
	const origsqlstr = `SELECT ` +
		`character_set_catalog, character_set_schema, character_set_name, character_repertoire, form_of_use, default_collate_catalog, default_collate_schema, default_collate_name ` +
		`FROM information_schema.character_sets`

	sqlstr := origsqlstr + order.String()

	var vals []*CharacterSet
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		cs := CharacterSet{}

		err = q.Scan(&cs.CharacterSetCatalog, &cs.CharacterSetSchema, &cs.CharacterSetName, &cs.CharacterRepertoire, &cs.FormOfUse, &cs.DefaultCollateCatalog, &cs.DefaultCollateSchema, &cs.DefaultCollateName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &cs)
	}
	return vals, nil
}
