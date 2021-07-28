package entity

// OrderAssets represents...
//
// swagger:model
type Ingredient struct {
	Base

	Name string `gorm:"type:varchar(128); unique_index; not null" json:"name"`

	Img string `gorm:"type:varchar(256); not null" json:"img"`
}
