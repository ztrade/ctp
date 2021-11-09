package ctp

/*
 #include <string.h>
 #include <float.h>
 #include <stdint.h>
 #include <stdlib.h>
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
 #include <stdio.h>
void dump_buf(char *buf, int bytes)
{
    int    i;
    for(i=0; i<bytes; i++)
    {
        printf("0x%02x ", buf[i]);
        if( 0 == (i+1)%16 )
            printf("\n");
    }
    printf("\n");
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
 void freeStr(char* s){
 free(s);
 }
 void freeStrArray(char **s, int l){
   for(int i=0; i!=l; i++){
     char **temp = s + i;
     free(*temp);
   }
   free(s);
 }
#cgo LDFLAGS: -L libs/linux64 -l:thostmduserapi_se.so -l:thosttraderapi_se.so
*/
import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"sync/atomic"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	ptrMap   sync.Map
	ptrIndex uint64
)

func GbkToUtf8(str string) (ret string, err error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewDecoder())
	decode, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	ret = string(decode)
	return
}

func getGoPtr(i uint64) interface{} {
	v, ok := ptrMap.Load(i)
	if ok {
		return v
	}
	return nil
}

func storeGoPtr(value interface{}) uint64 {
	n := atomic.AddUint64(&ptrIndex, 1)
	ptrMap.Store(n, value)
	return n
}

func c2goStr(p *C.char, n int) string {
	index := C.cstrEnd(p, C.int(n))
	str := C.GoStringN(p, C.int(index))
	ret, err := GbkToUtf8(str)
	if err != nil {
		fmt.Println("gbktoUtf8 error:", err.Error())
		return str
	}
	return ret
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

func go2cStrArray(strs []string) (ret **C.char) {
	cArray := C.malloc(C.size_t(len(strs)) * C.size_t(unsafe.Sizeof(uintptr(0))))
	a := (*[1<<30 - 1]*C.char)(cArray)
	for i, v := range strs {
		a[i] = C.CString(v)
	}
	return (**C.char)(cArray)
}

func go2cStrPtr(str string) *C.char {
	return C.CString(str)
}

func go2cBool(b bool) C.int8_t {
	if b {
		return 1
	}
	return 0
}

func cPtr2GoStr(p *C.char) string {
	return C.GoString(p)
}

func c2goBool(b C.int8_t) bool {
	if b == 0 {
		return false
	}
	return true
}

func freeCStr(str *C.char) {
	C.freeStr(str)
}

func freeCStrArray(str **C.char, n int) {
	C.freeStrArray(str, C.int(n))
}
