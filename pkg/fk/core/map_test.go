package core

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
)

func TestRandMap(testingT *testing.T) {

	randMap := NewRandMap(util.RandSeed())

	mockMap1 := map[any]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}
	mockMap2 := map[any][]any{
		1: {"value11", "value12"},
		2: {"value21", "value22"},
	}

	tFunc := gt.CreateTest(testingT)
	tFunc.Describe("GetRandomKeyValue", func() {
		tFunc.It("should return string key and string value", func() {
			kStr, vStr := GetRandomKeyValue(randMap, mockMap1)
			vExpected := mockMap1[kStr]
			gt.Expect(tFunc, &vStr).ToBe(vExpected)
		})

		tFunc.It("should return int key and slice value", func() {
			kInt, vSlice := GetRandomKeyValue(randMap, mockMap2)
			vExpected2 := mockMap2[kInt]
			gt.Expect(tFunc, &vSlice).ToBe(vExpected2)
		})
	})

	t := gt.CreateTest(testingT)
	t.Describe("RandMap", func() {
		t.Test("KeyValue should return key and the value", func() {
			k, v := randMap.KeyValue(mockMap1)

			vExpected := mockMap1[k]
			gt.Expect(t, &v).ToBe(vExpected)
		})
		t.Test("KeyValue should", func() {
			k, v := randMap.KeySliceValue(mockMap2)

			vExpected := mockMap2[k]
			gt.Expect(t, &v).ToBe(vExpected)
		})
	})

}
