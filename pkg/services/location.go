package services

import (
	"yuka-case/pkg/db"
	"yuka-case/pkg/entity"
	payloadSchema "yuka-case/pkg/schema/payload/v1"
)

type Location struct{}

func (l *Location) CreateLocation(payload *payloadSchema.Location) error {
	location := entity.Location{
		Name:        payload.Name,
		Latitude:    payload.Latitude,
		Longitude:   payload.Longitude,
		MarkerColor: payload.MarkerColor,
	}
	if err := db.Sql().Create(&location).Error; err != nil {
		panic(err)
	}

	return nil
}

func (l *Location) GetAllLocations() ([]entity.Location, error) {
	var locations []entity.Location
	if err := db.Sql().Find(&locations).Error; err != nil {
		panic(err)
	}

	return locations, nil
}

func (l *Location) GetLocations(page int) ([]entity.Location, error) {
	var locations []entity.Location
	if err := db.Sql().Offset((page - 1) * 20).Limit(20).Find(&locations).Error; err != nil {
		panic(err)
	}

	return locations, nil
}

func (l *Location) GetLocationById(id int) (entity.Location, error) {

	var location entity.Location
	if err := db.Sql().Where("id = ?", id).First(&location).Error; err != nil {
		panic(err)
	}

	return location, nil
}

func (l *Location) UpdateLocation(id int, payload *payloadSchema.Location) error {

	location := entity.Location{
		Name:        payload.Name,
		Latitude:    payload.Latitude,
		Longitude:   payload.Longitude,
		MarkerColor: payload.MarkerColor,
	}
	if err := db.Sql().Model(&location).Where("id = ?", id).Updates(&location).Error; err != nil {
		panic(err)
	}

	return nil
}

func (l *Location) DeleteLocation(id int) error {

	if err := db.Sql().Where("id = ?", id).Delete(&entity.Location{}).Error; err != nil {
		panic(err)
	}

	return nil
}
