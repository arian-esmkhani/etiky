// Getting data for cafe preview (slider, cart, etc)
export interface PreviewProductDto{
    id:         number;
    name:       string;
	img_url:    string;
	price:      number;
	discount:   number;
	final_price: number;
}

// For Fetching products from db witch filter
export interface ProductFilterDto {
	mark:    string;
	feature: string;
	tag:     string;
}