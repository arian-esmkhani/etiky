"use client";

import { useFinder } from "@/app/hooks/use-finder";
import { PreviewProductDto } from "@/app/types/data-type";
import { Card } from "@/app/ui/card";
import SingleChipFilter from "@/app/ui/single-chip-filter";
import PriceRangeSlider from "@/app/ui/price-range-slider";
import { useRouter, useSearchParams } from "next/navigation";
import { useCallback, useEffect, useState, useRef } from "react";
import { useToast } from "@/app/ui/toast";

const TAGS = ["ست", "کلیبس", "استیل", "کش مو", "هدیه", "رولکس"];

export default function SearchPage() {
    const router = useRouter();
    const searchParams = useSearchParams();
    const initialQuery = searchParams.get("q") || "";

    const { searchProducts, error: hookError } = useFinder();
    const toast = useToast();
    const [products, setProducts] = useState<PreviewProductDto[]>([]);
    const [page, setPage] = useState(0);
    const [hasNext, setHasNext] = useState(true);
    const [isInitialLoading, setIsInitialLoading] = useState(true);
    const [isLoadingMore, setIsLoadingMore] = useState(false);
    const [selectedCategory, setSelectedCategory] = useState<string | null>(null);
    const [priceRange, setPriceRange] = useState<{ min: number; max: number } | null>(null);

    const isLoadingRef = useRef(false);

    const fetchProducts = useCallback(async (
        pageNum: number,
        category: string | null,
        price: typeof priceRange
    ) => {
        if (isLoadingRef.current) return null;
        isLoadingRef.current = true;

        try {
            return await searchProducts(
                category || undefined,
                initialQuery || undefined,
                price?.min,
                price?.max,
                pageNum
            );
        } catch {
            toast.error(hookError || "مشکل در دریافت اطلاعات");
            console.log(hookError)
            setHasNext(false);
            return null;
        } finally {
            isLoadingRef.current = false;
        }
    }, [searchProducts, hookError, toast, initialQuery]);

    useEffect(() => {
        const loadInitial = async () => {
            setIsInitialLoading(true);
            setProducts([]);
            setPage(0);
            setHasNext(true);

            const data = await fetchProducts(0, selectedCategory, priceRange);
            if (data) {
                setProducts(data.products);
                setHasNext(data.has_next);
                if (data.products.length === 0) {
                    setHasNext(false);
                }
            }
            setIsInitialLoading(false);
        };

        loadInitial();
    }, [fetchProducts, selectedCategory, priceRange]);

    const loadMore = useCallback(async () => {
        if (isLoadingRef.current || !hasNext || isLoadingMore) return;
        
        const nextPage = page + 1;
        setIsLoadingMore(true);
        
        const data = await fetchProducts(nextPage, selectedCategory, priceRange);

        if (data && data.products.length > 0) {
            if (data.products.length === 0) {
                setHasNext(false);
            }

            const existingIds = new Set(products.map(p => p.id));
            const newProducts = data.products.filter(p => !existingIds.has(p.id));
            
            if (newProducts.length > 0) {
                setProducts((prev) => [...prev, ...newProducts]);
                setHasNext(data.has_next);
                setPage(nextPage);
            }
        }
        
        setIsLoadingMore(false);
    }, [page, hasNext, products, fetchProducts, selectedCategory, priceRange, isLoadingMore]);

    const handleCategorySelect = (category: string | null) => {
        setSelectedCategory(category);
    };

    const handlePriceChange = (values: { min: number; max: number }) => {
        setPriceRange(values);
    };

    useEffect(() => {
        if (!hasNext || isLoadingRef.current || isLoadingMore) return;
        
        let timeoutId: NodeJS.Timeout;
        
        const handleScroll = () => {
            clearTimeout(timeoutId);
            timeoutId = setTimeout(() => {
                const { scrollTop, scrollHeight, clientHeight } = document.documentElement;
                if (scrollTop + clientHeight >= scrollHeight - 300) {
                    loadMore();
                }
            }, 20);
        };

        window.addEventListener("scroll", handleScroll, { passive: true });
        return () => {
            window.removeEventListener("scroll", handleScroll);
            clearTimeout(timeoutId);
        };
    }, [loadMore, hasNext, isLoadingMore]);

    if (isInitialLoading) {
        return (
            <div className="max-w-8xl mx-[0.1vw] xl:mx-[5.5vw] 2xl:mx-30 lg:mt-2">
                <div className="lg:flex gap-6">
                    <div className="lg:flex-1">
                        <div className="mb-8">
                            <h1 className="text-3xl font-bold text-[rgb(217,47,74)] mb-2 text-right">
                                نتایج جستجو برای {initialQuery}
                            </h1>
                        </div>
                        <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5" dir="rtl">
                            {[...Array(10)].map((_, i) => (
                                <div key={i} className="min-h-[395px] w-70 p-5 text-center border border-neutral-200 bg-white rounded-lg">
                                    <div className="w-49 aspect-square rounded-lg bg-neutral-200 animate-pulse mx-auto" />
                                    <div className="h-3 bg-neutral-200 rounded w-28 mt-4 mx-auto animate-pulse" />
                                    <div className="flex items-center justify-center gap-2 mt-4">
                                        <div className="h-4 bg-neutral-200 rounded w-12 animate-pulse" />
                                        <div className="h-4 bg-neutral-200 rounded w-16 animate-pulse" />
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                    <div className="w-64 shrink-0 mt-25">
                        <div className="sticky top-52 space-y-4">
                            <PriceRangeSlider
                                min={0}
                                max={96000000}
                                step={100000}
                                onRangeChange={handlePriceChange}
                            />
                            <SingleChipFilter
                                title="فیلتر بر اساس تگ ها"
                                options={TAGS}
                                defaultSelected=""
                                onSelect={handleCategorySelect}
                            />
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    return (
        <div className="max-w-8xl mx-[0.1vw] xl:mx-[4.6vw] lg:mt-2">
            <div className="lg:flex gap-6">
                <div className="z-3 mb-1 lg:hidden backdrop-blur-lg
                shrink-0 sticky top-22 sm:top-28 md:top-39 flex flex-row">
                    <PriceRangeSlider
                        min={0}
                        max={96000000}
                        step={100000}
                        onRangeChange={handlePriceChange}
                    />
                    <SingleChipFilter
                        title="فیلتر بر اساس تگ ها"
                        options={TAGS}
                        defaultSelected=""
                        onSelect={handleCategorySelect}
                    />
                </div>
                <div className="lg:flex-1">
                    <div className="mb-8 hidden lg:block">
                        <h1 className="text-3xl font-bold text-[rgb(217,47,74)]
                        mb-2 text-right">
                            نتایج جستجو برای {initialQuery}
                        </h1>
                    </div>

                    {products.length === 0 ? (
                        <div className="text-center py-12">
                            <p className="text-gray-500 text-lg mb-4">
                                هیچ محصولی با این جستجو یافت نشد
                            </p>
                            <p className="text-gray-400 text-sm">
                                {selectedCategory
                                    ? `در دسته‌ی "${selectedCategory}"`
                                    : "سعی کنید عبارت دیگری جستجو کنید"}
                            </p>
                        </div>
                    ) : (
                        <>
                            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 
                            lg:grid-cols-4 2xl:grid-cols-5" dir="rtl">
                                {products.map((product) => (
                                    <Card
                                        key={product.id}
                                        imgSrc={"/api" + product.img_url}
                                        imgAlt={product.name}
                                        title={product.name}
                                        variant="magical"
                                        price={product.final_price}
                                        oldPrice={product.price}
                                        discountPercent={product.discount}
                                        size="found"
                                        onClick={() => router.push(`/components/${product.id}/product`)}
                                    />
                                ))}
                            </div>
                            {isLoadingMore && (
                                <div className="flex justify-center py-8">
                                    <div className="animate-spin rounded-full h-10 w-10 border-4 border-[rgb(217,47,74)] border-t-transparent"></div>
                                </div>
                            )}
                        </>
                    )}
                </div>

                <div className="w-64 shrink-0 mt-25 hidden lg:block ">
                    <div className="sticky top-52 space-y-4">
                        <PriceRangeSlider
                            min={0}
                            max={96000000}
                            step={100000}
                            onRangeChange={handlePriceChange}
                        />
                        <SingleChipFilter
                            title="فیلتر بر اساس تگ ها"
                            options={TAGS}
                            defaultSelected=""
                            onSelect={handleCategorySelect}
                        />
                    </div>
                </div>
            </div>
        </div>
    );
}