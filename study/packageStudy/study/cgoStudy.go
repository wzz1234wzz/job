package study

/*
#include <stdio.h>
#include <errno.h>
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
    const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
    const char* os = "darwin";
#elif defined(CGO_OS_LINUX)
    const char* os = "linux";
#else
#    error(unknown os)
#endif

void SayHello(const char* s);

static void SayHello0(const char* s) {
    puts(s);
}

static int div(int a, int b) {
    if(b == 0) {
        errno = EINVAL;
        return 0;
    }
    return a/b;
}

void printint(int v) {
    printf("printint: %d\n", v);
}

struct A {
    int i;
    float f;
    int type; // type 是 Go 语言的关键字
};
 */
import "C"
import "fmt"

func CgoMain(){
	println("hello cgo")
	C.puts(C.CString("Hello, World\n"))   // C.CString函数将Go语言字符串转为C语言字符串，最后调用CGO包的C.puts函数向标准输出窗口打印转换后的C字符串
	_, err := C.SayHello0(C.CString("Hello, ys\n")) // 使用自定义的函数,
	fmt.Println(err) //cgo对errno做了特殊处理，可以通过第二个返回值来获取C语言的错误状态。对于void类型函数，这个特性依然有效
	C.SayHello(C.CString("Hello, wzz\n")) // 外部

	v := 42
	C.printint(C.int(v)) //C.int(v)用于将一个Go中的int类型值强制类型转换转化为C语言中的int类型值，然后调用C语言定义的printint函数进行打印

	print(C.GoString(C.os),"\n")
	C.puts(C.os)

	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)
	fmt.Println(a._type) // _type 对应 type,如果结构体的成员名字中碰巧是Go语言的关键字，可以通过在成员名开头添加下划线来访问

	// CGO也针对<errno.h>标准库的errno宏做的特殊支持：在CGO调用C函数时如果有两个返回值，那么第二个返回值将对应errno错误状态。
	v0, err0 := C.div(2, 1)
	fmt.Println(v0, err0)

	v1, err1 := C.div(1, 0)
	fmt.Println(v1, err1)

	/*
	  既然SayHello函数已经放到独立的C文件中了，我们自然可以将对应的C文件编译打包为静态库或动态库文件供使用。
	  如果是以静态库或动态库方式引用SayHello函数的话，需要将对应的C源文件移出当前目录
	  （CGO构建程序会自动构建当前目录下的C源文件，从而导致C函数名冲突）
	 */

}
