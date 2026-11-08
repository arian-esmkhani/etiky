import { useCallback, useState } from "react";
import { finderService } from "../service/finder-service";
import { ApiError } from "../types/api-type";
import { ArticleDto, PreviewProductRes, ProductFilterDto } from "../types/data-type";

export const useFinder = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const getCheapest = useCallback(async (page?: number): Promise<PreviewProductRes | null> => {
        setLoading(true);
        setError(null);
        try {
            return await finderService.FindCheapestProducts(page);
        } catch (err) {
            const error = err as ApiError;
            let errorMessage = '';
            switch (error.type) {
                case 'VALIDATION_ERROR':
                    errorMessage = 'اطلاعات وارد شده معتبر نیست';
                    break;
                case 'RATE_LIMIT_EXCEEDED':
                    errorMessage = 'تعداد درخواست‌ های شما بیش از حد مجاز است، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                case 'INTERNAL_SERVER_ERROR':
                    errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                default:
                    errorMessage = 'خطا در دریافت محصولات ';
            }
            setError(errorMessage);
            throw new Error(errorMessage);
        } finally {
            setLoading(false);
        }
    }, []);

    const getBestSelling = useCallback(async (category?: string, page?: number): Promise<PreviewProductRes | null> => {
        setLoading(true);
        setError(null);
        try {
            return await finderService.FindBestSellingProducts(category, page);
        } catch (err) {
            const error = err as ApiError;
            let errorMessage = '';
            switch (error.type) {
                case 'VALIDATION_ERROR':
                    errorMessage = 'اطلاعات وارد شده معتبر نیست';
                    break;
                case 'RATE_LIMIT_EXCEEDED':
                    errorMessage = 'تعداد درخواست‌ های شما بیش از حد مجاز است، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                case 'INTERNAL_SERVER_ERROR':
                    errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                default:
                    errorMessage = 'خطا در دریافت محصولات پرفروش';
            }
            setError(errorMessage);
            throw new Error(errorMessage);
        } finally {
            setLoading(false);
        }
    }, []);

    const getByDiscount = useCallback(async (category?: string, page?: number): Promise<PreviewProductRes | null> => {
        setLoading(true);
        setError(null);
        try {
            return await finderService.FindProductsByDiscount(category, page);
        } catch (err) {
            const error = err as ApiError;
            let errorMessage = '';
            switch (error.type) {
                case 'VALIDATION_ERROR':
                    errorMessage = 'اطلاعات وارد شده معتبر نیست';
                    break;
                case 'RATE_LIMIT_EXCEEDED':
                    errorMessage = 'تعداد درخواست‌ های شما بیش از حد مجاز است، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                case 'INTERNAL_SERVER_ERROR':
                    errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                default:
                    errorMessage = 'خطا در دریافت محصولات تخفیف‌دار';
            }
            setError(errorMessage);
            throw new Error(errorMessage);
        } finally {
            setLoading(false);
        }
    }, []);

    const getCheapestProductsByFilter = useCallback(async (filter: ProductFilterDto, page?: number): Promise<PreviewProductRes | null> => {
        setLoading(true);
        setError(null);
        try {
            return await finderService.FindCheapestProductsByFilter(filter, page);
        } catch (err) {
            const error = err as ApiError;
            let errorMessage = '';
            switch (error.type) {
                case 'VALIDATION_ERROR':
                    errorMessage = 'اطلاعات فیلتر معتبر نیست';
                    break;
                case 'RATE_LIMIT_EXCEEDED':
                    errorMessage = 'تعداد درخواست‌های شما بیش از حد مجاز است، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                case 'INTERNAL_SERVER_ERROR':
                    errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                default:
                    errorMessage = 'خطا در دریافت محصولات پیشنهادی';
            }
            setError(errorMessage);
            throw new Error(errorMessage);
        } finally {
            setLoading(false);
        }
    }, []);

    const searchProducts = useCallback(async (
        category?: string,
        q?: string,
        lowPrice?: number,
        maxPrice?: number,
        page?: number
    ): Promise<PreviewProductRes | null> => {
        setLoading(true);
        setError(null);
        try {
            return await finderService.searchProducts(category, q, lowPrice, maxPrice, page);
        } catch (err) {
            const error = err as ApiError;
            let errorMessage = '';
            switch (error.type) {
                case 'VALIDATION_ERROR':
                    errorMessage = 'اطلاعات جستجو معتبر نیست';
                    break;
                case 'NOT_FOUND':
                    errorMessage = 'محصولی با این مشخصات یافت نشد';
                    break;
                case 'RATE_LIMIT_EXCEEDED':
                    errorMessage = 'تعداد درخواست‌ های شما بیش از حد مجاز است، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                case 'INTERNAL_SERVER_ERROR':
                    errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                default:
                    errorMessage = 'خطا در جستجوی محصولات';
            }
            setError(errorMessage);
            throw new Error(errorMessage);
        } finally {
            setLoading(false);
        }
    }, []);

    const searchSuggest = useCallback(async (q?: string): Promise<string[]> => {
        setLoading(true);
        setError(null);
        try {
            return await finderService.searchSuggest(q);
        } catch (err) {
            const error = err as ApiError;
            let errorMessage = '';
            switch (error.type) {
                case 'VALIDATION_ERROR':
                    errorMessage = 'اطلاعات جستجو معتبر نیست';
                    break;
                case 'RATE_LIMIT_EXCEEDED':
                    errorMessage = 'تعداد درخواست‌های شما بیش از حد مجاز است، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                case 'INTERNAL_SERVER_ERROR':
                    errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                default:
                    errorMessage = 'خطا در دریافت پیشنهادات جستجو';
            }
            setError(errorMessage);
            throw new Error(errorMessage);
        } finally {
            setLoading(false);
        }
    }, []);

    const findArticleByTitle = useCallback(async (title: string): Promise<ArticleDto | null> => {
        setLoading(true);
        setError(null);
        try {
            return await finderService.FindArticleByTitle(title);
        } catch (err) {
            const error = err as ApiError;
            let errorMessage = '';
            switch (error.type) {
                case 'VALIDATION_ERROR':
                    errorMessage = 'عنوان مقاله معتبر نیست';
                    break;
                case 'NOT_FOUND':
                    errorMessage = 'مقاله‌ای با این عنوان یافت نشد';
                    break;
                case 'INTERNAL_SERVER_ERROR':
                    errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                default:
                    errorMessage = 'خطا در دریافت مقاله';
            }
            setError(errorMessage);
            throw new Error(errorMessage);
        } finally {
            setLoading(false);
        }
    }, []);

    const findArticleBySlug = useCallback(async (slug: string): Promise<ArticleDto | null> => {
        setLoading(true);
        setError(null);
        try {
            return await finderService.FindArticleBySlug(slug);
        } catch (err) {
            const error = err as ApiError;
            let errorMessage = '';
            switch (error.type) {
                case 'VALIDATION_ERROR':
                    errorMessage = 'شناسه مقاله معتبر نیست';
                    break;
                case 'NOT_FOUND':
                    errorMessage = 'مقاله یافت نشد';
                    break;
                case 'INTERNAL_SERVER_ERROR':
                    errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                    break;
                default:
                    errorMessage = 'خطا در دریافت مقاله';
            }
            setError(errorMessage);
            throw new Error(errorMessage);
        } finally {
            setLoading(false);
        }
    }, []);

    return {
        loading,
        error,
        getCheapest,
        getBestSelling,
        getByDiscount,
        getCheapestProductsByFilter,
        searchProducts,
        searchSuggest,
        findArticleByTitle,
        findArticleBySlug,
    };
};