// +build ios

package tor

/*
#cgo CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

const char* GetTempDir() {
    NSString *tempDir = NSTemporaryDirectory();
    char *copy = strdup([tempDir UTF8String]);

    return copy;
}
*/
import "C"

import "unsafe"

func getTempDir() string {
    cstring := C.GetTempDir()
    tempDir := C.GoString(cstring)
    C.free(unsafe.Pointer(cstring))

    return tempDir
}
