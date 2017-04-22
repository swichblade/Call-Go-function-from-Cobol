package main
import "fmt"
import "encoding/json"

//#include <stdlib.h>
//#include <string.h>
//#include <stdio.h>
import "C"


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
