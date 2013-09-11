package rest

type ShopScore struct {
	DeliveryScore float32 `json:"delivery_score,string" bson:"delivery_score"`
	ItemScore     float32 `json:"item_score,string" bson:"item_score"`
	ServiceScore  float32 `json:"service_score,string" bson:"service_score"`
}

type Shop struct {
	AllCount    int        `json:"all_count" bson:"all_count"`
	Bulletin    string     `json:"bulletin" bson:"bulletin,omitempty"`
	Cid         int        `json:"cid" bson:"cid"`
	Created     string     `json:"created" bson:"created"`
	Desc        string     `json:"desc" bson:"desc"`
	Modified    string     `json:"modified" bson:"modified"`
	Nick        string     `json:"nick" bson:"nick"`
	PicPath     string     `json:"pic_path" bson:"pic_path"`
	RemainCount int        `json:"remain_count" bson:"remain_count"`
	ShopScore   *ShopScore `json:"shop_score" bson:"shop_score"`
	Sid         int        `json:"sid" bson:"sid"`
	Title       string     `json:"title" bson:"title"`
	UsedCount   int        `json:"used_count" bson:"used_count"`
}

type ShopGetResponse struct {
	Shop *Shop `json:"shop"`
}

type ShopGetResponseResult struct {
	Response *ShopGetResponse `json:"shop_get_response"`
}
