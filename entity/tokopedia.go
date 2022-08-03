package entity

type TokopediaResponseParent struct {
	Data ParentData `json:"data"`
}

type ParentData struct {
	AceSearchProductV4 AceSearchProductV4 `json:"ace_search_product_v4"`
}

type AceSearchProductV4 struct {
	Data AceSearchData `json:"data"`
}

type AceSearchData struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID            int
	Name          string `json:"name"`
	Price         string `json:"price"`
	OriginalPrice string `json:"originalPrice"`
	Rating        int8   `json:"rating"`
	RatingAverage string `json:"ratingAverage"`
	Url           string `json:"url"`
	ImageUrl      string `json:"imageUrl"`
}
