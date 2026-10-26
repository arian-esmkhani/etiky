'use client'

import { useRouter } from "next/navigation";
import { heroCardItems } from "../../lib/data";
import { InfiniteSlider } from "../../ui/carousel";

export function HeroSlider() {
    const router = useRouter();
    
    const itemsWithNavigation = heroCardItems.map(item => ({
        ...item,
        onClick: () => router.push(item.link),
    }));

    return(
        <InfiniteSlider
        items={itemsWithNavigation}
        showNavigation={false}
        dragFree={false}
        loop={true}
        autoplay={true}
        align="center"
        showDots={true}
        dotsPosition="inside"/>
    )
}