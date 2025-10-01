package utils

import (
	"fmt"
	"strings"
)

func AddLikeClause(field string) string {
	return fmt.Sprintf("%s LIKE ?", field)
}

func AddLikeValue(value string) string {
	return "%" + value + "%"
}

func MergeConditionAND(args ...string) string {
	return strings.Join(args, " AND ")
}
