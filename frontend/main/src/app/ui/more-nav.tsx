import { ChevronLeft } from 'lucide-react'
import { cn } from '../lib/utils';
import { useRouter } from 'next/navigation';

interface MoreNavProps {
    className?: string;
    href?: string;
    onClick?: () => void;
}

export function MoreNav({ className, href, onClick }: MoreNavProps) {
    const router = useRouter();

    const handleClick = () => {
        if (onClick) {
            onClick(); // Priority 1: Custom behavior
        } else if (href) {
            router.push(href); // Priority 2: Navigation
        }
    };

    return(
        <span 
            className={cn(
                "text-gray-300 text-sm cursor-pointer",
                "hover:text-gray-300/70 transition-colors duration-300",
                "flex items-center gap-0.5",
                className
            )}
            onClick={handleClick}
        >
            <ChevronLeft className="w-3 h-3 sm:w-4 sm:h-4" />
            <p className='text-xs sm:text-[1rem]'>مشاهده همه</p>
        </span>
    )
}