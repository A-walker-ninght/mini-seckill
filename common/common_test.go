package common

import (
	"fmt"
	"github.com/A-walker-ninght/mini-seckill/models"
	"testing"
)

func TestDataToStructByTagSql(t *testing.T) {
	data := map[string]string{
		"id":          "1",
		"productName": "produ",
		"productUrl":  "dsfag",
		"productNum":  "2",
	}
	product := &models.Product{}
	DataToStructByTagSql(data, product)
	fmt.Println(product)
}
