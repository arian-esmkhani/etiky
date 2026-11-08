// Getting data for cafe preview (slider, cart, etc)
export interface PreviewProductDto{
    id:         number;
    name:       string;
	img_url:    string;
	price:      number;
	discount:   number;
	final_price: number;
}

// Getting data for cafe preview (slider, cart, etc) with pagination
export interface PreviewProductRes {
	products: PreviewProductDto[];
	has_next: boolean;
}


// For Fetching products from db witch filter
export interface ProductFilterDto {
	mark:    string;
	feature: string;
	tag:     string;
}

// Inner content structure of the article
export interface ContentDto {
    content_title: string;
    title_tag: string;
    paragraph: string[];
    paragraph_tag: string[];
    new_line: boolean[];
    content_img_url?: string; // Optional due to omitempty tag in Go
    order: number;
}

// Full article structure
export interface ArticleDto {
    id: number;
    title: string;
    description: string;
    views_count: number;
    author_id?: number; // Optional due to pointer and omitempty in Go
    status: string;
    img_url?: string;
    video_url?: string;
    tags?: string[];
    
    // SEO
    slug: string;
    meta_title: string;
    meta_description: string;
    meta_robots: string;
    canonical_mode: string;
    canonical_url: string;
    breadcrumb_title: string;
    breadcrumb_parent: string;
    breadcrumb_url: string;
    schema_type: string;

    // Date (as string since they come from JSON)
    published_at?: string; 
    updated_at?: string;

    // Content
    content: ContentDto[];
}

export interface LoginDto {
    number: string;
    password: string;
}

export interface AuthResponse {
    message: string;
}

export interface LoginResponse {
    success: boolean;
}