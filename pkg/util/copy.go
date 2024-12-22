package util

import (
	"errors"
	"github.com/jinzhu/copier"
	"time"
)

func ConvertCopyOption(to, from any) copier.Option {
	return copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: int64(0),
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(int64)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return time.Unix(s, 0).Format(time.DateTime), nil
				},
			},
			{
				SrcType: int64(0),
				DstType: copier.Bool,
				Fn: func(src interface{}) (interface{}, error) {

					s, ok := src.(int64)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return s == 1, nil
				},
			},
		},
		FieldNameMapping: []copier.FieldNameMapping{
			{
				SrcType: from,
				DstType: to,
				Mapping: map[string]string{
					"CreatedAt": "CreatedTime",
					"UpdatedAt": "UpdatedTime",
				}},
		}}
}
