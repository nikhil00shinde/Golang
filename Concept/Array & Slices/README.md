## Agenda

### Arrays
- Creation
- Built-in functions
- Working with arrays


### Slices
- Creation
- Built-in functions
- Working with slices


# Summary

## Arrays
- Collection of items with same type
- Fixed size
- Declaration styles
    - a := [3]int{1,2,3} // literal style
    - a := [...]int{1,2,3} //literal style
    - var a [3]int
- Access via zero-based index
   - a := [3]int {1,2,3} //a[1] ==2
- **len** function returns size of array
- Copies refer to different underlying data


## Slices
- Backed by array
- Creation styles
   - Slice existing array or slice
   - Literal style
   - Via make function
      - a := make([]int,10) // create slice with capacity and length == 10
      - a := make([]int,10,100) // create slice with capacity == 100 and length == 10
- **len** function returns length of slice 
- **cap** function returns length of underlying array
- **append** function to add elements to slice
   - May cause expensive copy operation if underlying array is too small
- Copies refer to same underlying array