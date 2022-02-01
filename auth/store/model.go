package store

type OrderModel struct {
	order string
	sort  string
}

type PageQueryModel struct {
	Order  []OrderModel
	Offset int64
	Count  int64
}

type PageDataModel struct {
	Data  []interface{} `json:"data"`
	Count int64         `json:"count"`
}

type UserModel struct {
	ID       string `bson:"_id"`
	Account  string `bson:"account"`
	Password string `bson:"password"`
	Phone    string `bson:"phone"`
}
