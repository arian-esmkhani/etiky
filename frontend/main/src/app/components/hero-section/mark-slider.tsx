'use client'

import { useRouter } from "next/navigation";
import { markCardItems } from "../../lib/data";
import { InfiniteSlider } from "../../ui/carousel";

export function MarkSlider() {
    const router = useRouter();

    const itemsWithNavigation = markCardItems.map(item => ({
        ...item,
        onClick: () => router.push(item.link),
    }));

    return(
        <InfiniteSlider 
            items={itemsWithNavigation}
            className="mx-[1.7em] lg:mx-[8em] xl:mx-[11vw] mb-3 sm:my-4"
            showNavigation={false}
            align="end"
        />
    )
}