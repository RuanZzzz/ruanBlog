package controllers

import (
	"database/sql"
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"html/template"
	"net/http"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	// 1、获取URL参数
	id := route.GetRouteVariable("id", r)

	// 2、读取对应文章数据
	article, err := getArticleByID(id)

	// 3、如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 未找到数据
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章找不到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		//4. 读取成功
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).
			ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}

}
