'use client'

import { InfiniteSlider } from "../ui/carousel";
import Image from "next/image";
import { MoreNav } from "../ui/more-nav";
import { useSlider } from "../hooks/use-slider";
import { useEffect, useMemo, useState } from "react";
import { PreviewProductDto } from "../types/data-type";
import { useRouter } from "next/navigation";

export function DiscountSlider() {
    const [cards, setCards] = useState<PreviewProductDto[]>([])
    const { getByDiscount, error } = useSlider()
    const router = useRouter();

    // Fetch data from server
    useEffect(() => {
        const fetchNewest = async () => {
            const data = await getByDiscount();
            if(error != null) {
                console.log(error)
            } else if (data && data.length > 0) {
                setCards(data);
            } 
        };

        fetchNewest();
    }, [getByDiscount, error])

    // Creating slider items
    const items = useMemo(() => {
        return cards.map((item) => ({
            imgSrc:"/api" + item.img_url,
            imgAlt:item.name,
            title:item.name,
            variant:"magical" as const,
            price : item.final_price,
            oldPrice : item.price,
            discountPercent : item.discount,
            size: "md" as const,
            onClick: () => router.push(`/components/${item.id.toString()}/product`),
        }));
    }, [cards, router])

    return(
        <div className=" pt-[3vh] sm:pt-0 py-2">
            <div className="flex items-center justify-between gap-3 min-h-[247px] sm:min-h-[299px]
                            xl:mx-[4.7vw] bg-[rgb(217,47,74)] rounded-3xl py-3 sm:py-5">
                <div className="flex-1 min-w-0 hidden sm:block">
                    <InfiniteSlider items={items} align="end"/>
                </div>
                <div className="flex-1 min-w-0 sm:hidden">
                    <InfiniteSlider items={items} align="end" showNavigation={false}/>
                </div>

                <div className="flex flex-col items-end shrink-0">
                    <div className="flex flex-col items-center leading-tight mr-3 sm:mr-5">
                        <span className="text-xs sm:text-2xl font-light text-gray-300">تخفیف</span>
                        <span className="text-xs sm:text-2xl font-light text-gray-300 -mt-1">های</span>
                        <span className="text-sm sm:text-3xl font-bold text-gray-200 mt-1">شگفت‌انگیز</span>
                    </div>
                    <Image 
                        src="/discount.svg" 
                        alt="discount" 
                        width={100} 
                        height={26} 
                        priority
                        className="mt-5 mr-8 hidden md:block"
                    />
                    <Image 
                        src="/discount.svg" 
                        alt="discount" 
                        width={55} 
                        height={26} 
                        priority
                        className="mt-2 mr-3 md:hidden"
                    />
                    <MoreNav 
                        className="mt-4 mr-1 sm:mr-9" 
                        onClick={() => router.push('/products/discount')} 
                    />
                </div>
            </div>
        </div>
    )
}