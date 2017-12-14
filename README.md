Apache Avro for Golang
=====================

This is fork of elodina/go-avro for Guazi corp use.

***Generate avro go files Usage***:  

`go get -u github.com/Guazi-inc/go-avro/gen_avro_go`  

`gen_avro_go --schema=`   

***Register the schema to Schema-Registry***:  

`EXPORT  SCHEMA_REGISTRY_ADDR=your schema registry address`  

If SCHEMA_REGISTRY_ADDR not specified, the generated avro go code always set schema_id = 0, just for local develop purpose.


