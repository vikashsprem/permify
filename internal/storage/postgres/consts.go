package postgres

import "sort"

const (
	RelationTuplesTable   = "relation_tuples"
	AttributesTable       = "attributes"
	SchemaDefinitionTable = "schema_definitions"
	TransactionsTable     = "transactions"
	TenantsTable          = "tenants"
	BundlesTable          = "bundles"
)

// isSameArray - check if two arrays are the same
func isSameArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sortedA := make([]string, len(a))
	copy(sortedA, a)
	sort.Strings(sortedA)

	sortedB := make([]string, len(b))
	copy(sortedB, b)
	sort.Strings(sortedB)

	for i := range sortedA {
		if sortedA[i] != sortedB[i] {
			return false
		}
	}

	return true
}
