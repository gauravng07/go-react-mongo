package model


type Listing struct {
	Id 						string		`bson:"_id,omitempty"`
	UniqueId 				string 		`bson:"uniq_id,omitempty"`
	CrawlTimeStamp			string		`bson:"crawl_timestamp,omitempty"`
	ProductURL				string		`bson:"product_url,omitempty"`
	ProductName				string		`bson:"product_name,omitempty"`
	ProductCategoryTree		string		`bson:"product_category_tree,omitempty"`
	PID						string		`bson:"pid,omitempty"`
	RetailPrice				interface{}		`bson:"retail_price,omitempty"`
	DiscountedPrice			interface{}		`bson:"discounted_price,omitempty"`
	Image					string		`bson:"image,omitempty"`
	IsFKAdvanceProduct		string		`bson:"is_FK_Advantage_product,omitempty"`
	Description				string		`bson:"description,omitempty"`
	ProductRating			interface{}	`bson:"product_rating,omitempty"`
OverallRating				interface{}	`bson:"overall_rating,omitempty"`
	Brand					string		`bson:"brand,omitempty"`
	ProductSpecification 	string		`bson:"product_specifications,omitempty"`
}