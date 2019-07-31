package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

type Myhandler struct{}
type home struct {
	Title string
}

const (
	VIEW = "./view/"
	UPLOADFILE   = "./upload/"
	SERVERPORT = ":9890"

)

const(
	GB int = 1 << (iota*10)
	MB int = 1 << (iota*10)
	KB int = 1 << (iota*10)
)
func main() {
	server := http.Server{
		Addr:        SERVERPORT,
		Handler:     &Myhandler{},
		ReadTimeout: 10 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = index
	mux["/upload"] = upload
	mux["/file"] = StaticServer

	fmt.Println("GB",GB)
	fmt.Println("MB",MB)
	fmt.Println("KB",KB)

	fmt.Println("Create Server Port:",server.Addr)

	err :=server.ListenAndServe()

	if err != nil {
		fmt.Println("Create Server Error:",err.Error())
	}


	//fmt.Println(time.Now().Unix())
	//
	//h := md5.New()
	//h.Write([]byte("JH2fa0a5178082f5f89b5488fb50312ceb"+"387e4f2ea0aed8cbc8b8f2c3dc1bfd27"+"13517515583"+"1"+"156456233113517515583")) // 需要加密的字符串为 password
	//submit :=hex.EncodeToString(h.Sum(nil))
	//
	//fmt.Println("submit:",submit)

}

func (*Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		fmt.Println("time:",time.Now().Format("2006-01-02 15:04")," url:",r.URL.String())
		h(w, r)
		return
	}
	if ok, _ := regexp.MatchString("/css/", r.URL.String()); ok {
		http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))).ServeHTTP(w, r)
	} else {
		http.StripPrefix("/", http.FileServer(http.Dir("./upload/"))).ServeHTTP(w, r)
	}

}

func upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles(VIEW + "file.html")
		t.Execute(w, "上传文件")
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Fprintf(w, "%v", "上传错误")
			return
		}
		fileext := filepath.Ext(handler.Filename)
		if check(fileext) == false {
			fmt.Fprintf(w, "%v", "不允许的上传类型")
			return
		}
		filename := strconv.FormatInt(time.Now().Unix(), 10) + fileext
		f, _ := os.OpenFile(UPLOADFILE+filename, os.O_CREATE|os.O_WRONLY, 0660)
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Fprintf(w, "%v", "上传失败")
			return
		}
		filedir, _ := filepath.Abs(UPLOADFILE + filename)
		fmt.Fprintf(w, "%v", filename+"上传完成,服务器地址:"+filedir)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	title := home{Title: "首页"}
	t, _ := template.ParseFiles(VIEW + "index.html")
	t.Execute(w, title)
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/file", http.FileServer(http.Dir("./upload/"))).ServeHTTP(w, r)
}

func check(name string) bool {
	ext := []string{".exe", ".js", ".sh"}

	for _, v := range ext {
		if v == name {
			return false
		}
	}
	return true
}
