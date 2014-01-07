package rest

import "time"

type ShopScore struct {
	//	DeliveryScore float32   `json:"delivery_score,string" bson:"delivery_score"`
	//	ItemScore     float32   `json:"item_score,string" bson:"item_score"`
	//	ServiceScore  float32   `json:"service_score,string" bson:"service_score"`
	SemiScore    *ShopKPS     `bson:"semi_score"`                     //店铺半年内评分详情
	ServiceScore *ShopService `bson:"service_score"`                  //店铺30天内服务信息数据
	Credit       string       `json:"credit" bson:"credit"`           //店铺等级
	PraiseRate   float32      `json:"praise_rate" bson:"praise_rate"` //好评率
}

//店铺半年内动态评分
type ShopKPS struct {
	DescScore     *RateScore `bson:"desc_score"`
	ServiceScore  *RateScore `bson:"service_score"`
	DeliveryScore *RateScore `bson:"delivery_score"`
}

type RateScore struct {
	Score float32 `bson:"score"`
	Rate  float32 `bson:"rate"` //与同行业水平的差值比率
}

//店铺30天内服务信息数据,可参考http://rate.taobao.com/user-rate-1588199389.htm
type ShopService struct {
	AvgRefund       *Refund     `bson:"ave_refund" json:"avgRefund"`
	PunishCount     *Punish     `bson:"punish_count" json:"punish"`
	Raterefund      *RateRefund `bson:"rate_refund" json:"ratRefund"`
	ComplaintsCount *Complaints `bson:"complaints" json:"complaints"`
}

//平均退款速度
type Refund struct {
	IndVal   float32 `bson:"ind_val" json:"indVal"`     //行业均值
	LocalVal float32 `bson:"local_val" json:"localVal"` //本店值
}

//处罚次数
type Punish struct {
	IndVal         float32 `bson:"ind_val" json:"indVal"`
	LocalVal       float32 `bson:"local_val" json:"localVal"`
	ProhibitedInfo float32 `bson:"prohibited_info" json:"prohibitedInfoTimes"` //发布违禁信息被处罚次数
	Infringement   float32 `bson:"infringement" json:"infringement"`           //侵犯知识版权
	FalseMerch     float32 `bson:"false_merch" json:"falseMerchTimes"`         //虚假交易被处罚
}

//平均退款率
type RateRefund struct {
	Receive     int     `bson:"receive" json:"merchReceiveTimes"` //未收到货退款次数
	RefundCount int     `bson:"refund_count" json:"refundCount"`  //退款总次数
	IndVal      float32 `bson:"ind_val" json:"indVal"`            //行业均值退款率
	LocalVal    float32 `bson:"local_val" json:"localVal"`        //本店均值退款率
	NoReason    int     `bson:"no_reason" json:"noReasonTimes"`   //买家无理由退款次数
	Quality     int     `bson:"quality" json:"merchQualityTimes"` //商品质量问题退款次数
}

//投诉情况
type Complaints struct {
	IndVal    float32 `bson:"ind_val" json:"indVal"`
	LocalVal  float32 `bson:"local_val" json:"localVal"`
	AfterSale int     `bson:"after_sale" json:"afterSaleTimes"` //售后问题投诉次数
	Violation int     `bson:"violation" json:"violationTimes"`  //违规被投诉次数
	ComCount  int     `bson:"com_count" json:"complaintsCount"` //被投诉总次数
}
type Shop struct {
	AllCount     int        `json:"all_count" bson:"all_count"`         //已失效
	Bulletin     string     `json:"bulletin" bson:"bulletin,omitempty"` //已失效
	Cid          int        `json:"cid" bson:"cid"`
	Created      string     `json:"created" bson:"created"`
	Desc         string     `json:"desc" bson:"desc"`
	Modified     string     `json:"modified" bson:"modified"`
	Nick         string     `json:"nick" bson:"nick"`
	PicPath      string     `json:"pic_path" bson:"pic_path"`
	RemainCount  int        `json:"remain_count" bson:"remain_count"`
	ShopScore    *ShopScore `json:"shop_score" bson:"shop_score"`
	Sid          int        `json:"sid" bson:"sid"`
	Sellerid     int        `json:"sellerid" bson:"seller_id"`
	Title        string     `json:"title" bson:"title"`
	UsedCount    int        `json:"used_count" bson:"used_count"`
	MainProducts string     `json:"main_products" bson:"main_products"` //主营产品
	Company      string     `json:"company" bson:"company"`             //店铺所属公司
	Location     string     `json:"location" bson:"location"`           //所在地
	ShopType     string     `json:"shoptype" bson:"shop_type"`
	Synced       bool       `json:"synced" bson:"synced"` //是否已经同步到线下了
	ShopLink     string     `json:"shoplink" bson:"shop_link"`
	UpdatedTime  time.Time  `json:"updated_time",bson:"updated_time"`
}

type ShopGetResponse struct {
	Shop *Shop `json:"shop"`
}

type ShopGetResponseResult struct {
	Response *ShopGetResponse `json:"shop_get_response"`
}
