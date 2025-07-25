package controllers_test

import (
	"testing"

	"github.com/Danchitomoo/go_api_learning/controllers"
	"github.com/Danchitomoo/go_api_learning/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)
	m.Run()
}
