'use client'

import { useEffect, useMemo, useState } from "react";
import { InfiniteSlider } from "../ui/carousel";
import { MoreNav } from "../ui/more-nav";
import { PreviewProductDto } from "../types/data-type";
import { useSlider } from "../hooks/use-slider";
import { useRouter } from "next/navigation";

export function BestSellingSlider() {
    const [cards, setCards] = useState<PreviewProductDto[]>([]);
    const { getBestSelling, error, loading } = useSlider()
    const router = useRouter();

    // Fetch data from server
    useEffect(() => {
        const fetchNewest = async () => {
            const data = await getBestSelling();
            if(error != null) {
                console.log(error)
            } else if (data && data.length > 0) {
                setCards(data);
            } 
        };

        fetchNewest();
    }, [getBestSelling, error])

    const items = useMemo(() => {
        return cards.map((item) => ({
            imgSrc: `/api${item.img_url}`,
            imgAlt: item.name,
            title: item.name,
            variant: "magical" as const,
            price: item.final_price,
            oldPrice: item.price,
            lazyLoad: true,
            discountPercent: item.discount,
            size: "md" as const,
            onClick: () => router.push(`/components/${item.id.toString()}/product`),
        }));
    }, [cards, router]);

    if (loading) {
        return (
            <div className="py-8 max-w-7xl mx-auto space-y-12">
                {[1, 2].map((_, idx) => (
                    <div key={idx} className="bg-white/70 p-6 rounded-2xl shadow-sm">
                        <div className="flex justify-between">
                            <div className="h-4 w-14 bg-gray-300 rounded animate-pulse" />
                            <div className="h-4 w-20 bg-gray-300 rounded animate-pulse" />
                        </div>
                        <div className="mt-6 flex gap-3">
                            {[1, 2, 3, 4].map((_, i) => (
                                <div key={i} className="h-32 w-40 bg-gray-300
                                rounded animate-pulse shrink-0" />
                            ))}
                        </div>
                    </div>
                ))}
            </div>
        );
    }

    return (
        <div className="py-8 space-y-7 sm:space-y-12 max-w-7xl mx-auto">
            <div>
                <div className="relative">
                    <div className="relative bg-white rounded-3xl shadow-gray-200/50 border border-gray-100/80 p-6">
                        <div className="flex items-center justify-between">
                            <MoreNav className="group relative px-4 py-2 text-sm font-medium 
                            text-[rgb(217,47,74)] hover:text-[rgb(250,91,117)]"
                            onClick={() => router.push("/")}/> 
                            <div className="flex items-center gap-3">
                                <h3 className="text-lg sm:text-xl text-[rgb(217,47,74)]">
                                    {"پرفروش ترین ها"}
                                </h3>
                            </div>
                        </div>
                        <div className="mt-6 relative hidden sm:block">
                            <InfiniteSlider items={items} align="end" />
                        </div>
                        <div className="mt-6 relative sm:hidden">
                            <InfiniteSlider items={items} align="end" showNavigation={false}/>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}