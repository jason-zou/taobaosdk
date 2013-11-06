package rest

type Feature struct {
    AttrKey string  `json:"attr_key" bson:"attr_key"`
    AttrValue string `json:"attr_value bson:"attr_value"`
}

type ItemCat struct {
    Features []*Feature     `json:"features" bson:"features"`
    Cid int     `json:"cid" bson:"cid"`
    ParentCid int   `json:"parent_cid" bson:"parent_cid"`
    Name string `json:"name" bson:"name"`
    IsParent bool `json:"is_parent" bson:"is_parent"`
    Status string `json:"status" bson:"status"` // normal, deleted
    SortOrder int  `json:"sort_border" bson:"sort_border"`
}

type ItemCats struct {
    ItemCatArray []*ItemCat `json:"item_cat"`
}

type ItemCatsGetResponse struct {
    LastModified string `json:"last_modified"`
    ItemCats *ItemCats `json:"item_cats"`
}

type ItemCatsGetResponseResult struct {
    Response *ItemCatsGetResponse `json:"itemcats_get_response"`
}
