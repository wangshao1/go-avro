package avro

import (
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	// 	raw := `{"type": "record", "name": "TestRecord", "fields": [
	// 		{"name": "longRecordField", "type": "long"},
	// 		{"name": "stringRecordField", "type": "string"},
	// 		{"name": "intRecordField", "type": "int"},
	// 		{"name": "floatRecordField", "type": "float"}
	// ]}`
	// 	s, err := ParseSchema(raw)
	// assert(t, err, nil)
	now := time.Now()
	regClient := NewCachedSchemaRegistryClient("")

	// id, err := regClient.Register("this_is_a_test_schema", s)
	// assert(t, err, nil)
	// t.Logf("register id is:%v", id)

	id := int32(735)
	schema, err := regClient.GetByID(id)
	assert(t, err, nil)
	t.Logf("getByID scehma:%v", schema)

	for i := 0; i < 10; i++ {
		t.Logf("getByID:%v", time.Since(now))
		regClient.GetByID(id)
		assert(t, err, nil)
		//		t.Logf("getByID scehma:%v", schema)
		now = time.Now()
	}

	for i := 0; i < 10; i++ {
		idtemp, err := regClient.GetIDBySchema("this_is_a_test_schema", schema)
		assert(t, err, nil)
		t.Logf("getIDBySchema id:%v", idtemp)
		t.Logf("duration:%v", time.Since(now))
		now = time.Now()
	}

	idtemp, err := regClient.GetIDBySchema("this_is_a_test_schema", schema)
	assert(t, err, nil)
	t.Logf("twice getIDBySchema id:%v", idtemp)

}
