
###一百行golng代码写一个静态文件服务器

####包含功能 

1. 静态文件模板
2. 文件上传
3. 文件查看和下载

####the use pkg

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

####Contains the knowledge

	//static server
	http.FileServer(http.Dir("目录"))

	//Manually configure the service and routing
	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe();

	// path/filepath The use of the package 

The sparrow is small, but there are a lot of basic knowledge for learning golang friends to play together

# Upload
#文件上传服务器

#内容逐渐完善中

