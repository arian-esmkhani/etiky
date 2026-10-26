import { PreviewProductDto, ProductFilterDto } from "../types/data-type";


// Base API path for slider-related endpoints.
const SLIDER_API_BASE = "/api/slider";

/**
 * Generic fetcher function
 * 
 * @template T - The expected return type of the response.
 * @param endpoint - The specific endpoint path (e.g., "discount", "best-selling").
 * @returns A promise resolving to the parsed response data.
 * @throws Will throw an error if the network request fails or response is not OK.
*/
async function fetchSliderData<T>(endpoint: string): Promise<T> {
  const url = `${SLIDER_API_BASE}/${endpoint}`;
  
  const response = await fetch(url, {
    method: "GET",
  });

  if (!response.ok) {
    throw new Error(`Failed to fetch slider data from ${endpoint}: ${response.status} ${response.statusText}`);
  }

  return response.json() as Promise<T>;
}

/**
 * Service object containing all slider-related API calls.
 * Each method provides a clean, typed interface to fetch specific product lists.
*/
export const sliderService = {
  async GetProductsByDiscount(): Promise<PreviewProductDto[]> {
    return fetchSliderData<PreviewProductDto[]>("discount");
  },

  async GetBestSellingProducts(): Promise<PreviewProductDto[]> {
    return fetchSliderData<PreviewProductDto[]>("best-selling");
  },

  async GetCheapestProducts(): Promise<PreviewProductDto[]> {
    return fetchSliderData<PreviewProductDto[]>("cheapest");
  },

  async GetProductsByFilter(filter: ProductFilterDto): Promise<PreviewProductDto[]> {
    const url = `${SLIDER_API_BASE}/filter`;

    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(filter),
    });

    if (!response.ok) {
      throw new Error(
        `Failed to fetch filtered products: ${response.status} ${response.statusText}`
      );
    }

    return response.json() as Promise<PreviewProductDto[]>;
  },

};