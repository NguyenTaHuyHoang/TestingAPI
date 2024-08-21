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
	DB := Initializers.DB

	data.Process()
	var resultListKols []*DTO.KolDTO

	query := DB.Table(DTO.KolDTO{}.TableName())
	if data.SearchParams != nil {
		for _, param := range *data.SearchParams {
			if param.Key != "" && param.Value != "" {
				query = query.Where(param.Key+" = ?", param.Value)
			}
		}
	}

	var totalRecords int64
	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, err
	}
	data.PageSize = totalRecords

	if err := query.Offset(int(data.PageIndex-1) * int(data.PageLimit)).Limit(int(data.PageLimit)).Find(&resultListKols).Error; err != nil {
		return nil, err
	}

	return resultListKols, nil
}
