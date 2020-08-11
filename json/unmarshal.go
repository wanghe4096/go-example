package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	bytes := []byte(`{
    "13-1963": {
        "id": 24047,
        "product_name": "暗时间",
        "product_id": 1963,
        "product_type": 13,
        "product_key": "13-1963",
        "cover_img": "img/201609/21/201609212105075040011632.jpg",
        "price": "4.99",
        "sku": "XLJD3X160921000001",
        "status": 3,
        "created_at": "2020-05-14T16:40:17+08:00",
        "updated_at": "2020-05-14T16:40:17+08:00"
    },
    "66-5": {
        "id": 25456,
        "product_name": "2刘嘉·心理学基础30讲",
        "product_id": 5,
        "product_type": 66,
        "product_key": "66-5",
        "cover_img": "img/201801/15/201801151513264434146226.jpg",
        "price": "1",
        "sku": "onlynana66666",
        "status": 1,
        "created_at": "2020-05-14T16:41:03+08:00",
        "updated_at": "2020-05-14T16:41:03+08:00"
    }
}`)

	result := make(map[string]Product)

	err := json.Unmarshal(bytes, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	price := 1
	s := fmt.Sprintf("%.2f", float64(price)/100.0)
	fmt.Println(s)

}


type Product struct {
	Id          int64     `json:"id" xorm:"pk autoincr comment('主键id') BIGINT(20)"`
	ProductName string    `json:"product_name" xorm:"not null default '' comment('商品名称') VARCHAR(255)"`
	ProductId   int64     `json:"product_id" xorm:"not null default 0 comment('商品ID') index BIGINT(20)"`
	ProductType int64     `json:"product_type" xorm:"not null default 0 comment('商品类型') index BIGINT(255)"`
	ProductKey  string    `json:"product_key" xorm:"not null default '' comment('索引字段') VARCHAR(128)"`
	CoverImg    string    `json:"cover_img" xorm:"not null default '' comment('封面') VARCHAR(255)"`
	Price       string    `json:"price" xorm:"not null default '' comment('价格') VARCHAR(10)"`
	Sku         string    `json:"sku" xorm:"not null default '' comment('商品编码') index VARCHAR(255)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('0-下架 1-上架') int"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') DATETIME"`
}
