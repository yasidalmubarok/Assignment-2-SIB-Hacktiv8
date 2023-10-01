package item_repository

import (
	"assignment-2/entity"
	"assignment-2/pkg/errs"
)

type Repository interface {
	GetItemsByCodes(ItemCodes []any) ([]entity.Item, errs.Error)
}
