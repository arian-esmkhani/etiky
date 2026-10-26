'use client'

import { useEffect, useMemo, useState } from "react";
import { InfiniteSlider } from "../ui/carousel";
import { MoreNav } from "../ui/more-nav";
import { PreviewProductDto, ProductFilterDto} from "../types/data-type";
import { useSlider } from "../hooks/use-slider";
import { useRouter } from "next/navigation";

const FILTER: ProductFilterDto = { mark: '', feature: '', tag: 'ترند' };

export function TrendSlider() {
    const [cards, setCards] = useState<PreviewProductDto[]>([]);
    const { getByFilter, error, loading } = useSlider()
    const router = useRouter();

    // Fetch data from server
    useEffect(() => {
        const fetchNewest = async () => {
            const data = await getByFilter(FILTER);
            if(error != null) {
                console.log(error)
            } else if (data && data.length > 0) {
                setCards(data);
            } 
        };

        fetchNewest();
    }, [getByFilter, error])

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
                <div className="relative rounded-2xl bg-white/70 backdrop-blur-sm shadow-xl shadow-[rgb(207,192,194)] border border-white/50 p-6 transition-all duration-300">
                    <div className="absolute top-0 left-0 right-0 h-1 bg-linear-to-r from-transparent via-[rgb(217,47,74)] to-[rgb(126,0,21)]" />
                    
                    <div className="flex items-center justify-between">
                        <MoreNav className="group relative px-4 py-2 text-sm font-medium 
                        text-[rgb(217,47,74)] hover:text-[rgb(250,91,117)]"
                        onClick={() => router.push("/")}/>
                        <h3 className="text-lg sm:text-xl bg-linear-to-r from-[rgb(217,47,74)] 
                        to-[rgb(126,0,21)] bg-clip-text text-transparent">
                            {"جدید ترین ها"}
                        </h3>
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
    );
}