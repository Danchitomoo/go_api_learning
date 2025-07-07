package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Danchitomoo/go_api_learning/models"
	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle

	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	log.Println(page)

	articleList := []models.Article{models.Article1, models.Article2}

	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}
	var article models.Article
	if articleID == 1 {
		article = models.Article1
	} else if articleID == 2 {
		article = models.Article2
	} else {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reqComment)
}
