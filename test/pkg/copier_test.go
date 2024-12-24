package pkg

import (
	"errors"
	"github.com/jinzhu/copier"
	"testing"
	"time"
)

func TestCopierOption(t *testing.T) {
	type User1 struct {
		Status    int64
		CreatedAt int64
	}

	type User2 struct {
		Status      bool
		CreatedTime string
	}
	u := &User1{
		Status:    1,
		CreatedAt: 111111111,
	}
	m := new(User2)
	err := copier.CopyWithOption(m, u, convertOption(*m, *u))
	if err != nil {
		t.Error(err)
	}
	t.Log(m)
}

func convertOption(m any, u any) copier.Option {
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
				SrcType: u,
				DstType: m,
				Mapping: map[string]string{
					"CreatedAt": "CreatedTime",
				}},
		}}
}
