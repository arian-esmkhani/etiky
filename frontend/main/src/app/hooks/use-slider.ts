import { useCallback, useState } from "react";
import { sliderService } from "../service/slider-service";
import { ProductFilterDto } from "../types/data-type";

/**
 * Hook for managing slider product data.
 * Handles loading state, error state, and provides methods
 * to fetch different product lists via sliderService.
 */
export const useSlider = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null); 

    const getByDiscount = useCallback(async () => {
        setLoading(true);
        setError(null);
        try {
            const result = await sliderService.GetProductsByDiscount();
            return result;
        } catch (err) {
            setError(err instanceof Error ? err.message : 'An error occurred');
            return null;
        } finally {
            setLoading(false);
        }
    }, []);

    const getBestSelling = useCallback(async () => {
        setLoading(true);
        setError(null);
        try {
            const result = await sliderService.GetBestSellingProducts();
            return result;
        } catch (err) {
            setError(err instanceof Error ? err.message : 'An error occurred');
            return null;
        } finally {
            setLoading(false);
        }
    }, []);

    const getCheapest = useCallback(async () => {
        setLoading(true);
        setError(null);
        try {
            const result = await sliderService.GetCheapestProducts();
            return result;
        } catch (err) {
            setError(err instanceof Error ? err.message : 'An error occurred');
            return null;
        } finally {
            setLoading(false);
        }
    }, []);

    const getByFilter = useCallback(async (filter: ProductFilterDto) => {
        setLoading(true);
        setError(null);
        try {
            return await sliderService.GetProductsByFilter(filter);
        } catch (err) {
            setError(err instanceof Error ? err.message : "An error occurred");
            return null;
        } finally {
            setLoading(false);
        }
    }, []);

    return {
        loading,
        error,
        getByDiscount,
        getBestSelling,
        getCheapest,
        getByFilter,
    };
}