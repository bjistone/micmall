package model

import (
	"database/sql/driver"
	"encoding/json"
)


type GoodsDetail struct {
	Goods int32 //商品id
	Num   int32 //扣减数量
}
type GoodsDetailList []GoodsDetail

func (g GoodsDetailList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

//Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GoodsDetailList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type Inventory struct {
	BaseModel
	Goods   int32 `gorm:"type:int;index"` //商品id
	Stocks  int32 `gorm:"type:int"`       //存货
	Version int32 `gorm:"type:int"`       //分布式锁的乐观锁
}

//StockSellDetail 记录分布式事务的订单状态
type StockSellDetail struct {
	OrderSn string          `gorm:"type:varchar(200);index:idx_order_sn,unique;"`
	Status  int32           `gorm:"type:varchar(200)"` //1 表示已扣减 2. 表示已归还
	Detail  GoodsDetailList `gorm:"type:varchar(200)"` //记录商品id和这个商品本次扣减的库存
}

func (StockSellDetail) TableName() string {
	return "stockselldetail"
}
