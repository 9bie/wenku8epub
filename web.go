package main

import (
	"net/http"
)

// 一个十分暴力的小心数据筛选反代服务
var session =make(map[string]string)
func versionSession(c string)bool{
return true;
}
func login(w http.ResponseWriter, r *http.Request){
	c, err := r.Cookie("session")

		if err != nil{
			if versionSession(c.String()){
				http.Redirect(w,r,"/",302 );
			}
	}

}
func index(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(`<a herf="/topList?type=dayvisit">今日热榜</a>\n<a herf="/topList?type=monthvisit">本月热点</a>\n<a herf="/topList?type=goodnum">最受关注</a>\n<a herf="/topList?type=anime">已动画化</a>\n<a herf="/topList?type=postdate">最新入库</a>\n	`));
}
func topList(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	if _, ok := r.Form["type"]; ok {
		html := getList("https://www.wenku8.net/modules/article/toplist.php?sort="+r.Form["type"][0]);

	}else{
		http.Redirect(w,r,"/",302 );
	}
	}
func push(w http.ResponseWriter, r *http.Request){
	if _, ok := r.Form["id"]; ok {

	}
}