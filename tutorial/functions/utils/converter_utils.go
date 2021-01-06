package utils

import (
	"fmt"
	"strconv"
)

func ConvertPaginationParams(pageStr, sizeStr string) (page int, size int, err error) {
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return page, size, fmt.Errorf("Fail to convert page :%s", err.Error())
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		return
	}
	return
}

