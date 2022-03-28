# Agenda 

- Boolean type
- Numeric type
   - Integers
   - Floating point
   - Complex numbers
- Text types


## Boolean Type
  - Values are true or false
  - Not an alias for other types(e.g. int)
  - Zero value is false

## Numeric types
### Integer
  - int type has varying size, but min 32 bits
  - 8 bit (int8) through 64 bit(int64)

### Unsigned integers
  - 8 bit (byte and uint8) through 32 bit(uint32)

### Arthemetic operations
  - Addition, Substraction, multiplication, division, remainder

### Bitwise operations
  -  And,or, xor, and not

### Zero value is 0

### Can't mix types in same family(uint16 + uint32 = error)

### Floating point numbsers
  - Follow IEEE-754 standart
  - Zero value is 0
  - 32 and 64 bit versions
  - Literal styles
     - Decimal(3.14)
     - Exponential (13e18 or 2E10)
     - mixed(13.7e12)

### Complex Numbers
  - Zero value is (0+0i)
  - 64 and 128 bit versions
  - Built-in functions
     - complex - make complex number from two floats
     - real    - get real part as float
     - imag    - get iamginary part as float


### Text Types
  - Strings
    - UTF-8
    - Immutable
    - can be concatenated with plus (+) operator
    - can be converted to []byte
   
  - Rune
   - UTF-32
   - Alias for int32
   - Special methods normally required to process
      - e.g. string.Reader#ReadRune