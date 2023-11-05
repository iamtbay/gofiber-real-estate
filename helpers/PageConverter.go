package helpers

import "math"

// page setter
func PageConverter(dataCount int64, productPerPage int64, page int) (int, int) {

	var flTotalPage float64 = float64(dataCount) / float64(productPerPage)
	totalPage := math.Ceil(flTotalPage)

	if page < 1 || page == 0 {
		page = 1
	} else if page > int(totalPage) {
		page = int(totalPage)
	}

	return int(totalPage), page

}
