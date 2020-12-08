package persist

import (
	"github.com/bingoohuang/go-rest-template/internal/pkg/db"
	"github.com/jinzhu/gorm"
)

// Create
func Create(value interface{}) error {
	return db.GetDB().Create(value).Error
}

// Save
func Save(value interface{}) error {
	return db.GetDB().Save(value).Error
}

// Updates
func Updates(where interface{}, value interface{}) error {
	return db.GetDB().Model(where).Updates(value).Error
}

// Delete
func DeleteByModel(model interface{}) (count int64, err error) {
	d := db.GetDB().Delete(model)
	err = d.Error
	if err != nil {
		return
	}
	count = d.RowsAffected
	return
}

// Delete
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	d := db.GetDB().Where(where).Delete(model)
	err = d.Error
	if err != nil {
		return
	}
	count = d.RowsAffected
	return
}

// Delete
func DeleteByID(model interface{}, id uint64) (count int64, err error) {
	d := db.GetDB().Where("id=?", id).Delete(model)
	err = d.Error
	if err != nil {
		return
	}
	count = d.RowsAffected
	return
}

// Delete
func DeleteByIDS(model interface{}, ids []uint64) (count int64, err error) {
	d := db.GetDB().Where("id in (?)", ids).Delete(model)
	err = d.Error
	if err != nil {
		return
	}
	count = d.RowsAffected
	return
}

// First
func FirstByID(out interface{}, id string) (notFound bool, err error) {
	err = db.GetDB().First(out, id).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// First
func First(where interface{}, out interface{}, associations []string) (notFound bool, err error) {
	d := db.GetDB()
	for _, a := range associations {
		d = d.Preload(a)
	}
	err = d.Where(where).First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// Find
func Find(where interface{}, out interface{}, associations []string, orders ...string) error {
	d := db.GetDB()
	for _, a := range associations {
		d = d.Preload(a)
	}
	d = d.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			d = d.Order(order)
		}
	}
	return d.Find(out).Error
}

// Scan
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = db.GetDB().Model(model).Where(where).Scan(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// ScanList
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	d := db.GetDB().Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			d = d.Order(order)
		}
	}
	return d.Scan(out).Error
}
