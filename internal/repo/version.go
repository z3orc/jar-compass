package repo

import (
	"github.com/charmbracelet/log"
	"github.com/z3orc/compass/internal/data"
	"github.com/z3orc/compass/internal/model"
)

type VersionRepository struct {
	datasources map[model.Flavour]data.IDataSource
	flavours    []model.Flavour
}

func NewVersionRepository(datasources ...data.IDataSource) *VersionRepository {
	if len(datasources) <= 0 {
		return nil
	}

	r := VersionRepository{
		datasources: make(map[model.Flavour]data.IDataSource),
		flavours:    make([]model.Flavour, 0),
	}

	for _, ds := range datasources {
		r.flavours = append(r.flavours, ds.GetFlavour())
		r.datasources[ds.GetFlavour()] = ds

		log.Debug("Added datasource to VersionRepository", "flavour", ds.GetFlavour().ToString())
	}

	log.Info("Created VersionRepository", "datasource_count", len(r.datasources), "flavours", r.flavours)

	return &r
}

func (r *VersionRepository) GetVersion(flavour model.Flavour, id string) (*model.Version, data.DataError) {

	src, ok := r.datasources[flavour]
	if !ok {
		log.Error("Unknown flavour requested", "flavour_int", flavour, "flavour", flavour.ToString())
		return nil, &InvalidFlavourError{}
	}

	version, err := src.GetVersion(id)
	if err != nil {
		return nil, err
	}

	return version, err
}

func (r *VersionRepository) GetFlavours() []model.Flavour {
	return r.flavours
}