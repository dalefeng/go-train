package orm

import (
	"github.com/stretchr/testify/assert"
	"go-train/orm/internal/errs"
	"testing"
)

func Test_parseModel(t *testing.T) {
	testCases := []struct {
		name   string
		entity any

		wantModel any
		wantErr   error
	}{
		{
			name:      "struct",
			entity:    TestModel{},
			wantModel: nil,
			wantErr:   errs.ErrPointerOnly,
		},
		{
			name:   "pointer",
			entity: &TestModel{},
			wantModel: &model{
				tableName: "test_model",
				fields: map[string]*field{
					"Id": {
						colName: "id",
					},
					"FirstName": {
						colName: "first_name",
					},
					"LastName": {
						colName: "last_name",
					},
					"Age": {
						colName: "age",
					},
				},
			},
		},
	}

	r := &registry{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := r.parseModel(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantModel, res)
		})
	}
}