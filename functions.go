package main
import "fmt"
import "encoding/json"

//#include <stdlib.h>
//#include <string.h>
//#include <stdio.h>
import "C"

type TestItem struct {
	Id	int
	Name	string
	Phone	string
}

var slicetestStrings []TestItem
var currentslice TestItem

//export SelectSliceAtIndex
func SelectSliceAtIndex(i int){ 
   currentslice = slicetestStrings[i]
   fmt.Println(currentslice)   
}


//export SelectIdAtIndex
func SelectIdAtIndex(i int) int{
   fmt.Println(slicetestStrings[i].Id)
   return slicetestStrings[i].Id 

}

//export PrintSlice
func PrintSlice(){
  fmt.Println(slicetestStrings)
  fmt.Println(currentslice)
}

//export getSliceAttribute
func getSliceAttribute(idptr *C.int, name *C.char, phone *C.char){
    *idptr = C.int(currentslice.Id)
    C.strcpy(name,C.CString(currentslice.Name))
    C.strcpy(phone,C.CString(currentslice.Phone))
}

//export AppendSlice
func AppendSlice(iptr C.int,str2 *C.char,str3 *C.char){
   tmp := TestItem{iptr,C.GoStringN(str2, 6),C.GoStringN(str3, 6)}
   slicetestStrings = append(slicetestStrings,tmp)
}


//export zeroptr
func zeroptr(iptr *C.int, dptr *C.double, fptr *C.float, str *C.char) {
    t := *iptr
    println(t)
    tmp := 13
    *iptr = C.int(tmp)

    u := *dptr
    println(u)
    tmp2 := 2.1
    *dptr = C.double(tmp2)

    i := *fptr
    println(i)
    tmp3 := 2.2
    *fptr = C.float(tmp3)
   
    tmp4 := C.GoStringN(str, 6)
    println(tmp4)

    tmp4 = tmp4[:3] + "z" + tmp4[3:]
   
    strtmp := C.CString(tmp4)
    C.strcpy(str,strtmp)
}

//export JsonExample
func JsonExample(iptr *C.int, dptr *C.double, fptr *C.float, str *C.char) {

    byt := []byte(`{"integer":69,"double":1.1,"float":1.2,"strs":["hello","world"]}`)
    var dat map[string]interface{}
	
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)


    ii := dat["integer"].(float64)
    fmt.Println(ii)
    *iptr = C.int(int(ii))


    d := dat["double"].(float64)
    fmt.Println(d)
    *dptr = C.double(d)

    f := dat["float"].(float64)
    fmt.Println(f)
    *fptr = C.float(f)

    strs := dat["strs"].([]interface{})
    str1 := strs[1].(string)
    fmt.Println(str1)
   
    strtmp := C.CString(str1)
    C.strcpy(str,strtmp)
}

func main() {}
