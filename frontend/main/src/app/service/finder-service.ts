import { ArticleDto, PreviewProductRes, ProductFilterDto } from "../types/data-type";

export const finderService = {
    async FindCheapestProducts(page?: number): Promise<PreviewProductRes> {
        const url = new URL(`/api/finder/cheapest`, window.location.origin);
        if (page !== undefined) url.searchParams.append('page', String(page));
        const response = await fetch(url.toString(), { method: 'GET' });

        checkFinderResponse(response);

        return response.json();
    },

    async FindBestSellingProducts(category?: string, page?: number): Promise<PreviewProductRes> {
        const url = new URL(`/api/finder/best-selling`, window.location.origin);
        if (category) url.searchParams.append('category', category);
        if (page !== undefined) url.searchParams.append('page', String(page));
        const response = await fetch(url.toString(), { method: 'GET' });

        checkFinderResponse(response);

        return response.json();
    },

    async FindProductsByDiscount(category?: string, page?: number): Promise<PreviewProductRes> {
        const url = new URL(`/api/finder/discount`, window.location.origin);
        if (category) url.searchParams.append('category', category);
        if (page !== undefined) url.searchParams.append('page', String(page));
        const response = await fetch(url.toString(), { method: 'GET' });

        checkFinderResponse(response);

        return response.json();
    },

    async FindCheapestProductsByFilter(filter: ProductFilterDto,
        page?: number): Promise<PreviewProductRes> {
        const url = new URL(`/api/finder/suggested`, window.location.origin);
        if (page !== undefined) url.searchParams.append('page', String(page));
        const response = await fetch(url.toString(), {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(filter),
        });

        checkFinderResponse(response);

        return response.json();
    },

    async searchProducts(
        category?: string,
        q?: string,
        lowPrice?: number,
        maxPrice?: number,
        page?: number
    ): Promise<PreviewProductRes> {
        const url = new URL(`/api/finder/search`, window.location.origin);
        if (q) url.searchParams.append('q', q);
        if (category) url.searchParams.append('category', category);
        if (lowPrice !== undefined) url.searchParams.append('low_price', String(lowPrice));
        if (maxPrice !== undefined) url.searchParams.append('max_price', String(maxPrice));
        if (page !== undefined) url.searchParams.append('page', String(page));
        const response = await fetch(url.toString(), {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
        });

        if (!response.ok) {
            switch (response.status) {
                case 400:
                    throw {
                        type: 'VALIDATION_ERROR'
                    };
                case 404:
                    throw {
                        type: 'NOT_FOUND'
                    };
                case 429:
                    throw {
                        type: 'RATE_LIMIT_EXCEEDED'
                    };
                case 500:
                    throw {
                        type: 'SERVER_ERROR'
                    };
                default:
                    throw {
                        type: 'UNKNOWN_ERROR'
                    };
            }
        }

        return response.json();
    },

    async searchSuggest(
        q?: string,
    ): Promise<string[]> {
        const url = new URL(`/api/finder/search/suggestions`, window.location.origin);
        if (q) url.searchParams.append('q', q);
        const response = await fetch(url.toString(), {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });

        if (!response.ok) {
            switch (response.status) {
                case 400:
                    throw {
                        type: 'VALIDATION_ERROR'
                    };
                case 429:
                    throw {
                        type: 'RATE_LIMIT_EXCEEDED'
                    };
                case 500:
                    throw {
                        type: 'SERVER_ERROR'
                    };
                default:
                    throw {
                        type: 'UNKNOWN_ERROR'
                    };
            }
        }

        return response.json();
    },

    // Articles
    async FindArticleByTitle(title: string): Promise<ArticleDto> {
        const url = new URL(`/api/finder/articles/${encodeURIComponent(title)}`, window.location.origin);
        
        const response = await fetch(url.toString(), { method: 'GET' });


        if (!response.ok) {
            switch (response.status) {
                case 400:
                    throw {
                        type: 'VALIDATION_ERROR'
                    };
                case 404:
                    throw {
                        type: 'NOT_FOUND'
                    };
                case 429:
                    throw {
                        type: 'RATE_LIMIT_EXCEEDED'
                    };
                case 500:
                    throw {
                        type: 'SERVER_ERROR'
                    };
                default:
                    throw {
                        type: 'UNKNOWN_ERROR'
                    };
            }
        }

        return response.json();
    },

    async FindArticleBySlug(slug: string): Promise<ArticleDto> {
        const url = new URL(`/api/finder/articles/${encodeURIComponent(slug)}`, window.location.origin);
        
        const response = await fetch(url.toString(), { method: 'GET' });

        if (!response.ok) {
            switch (response.status) {
                case 400:
                    throw {
                        type: 'VALIDATION_ERROR'
                    };
                case 404:
                    throw {
                        type: 'NOT_FOUND'
                    };
                case 429:
                    throw {
                        type: 'RATE_LIMIT_EXCEEDED'
                    };
                case 500:
                    throw {
                        type: 'SERVER_ERROR'
                    };
                default:
                    throw {
                        type: 'UNKNOWN_ERROR'
                    };
            }
        }

        return response.json();
    },
};

function checkFinderResponse(response: Response): void {
    if (response.status === 204) {
        throw {
            type: 'NO_CONTENT'
        }
    }

    if (!response.ok) {
        switch (response.status) {
            case 400:
                throw {
                    type: 'VALIDATION_ERROR'
                };
            case 429:
                throw {
                    type: 'RATE_LIMIT_EXCEEDED'
                };
            case 500:
                throw {
                    type: 'INTERNAL_SERVER_ERROR'
                };
            default:
                throw {
                    type: 'UNKNOWN_ERROR'
                };
        }
    }
}