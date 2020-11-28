package main

/*
【http包执行流程】
1、创建Listen Socket监听指定的端口，等待客户端请求的到来
2、Listen Socket接受客户端的请求，得到Client Socket,接下来通过Client Socket与客户端通信
3、处理客户端的请求, 首先从Client Socket读取HTTP请求的协议头, 如果是POST方法, 还可能要读取客户端提
交的数据, 然后交给相应的handler处理请求, handler处理完毕准备好客户端需要的数据, 通过Client Socket写给客户端。
*/
import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello 浩远!") // 这个写入到 w 的是输出到客户端的
}
func main() {
	http.HandleFunc("/", sayhelloName) // 设置访问的路由
	/*
		初始化一个server对象，然后调用了 net.Listen("tcp", addr) ，也就是底层用TCP协议搭建了一个服务，然后监控我们设置的端口。
	*/
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
