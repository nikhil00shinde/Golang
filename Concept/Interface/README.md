# Agenda

## Basics
- type Writer interface{ Write([]byte) (int,error)}
- type ConsoleWriter struct {}
- func (cw ConsoleWriter) Write(data []byte) (int,error){
  n,err := fmt.Println(string(data))
  return n,err
}
## Composing Interface
- type Writer interface{ Write([]byte) (int,error)}
- type Closer interface{ Close([]byte) error}
- type WriterCloser interface {
  Writer
  Closer
} 

## Type Conversion
- var wc WriterCloser = NewBufferedWriterCloser()
- bwc := wc.(*BufferedWriterCloser)
   ## The empty Interface
      - var i interface{} = 0
   ## Type switches
      - switch i.(type){
        case int:
        fmt.Println()
        case string:
        fmt.Println()
        default:
        fmt.Println()
      }

## Implementing with values vs. pointer
- Method set of **value** is all methods with value receivers
- Method set of **pointer** is all methods, regardless of receiver type  


# Best Practices
- Use many, small interfaces
   - Single method interfaces are some of the most and flexible
      - io.Writer, io.Reader, interface{}

- Don't export interfaces for types that will be consumed
- Do export interfaces for types that will be used by package 
- Design functions and methods to recieve interfaces whenever possible