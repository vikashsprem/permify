package factories

import (
	"github.com/Permify/permify/internal/repositories"
	MMRepository "github.com/Permify/permify/internal/repositories/memory"
	PQRepository "github.com/Permify/permify/internal/repositories/postgres"
	"github.com/Permify/permify/pkg/database"
	MMDatabase "github.com/Permify/permify/pkg/database/memory"
	PQDatabase "github.com/Permify/permify/pkg/database/postgres"
)

// RelationTupleFactory -
func RelationTupleFactory(db database.Database) (repo repositories.IRelationTupleRepository) {
	switch db.GetEngineType() {
	case "postgres":
		return PQRepository.NewRelationTupleRepository(db.(*PQDatabase.Postgres))
	case "memory":
		return MMRepository.NewRelationTupleRepository(db.(*MMDatabase.Memory))
	default:
		return MMRepository.NewRelationTupleRepository(db.(*MMDatabase.Memory))
	}
}

// EntityConfigFactory -
func EntityConfigFactory(db database.Database) (repo repositories.IEntityConfigRepository) {
	switch db.GetEngineType() {
	case "postgres":
		return PQRepository.NewEntityConfigRepository(db.(*PQDatabase.Postgres))
	case "memory":
		return MMRepository.NewEntityConfigRepository(db.(*MMDatabase.Memory))
	default:
		return PQRepository.NewEntityConfigRepository(db.(*PQDatabase.Postgres))
	}
}
