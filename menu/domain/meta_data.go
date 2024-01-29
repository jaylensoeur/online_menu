package domain

type MetaData struct {
	Page      int
	Sort      string
	PageTotal int
	Count     int
	Limit     int
}

func NewMetaData(page int, order string, pageTotal int, count int, limit int) MetaData {
	return MetaData{
		Page:      page,
		PageTotal: pageTotal,
		Limit:     limit,
		Sort:      order,
		Count:     count,
	}
}
