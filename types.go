package ctp

/*
 #include <string.h>
 #include <float.h>
 #include "ctp.h"
 int go2cStr(_GoString_ str, char* p, int l){
    size_t n = _GoStringLen(str);
    if (n == 0){
      return 1;
    }
    if (n+1 > l){
       return 2;
    }
    const char *src = _GoStringPtr(str);
    memcpy(p, src, n);
    p[n] = 0;
    return 0;
 }
 int cstrEnd(char* p, int n){
   char* temp = NULL;
   for(int i=0; i!=n; i++){
     temp = p + i;
     if (*temp == '\0'){
       return i;
     }
   }
   return n;
 }
 int invalid(double f){
   if( f == DBL_MAX){
     return 1;
   }
   return 0;
 }
*/
import "C"
import (
	"errors"
)

func c2goStr(p *C.char, n int) string {
	index := C.cstrEnd(p, C.int(n))
	return C.GoStringN(p, C.int(index))
}

func go2cStr(str string, p *C.char, l int) (err error) {
	ret := C.go2cStr(str, p, C.int(l))
	if ret == 1 {
		err = errors.New("str empty")
	} else if ret == 2 {
		err = errors.New("str too long")
	}
	return
}

func goFloat64(f C.double) float64 {
	if C.invalid(f) == 1 {
		return 0
	}
	return float64(f)
}

func go2cStrArray(strs []string) (ret []*C.char) {
	for _, v := range strs {
		ret = append(ret, C.CString(v))
	}
	return ret
}
