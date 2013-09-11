package rest

type Location struct {
	City  string `json:"city" bson:"city"`
	State string `json:"state" bson:"state"`
}

type ItemImg struct {
	Id       int    `json:"id" bson:"id"`
	Url      string `json:"url" bson:"url"`
	Position int    `json:"position" bson:"position"`
	Created  string `json:"created" bson:"created"`
}

type ItemImgs struct {
	ItemImgArray []*ItemImg `json:"item_img" bson:"item_img"`
}

type Sku struct {
	PropertiesName string  `json:"properties_name" bson:"properties_name"`
	Properties     string  `json:"properties" bson:"properties"`
	SkuId          int     `json:"sku_id" bson:"sku_id"`
	Quantity       int     `json:"quantity" bson:"quantity"`
	Price          float32 `json:"price,string" bson:"price"`
	Created        string  `json:"created" bson:"created"`
	Modified       string  `json:"modified" bson:"modified"`
}

type Skus struct {
	SkuArray []*Sku `json:"sku" bson:"sku"`
}

type PropImg struct {
	Id         int    `json:"id" bson:"id"`
	Url        string `json:"url" bson:"url"`
	Properties string `json:"properties" bson:"properties"`
	Created    string `json:"created" bson:"created"`
}

type PropImgs struct {
	PropImgArray []*PropImg `json:"prop_img" bson:"prop_img"`
}

type Item struct {
	DetailUrl       string    `json:"detail_url" bson:"detail_url"`
	NumIid          int       `json:"num_iid" bson:"num_iid"`
	Title           string    `json:"title" bson:"title"`
	Nick            string    `json:"nick" bson:"nick"`
	Type            string    `json:"type" bson:"type"`
	Desc            string    `json:"desc" bson:"desc"`
	Cid             int       `json:"cid" bson:"cid"`
	PicUrl          string    `json:"pic_url" bson:"pic_url"`
	Num             int       `json:"num" bson:"num"`
	ListTime        string    `json:"list_time" bson:"list_time"`
	DelistTime      string    `json:"delist_time" bson:"delist_time"`
	StuffStatus     string    `json:"stuff_status" bson:"stuff_status"`
	Location        *Location `json:"location" bson:"location"`
	Price           float32   `json:"price,string" bson:"price"`
	GlobalStockType string    `json:"global_stock_type" bson:"global_stock_type"` //1现货, 2代购
	ItemImgs        *ItemImgs `json:"item_imgs" bson:"item_imgs"`
	Skus            *Skus     `json:"skus" bson:"skus"`
	PropsName       string    `json:"props_name" bson:"props_name"`
	PropImgs        *PropImgs `json:"prop_imgs" bson:"prop_imgs"`
}

type ItemGetResponse struct {
	Item *Item `json:"item"`
}

type ItemGetResponseResult struct {
	Response *ItemGetResponse `json:"item_get_response"`
}
