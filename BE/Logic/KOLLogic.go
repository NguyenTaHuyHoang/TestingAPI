package Logic

import (
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message

func GetKolLogic(data *DTO.GetSearchParam) ([]*DTO.KolDTO, error) {

	// Query the database and get the list of KOLs
	DB := Initializers.DB

	// check pagelimit
	data.Process()
	var resultListKols []*DTO.KolDTO

	if err := DB.Table(DTO.KolDTO{}.TableName()).Count(&data.PageSize).Error; err != nil {
		return nil, nil
	}
	if err := DB.Offset(int(data.PageIndex-1) * int(data.PageLimit)).Limit(int(data.PageLimit)).Find(&resultListKols).Error; err != nil {
		return nil, err
	}

	return resultListKols, nil
}
