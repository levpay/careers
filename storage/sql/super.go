package sql

import (
	"github.com/dvdscripter/superheroapi/model"
	"github.com/google/uuid"
)

func (db *DB) findOrCreateGroup(group model.Group) (model.Group, error) {
	if err := db.Where("name = ?", group.Name).FirstOrCreate(&group).Error; err != nil {
		return group, err
	}

	return group, nil
}

func (db *DB) findOrCreateRelative(relative model.Relative) (model.Relative, error) {
	if err := db.Where("name = ?", relative.Name).FirstOrCreate(&relative).Error; err != nil {
		return relative, err
	}

	return relative, nil
}

func (db *DB) CreateSuper(super model.Super) (model.Super, error) {
	var relatives []*model.Relative
	var groups []*model.Group

	if super.Relatives != nil {
		relatives = make([]*model.Relative, len(super.Relatives))
		copy(relatives, super.Relatives)
		super.Relatives = nil
	}
	if super.Groups != nil {
		groups = make([]*model.Group, len(super.Groups))
		copy(groups, super.Groups)
		super.Groups = nil
	}

	if err := db.Create(&super).Error; err != nil {
		return super, err
	}

	// find/create relationships and append
	for i := range relatives {
		relative, err := db.findOrCreateRelative(*relatives[i])
		if err != nil {
			return super, err
		}
		if err := db.Model(&super).Association("Relatives").Append(&relative).Error; err != nil {
			return super, err
		}
	}
	for i := range groups {
		group, err := db.findOrCreateGroup(*groups[i])
		if err != nil {
			return super, err
		}
		if err := db.Model(&super).Association("Groups").Append(&group).Error; err != nil {
			return super, err
		}

	}

	return super, nil
}

func (db *DB) ListAllSuper() ([]model.Super, error) {
	var supers []model.Super
	if err := db.Preload("Groups").Preload("Relatives").Find(&supers).Error; err != nil {
		return nil, err
	}

	return supers, nil
}

func (db *DB) ListAllGood() ([]model.Super, error) {
	var supers []model.Super
	if err := db.Preload("Groups").Preload("Relatives").Where("alignment = ?", "good").Find(&supers).Error; err != nil {
		return nil, err
	}

	return supers, nil
}

func (db *DB) ListAllBad() ([]model.Super, error) {
	var supers []model.Super
	if err := db.Preload("Groups").Preload("Relatives").Where("alignment = ?", "bad").Find(&supers).Error; err != nil {
		return nil, err
	}

	return supers, nil
}

func (db *DB) FindByName(name string) (model.Super, error) {
	super := model.Super{}
	if err := db.Preload("Groups").Preload("Relatives").Where("name = ?", name).Limit(1).Find(&super).Error; err != nil {
		return super, err
	}

	return super, nil
}

func (db *DB) FindByID(id string) (model.Super, error) {
	super := model.Super{}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return super, err
	}
	if err := db.Preload("Groups").Preload("Relatives").Where("id = ?", uuid).Limit(1).Find(&super).Error; err != nil {
		return super, err
	}

	return super, nil
}

func (db *DB) DeleteByID(id string) error {
	super, err := db.FindByID(id)
	if err != nil {
		return err
	}

	if err := db.Model(&super).Association("Relatives").Delete(super.Relatives).Error; err != nil {
		return err
	}

	if err := db.Model(&super).Association("Groups").Delete(super.Groups).Error; err != nil {
		return err
	}

	return db.Delete(&super).Error
}
