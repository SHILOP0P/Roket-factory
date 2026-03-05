package part

import (
	"context"
	"inventory/internal/model"
	repoConverter "inventory/internal/repository/converter"
	repoModel "inventory/internal/repository/model"
)

func (s *repository) ListParts(_ context.Context, filter model.PartsFilter) ([]model.Part, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	f := repoConverter.PartsFilterToRepoModel(filter)
	parts := make([]*repoModel.Part, 0, len(s.parts))
	for _, p := range s.parts{
		if !matchByFilter(p, &f){
			continue
		}
		parts = append(parts, p)
	}
	
	out := make([]model.Part, 0, len(parts))
	for _, p := range parts{
		out = append(out, repoConverter.PartToModel(p))
	}

	return out, nil
}


func matchByFilter(p *repoModel.Part, f *repoModel.PartsFilter) bool {
	if f == nil {
		return true
	}
	if p == nil{
		return false
	}

	// uuids
	if len(f.Uuids) > 0 {
		ok := false
		for _, id := range f.Uuids {
			if p.Uuid == id {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	// names
	if len(f.Names) > 0 {
		ok := false
		for _, name := range f.Names {
			if p.Name == name {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	// categories
	if len(f.Categories) > 0 {
		ok := false
		for _, c := range f.Categories {
			if p.Category == c {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	// manufacturer_countries
	if len(f.ManufacturerCountries) > 0 {
		if p.Manufacturer == nil {
			return false
		}
		ok := false
		country := p.Manufacturer.Country
		for _, c := range f.ManufacturerCountries {
			if country == c {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	// tags (есть хотя бы один общий тег)
	if len(f.Tags) > 0 {
		ok := false
		for _, pt := range p.Tags {
			for _, ft := range f.Tags {
				if pt == ft {
					ok = true
					break
				}
			}
			if ok {
				break
			}
		}
		if !ok {
			return false
		}
	}

	return true
}