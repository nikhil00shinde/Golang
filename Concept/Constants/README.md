## Agenda

- Naming convention
- Typed constants
- Untyped constants
- Enumerated constants (method of generating constants)
- Enumeration Expression


### Immutable, but can be shadowed

### Replaced by the compilrer at compile time
   - Value must be calculate at compiler time

### Named like variables
   - PascalCase for exported constants
   - camelCase for internal constants

### Typed Constants work like immutable variables 
   - const a int = 12
   - Can interoperate only with same types

### Untyped constants work like literals  
  - const a = 12
  - Can interoperate with similar types


### Enumerated constants
  - Special symbol *iota* allows related constants to be created easily
  - Iota starts at 0 in each block and increments by one
  - Watch out of constant values that match zero values for variables

### Enumerated expressions
  - Operations that can be determined at compile time are allowed
       - Arithmetic
       - Bitwise operations
       - Bitshifting