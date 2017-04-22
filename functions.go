package main

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

func main() {}
