'use client'

import { useFinder } from "@/app/hooks/use-finder";
import { PreviewProductDto, ProductFilterDto } from "@/app/types/data-type";
import { Card } from "@/app/ui/card";
import { useToast } from "@/app/ui/toast";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState, useRef } from "react";

const FILTER: ProductFilterDto = { mark: '', feature: '', tag: 'پیشنهادی' };

export default function SuggestedPage() {
    const { getCheapestProductsByFilter, error: hookError } = useFinder();
    const [products, setProducts] = useState<PreviewProductDto[]>([]);
    const [page, setPage] = useState(0);
    const [hasNext, setHasNext] = useState(true);
    const [isInitialLoading, setIsInitialLoading] = useState(true);
    const [isLoadingMore, setIsLoadingMore] = useState(false);
    const isLoadingRef = useRef(false);
    const router = useRouter();
    const toast = useToast();

    const fetchProducts = useCallback(async (pageNum: number) => {
        if (isLoadingRef.current) return null;
        
        isLoadingRef.current = true;
        
        try {
            const data = await getCheapestProductsByFilter(FILTER ,pageNum);
            return data;
        } catch {
            toast.error(hookError || "مشکل در دریافت اطلاعات");
            return null;
        } finally {
            isLoadingRef.current = false;
        }
    }, [getCheapestProductsByFilter, hookError, toast]);

    useEffect(() => {
        const loadInitialData = async () => {
            setIsInitialLoading(true);
            setProducts([]);
            setPage(0);
            setHasNext(true);
            
            const data = await fetchProducts(0);
            if (data) {
                setProducts(data.products);
                setHasNext(data.has_next);
            }
            setIsInitialLoading(false);
        };
        
        loadInitialData();
    }, [fetchProducts]);

    const loadMore = useCallback(async () => {
        if (isLoadingRef.current || !hasNext || isLoadingMore) return;
        
        const nextPage = page + 1;
        setIsLoadingMore(true);
        
        const data = await fetchProducts(nextPage);
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
    }, [page, hasNext, products, fetchProducts, isLoadingMore]);

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
            <div className="max-w-8xl mx-57 mt-2">
                <div className="flex">
                    <div className="flex-1">
                        <div className="mb-8">
                            <h1 className="text-3xl font-bold text-[rgb(217,47,74)] mb-2 mr-2 text-right">
                                با این پیشنهاد های شگفت انگیز استایلت رو کامل کن
                            </h1>
                            <p className="text-gray-600 text-right mr-2">
                                با اکسسوری زیبا تر می شویم ؟ چرا که نه
                            </p>
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
                </div>
            </div>
        );
    }

    return (
        <div className="max-w-8xl mx-[0.1vw] lg:mx-[9.1vw] xl:mx-[13.1vw] 2xl:mx-[12.3vw] lg:mt-2">
            <div className="flex">
                <div className="flex-1">
                    <div className="mb-8">
                        <h1 className="text-3xl font-bold text-[rgb(217,47,74)] mb-2 mr-2 text-right">
                            با این پیشنهاد های شگفت انگیز استایلت رو کامل کن
                        </h1>
                        <p className="text-gray-600 text-right mr-2">
                            با اکسسوری زیبا تر می شویم ؟ چرا که نه
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
                                        key={product.id} // ← کلید پایدار
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
            </div>
        </div>
    );
}