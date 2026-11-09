'use client'

import { useFinder } from "@/app/hooks/use-finder";
import { PreviewProductDto } from "@/app/types/data-type";
import { Card } from "@/app/ui/card";
import SingleChipFilter from "@/app/ui/single-chip-filter";
import { useToast } from "@/app/ui/toast";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState, useRef } from "react";

export default function BestSellingPage() {
    const { getBestSelling, error: hookError } = useFinder();
    const toast = useToast();
    const [products, setProducts] = useState<PreviewProductDto[]>([]);
    const [page, setPage] = useState(0);
    const [hasNext, setHasNext] = useState(true);
    const [isInitialLoading, setIsInitialLoading] = useState(true);
    const [isLoadingMore, setIsLoadingMore] = useState(false);
    const [selectedCategory, setSelectedCategory] = useState<string | null>(null);
    const isLoadingRef = useRef(false);
    const router = useRouter();

    const categories = ["دست‌بند", "ست", "ساعت", "زیورالات مو", "گردن‌بند", "گوشواره",
        "آویز", "انگشتر", "هدیه", "زنانه", "جاکلیدی"
    ];

    const fetchProducts = useCallback(async (pageNum: number, category: string | null) => {
        if (isLoadingRef.current) return null;
        
        isLoadingRef.current = true;
        
        try {
            const data = await getBestSelling(category || undefined, pageNum);
            return data;
        } catch {
            toast.error(hookError || "مشکل در دریافت اطلاعات");
            return null;
        } finally {
            isLoadingRef.current = false;
        }
    }, [getBestSelling, hookError, toast]);

    useEffect(() => {
        const loadInitialData = async () => {
            setIsInitialLoading(true);
            setProducts([]);
            setPage(0);
            setHasNext(true);
            
            const data = await fetchProducts(0, selectedCategory);
            if (data) {
                setProducts(data.products);
                setHasNext(data.has_next);
            }
            setIsInitialLoading(false);
        };
        
        loadInitialData();
    }, [fetchProducts, selectedCategory]);

    const loadMore = useCallback(async () => {
        if (isLoadingRef.current || !hasNext || isLoadingMore) return;
        
        const nextPage = page + 1;
        setIsLoadingMore(true);
        
        const data = await fetchProducts(nextPage, selectedCategory);
        if (data && data.products.length > 0) {
            const existingIds = new Set(products.map(p => p.id));
            const newProducts = data.products.filter(p => !existingIds.has(p.id));
            
            if (newProducts.length > 0) {
                setProducts(prev => [...prev, ...newProducts]);
                setHasNext(data.has_next);
                setPage(nextPage);
            }
        }
        
        setIsLoadingMore(false);
    }, [page, hasNext, products, fetchProducts, selectedCategory, isLoadingMore]);

    const handleCategorySelect = (category: string | null) => {
        setSelectedCategory(category);
    };

    // Infinite scroll
    useEffect(() => {
        if (!hasNext || isLoadingRef.current || isLoadingMore) return;
        
        let timeoutId: NodeJS.Timeout;
        
        const handleScroll = () => {
            clearTimeout(timeoutId);
            timeoutId = setTimeout(() => {
                const { scrollTop, scrollHeight, clientHeight } = document.documentElement;
                const isNearBottom = scrollTop + clientHeight >= scrollHeight - 300;
                
                if (isNearBottom) {
                    loadMore();
                }
            }, 20);
        };

        window.addEventListener('scroll', handleScroll, { passive: true });
        return () => {
            window.removeEventListener('scroll', handleScroll);
            clearTimeout(timeoutId);
        };
    }, [loadMore, hasNext, isLoadingMore]);

    if (isInitialLoading) {
        return (
            <div className="max-w-8xl mx-30 mt-2">
                <div className="flex">
                    <div className="flex-1">
                        <div className="hidden lg:block">
                            <h1 className="text-3xl font-bold text-[rgb(217,47,74)] mb-2 text-right">
                                پر فروش ترین ها
                            </h1>
                            <p className="text-gray-600 text-right">
                                محبوب ترین کالا ها در ماه های اخیر
                            </p>
                        </div>
                        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 2xl:grid-cols-5" dir="rtl">
                            {[...Array(10)].map((_, i) => (
                                <div key={i} className="min-h-[395px] lg:w-[20.7vw] xl:w-[18.27vw] 2xl:w-[15vw] p-5 text-center border border-neutral-200 bg-white rounded-lg flex flex-col items-center">
                                    <div className="w-49 aspect-square rounded-lg bg-neutral-100 animate-pulse" />
                                    <div className="w-28 h-3 bg-neutral-100 rounded mt-4 animate-pulse" />
                                    <div className="flex items-center justify-center gap-2 mt-4">
                                        <div className="w-12 h-4 bg-neutral-100 rounded animate-pulse" />
                                        <div className="w-16 h-4 bg-neutral-100 rounded animate-pulse" />
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    return (
        <div className="max-w-8xl mx-[0.1vw] xl:mx-[5.5vw] 2xl:mx-30 lg:mt-2">
            <div className="lg:flex">
                <div className="sticky top-22 sm:top-28 md:top-39 z-3 mb-1
                lg:hidden backdrop-blur-lg">
                    <SingleChipFilter
                        defaultSelected={selectedCategory || ""}
                        title="فیلتر بر اساس دسته بندی ها"
                        options={categories}
                        onSelect={handleCategorySelect}
                    />
                </div>
                <div className="lg:flex-1">
                    <div className="mb-8 hidden lg:block">
                            <h1 className="text-3xl font-bold text-[rgb(217,47,74)] 
                            mb-2 text-right">
                                پر فروش ترین ها
                            </h1>
                            <p className="text-gray-600 text-right">
                                محبوب ترین کالا ها در ماه های اخیر
                            </p>
                    </div>
                    {products.length === 0 ? (
                        <div className="text-center py-12">
                            <p className="text-gray-500 dark:text-gray-400 text-lg mb-4">
                                محصولی یافت نشد
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
                                        onClick={() => router.push(`/components/${product.id.toString()}/product`)}
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
                <div className="w-54 shrink-0 mt-25 ml-2 hidden lg:block">
                    <div className="sticky top-52">
                        <SingleChipFilter
                            defaultSelected={selectedCategory || ""}
                            title="فیلتر بر اساس دسته بندی ها"
                            options={categories}
                            onSelect={handleCategorySelect}
                        />
                    </div>
                </div>
            </div>
        </div>
    );
}