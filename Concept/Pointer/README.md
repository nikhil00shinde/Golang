## Pointer 

### Creating Pointer
- Pointer types use an asterisk (*) as a prefix to type pointed to
    - **int* - a pointer to an integer
- Use the addressof operator (&) to get address of variable

### Dereferencing Pointer
- Dereference a pointer by preceding with an asterisk (*)
- Complex types (e.g. structs) are automatically dereferenced

### Creating pointers to objects
-  Can use the addressof operator (&) if value type already exists
   - ms := myStruct{foo:42}
   - p :=&ms

- Use addressof operator before initializer
   - &myStruct{foo:42}

- Use the **new** keyword
  - Can't initialize fields at the same time

### The new function
### Working with nil 
### Types with internal pointers
- All assignment operations in Go are copy operations
- Slices and maps contain internal pointers, so copies point to same underlying data