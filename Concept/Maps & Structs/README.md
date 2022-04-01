## Agenda

### Maps
- What are they?
- Creating
- Manipulation

### Structs
- What are they?
- Creating
- Naming Convention
- Embedding
- Tags


## Maps
- Collections of value types that are accessed via keys
- Created via literals or via make function
- Members accessed via [key] syntax
   - myMap["key"] = "value"
- Check for presence with "value,ok" form of result
- Multiple assignments refer to same underlying data 


## Structs
- Collection of disparate data types that describe single concept
- Keyed by named fields
- Normally created as types, but anonymus structs are allowed
- Structs are value types
- No inheritance, but can use composition via embedding
- Tags can be added to struct fields to describe field