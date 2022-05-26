# Schema Builder

Extract metadata from Golang Objects.

(OpenAPI 3.0.3 Schema)[https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schema]


### Options

Property Name -> json-tag | struct-field-name 
Property Name Filter -> camel-case | snake-case | none
Respect Omitempty -> true | false
Tag Prefix -> default: qd
Follow Pointers -> true | false

### Input Types

Int, Int8, Int16, Int32, Int64
UInt, UInt8, UInt16, UInt32, UInt64
Float32, Float64
Bool
String
Slice
Map
Struct


### Target Types

Integer
Number
Boolean
String
Struct
