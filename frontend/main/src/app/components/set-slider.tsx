'use client'

import { useEffect, useMemo, useState } from "react";
import { InfiniteSlider } from "../ui/carousel";
import { MoreNav } from "../ui/more-nav";
import { PreviewProductDto, ProductFilterDto } from "../types/data-type";
import { useSlider } from "../hooks/use-slider";
import { useRouter } from "next/navigation";

const FILTER: ProductFilterDto = { mark: '', feature: '', tag: 'ست' };

export function SetSlider() {
    const [cards, setCards] = useState<PreviewProductDto[]>([]);
    const { getByFilter, error } = useSlider()
    const router = useRouter();

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
            imgSrc:"/api" + item.img_url,
            imgAlt:item.name,
            title:item.name,
            variant:"magical" as const,
            price:item.final_price,
            oldPrice:item.price,
            lazyLoad: true,
            discountPercent:item.discount,
            size: "lg" as const,
            onClick: () => router.push(`/components/${item.id.toString()}/product`),
        }));
    }, [cards, router])


    return (
        <div className="py-2 pb-9">
            <div className="max-w-7xl mx-auto sm:px-4">
                <div className="relative rounded-2xl overflow-hidden">
                    <div className="bg-[rgb(217,47,74)] px-6 py-5 sm:py-10">
                        <div className="flex items-center justify-between">
                            <MoreNav onClick={() => router.push("/")}/>
                            <h3 className="text-xl sm:text-2xl font-bold text-gray-300/80">
                                {"یک انتخاب دو نفره"}
                            </h3>
                        </div>
                    </div>

                    <div className="bg-white pr-1 sm:px-6 py-8 border-t-2 border-[rgb(217,47,74)]/20">
                        <div className="relative hidden sm:block">
                            <InfiniteSlider items={items} align="end"/>
                        </div>
                        <div className="relative sm:hidden">
                            <InfiniteSlider items={items} align="end" showNavigation={false}/>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}