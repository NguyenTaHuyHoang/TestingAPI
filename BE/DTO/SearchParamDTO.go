package DTO

type ListDTOModel[T any] struct {
	PageIndex   int64
	PageSize    int64
	PageLimit   int64
	TotalCount  int64
	TotalNotify int64
	CalValue1   float64
	CalValue2   float64
	CalValue3   float64
	Source      *[]T
}

type SearchParam struct {
	Key   string `json:"key" form:"key"`
	Value string `json:"value" form:"value"`
}

type AddedParam map[string]string

// pageNumber = pageIndex
// totalCount = pageSize
// pageLimit = totalCount

type GetSearchParam struct {
	PageIndex    int64          `json:"pageIndex" form:"pageIndex"`
	PageLimit    int64          `json:"pageLimit" form:"pageLimit"`
	PageSize     int64          `json:"pageSize" form:"-"`
	SearchParams *[]SearchParam `json:"searchParams" form:"searchParams"`
}

func (p *GetSearchParam) Process() {
	if p.PageIndex <= 0 {
		p.PageIndex = 1
	}
	if p.PageLimit <= 0 && p.PageLimit >= 100 {
		p.PageLimit = 10
	}
}
