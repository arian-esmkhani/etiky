import { ElementType, ReactNode, useState } from 'react'
import { cn } from "@/app/lib/utils"
import Image from 'next/image'
import { TomanIcon } from './toman'

type CardVariant = 'magical' | 'neutral'
type CardSize = 'sm' | 'md' | 'lg' | 'xl' | 'found'

export interface CardProps {
  as?: 'article' | 'aside' | 'section'
  title?: string
  imgSrc?: string
  imgAlt?: string
  children?: ReactNode // Custom content
  variant?: CardVariant
  size?: CardSize
  className?: string
  onClick?: () => void
  price?: number
  oldPrice?: number
  lazyLoad?: boolean // Lazy load image
  discountPercent?: number
}

const base = 'group relative flex flex-col items-center overflow-hidden'

const variants: Record<CardVariant, string> = {
  magical: 'border border-neutral-200 bg-white',
  neutral: '',
}

const sizes: Record<CardSize, string> = {
  sm: 'w-19 h-19 sm:h-24 sm:w-28 text-center',
  md: 'h-[227px] sm:min-h-[299px] w-32 sm:w-42 p-5 text-center ',
  lg: 'w-[81vw]  sm:w-126 h-50 p-3 sm:p-5 flex-row gap-7',
  xl: 'sm:w-full h-[30vh] sm:h-118',
  found:' sm:min-h-[395px] lg:w-[20.7vw] xl:w-[18.27vw] 2xl:w-[15vw] p-5 text-center flex-row sm:flex-col'
}

function Skeleton({ className }: { className?: string }) {
  return (
    <div 
      className={cn(
        "relative overflow-hidden bg-neutral-200/70",
        "before:absolute before:inset-0 before:-translate-x-full",
        "before:animate-[shimmer_1.5s_infinite]",
        "before:bg-linear-to-r before:from-transparent before:via-white/60 before:to-transparent",
        className
      )}
    />
  )
}

/**
 * Versatile card component for displaying products, articles, or content.
 * 
 * Features:
 * - Multiple sizes (sm, md, lg, xl, found)
 * - Image with lazy loading and skeleton placeholder
 * - Price display with discount formatting
 * - Toman currency icon integration
 * - Custom HTML wrapper (article/aside/section)
 * - Loading and error states for images
 * - Responsive design
*/
export function Card({
  as = 'article',
  title,
  price,
  oldPrice,
  discountPercent,
  imgSrc,
  imgAlt = "card image",
  children,
  variant = 'neutral',
  size = 'md',
  lazyLoad = true,
  className,
  onClick,
}: CardProps) {
  const Wrapper: ElementType = as
  const [isLoading, setIsLoading] = useState(true)
  const [hasError, setHasError] = useState(false)

  const formatPrice = (value?: number) => {
    if (!value) return ''
    return value.toLocaleString('fa-IR')
  }

  const getImageSize = () => {
    switch(size) {
      case 'sm': return { width: 96, height: 96 }
      case 'md': return { width: 140, height: 140 }
      case 'lg': return { width: 168, height: 168 }
      case 'xl': return { width: 800, height: 472 }
      case 'found': return { width: 196, height: 196 }
      default: return { width: 140, height: 140 }
    }
  }

  const imageSize = getImageSize()

  return (
    <Wrapper
      className={cn(base, variants[variant], sizes[size], className)}
      onClick={onClick}
    >
      {/* Image section */}
      {imgSrc && (
        <div className={cn(
          "relative overflow-hidden shrink-0 select-none",
          size === 'sm' && "w-14 h-14 sm:w-24 sm:h-24 rounded-full",
          size === 'md' && "w-26 h-28 sm:h-35 sm:w-35 aspect-square rounded-lg",
          size === 'lg' && "w-34 sm:w-42 h-42",
          size === 'xl' && "w-full h-full",
          size === 'found' && "w-49 aspect-square rounded-lg",
        )}>
          {/* Skeleton placeholder while loading */}
          {isLoading && (
            <Skeleton 
              className={cn(
                "absolute inset-0 z-10",
                size === 'sm' && "rounded-full",
                (size === 'md' || size === 'found') && "rounded-lg",
                size === 'lg' && "rounded-lg",
                size === 'xl' && "rounded-none"
              )}
            />
          )}

          {/* Main image with Next.js optimization */}
          <Image
            src={imgSrc}
            alt={imgAlt}
            width={imageSize.width}
            height={imageSize.height}
            className={cn(
              "object-fill sm:object-cover transition-opacity duration-300",
              isLoading ? "opacity-0" : "opacity-100",
              size === 'xl' && "w-full h-full"
            )}
            loading={lazyLoad ? "lazy" : "eager"} // Performance optimization
            onLoad={() => setIsLoading(false)}
            onError={() => {
              setIsLoading(false)
              setHasError(true)
            }}
          />

          {/* Error fallback */}
          {hasError && (
            <div className="absolute inset-0 flex items-center justify-center bg-neutral-100 text-neutral-400 text-sm">
              <span>خطا در بارگذاری</span>
            </div>
          )}
        </div>
      )}

      {/* Content section - hidden for XL (uses absolute positioned children) */}
      {size !== 'xl' && (
        <div className={cn(
          "flex flex-col w-full",
          "items-center text-center font-bold",
        )}>
          {/* Title skeleton while loading */}
          {title && isLoading && (
            <Skeleton className={cn(
              "h-3 rounded",
              size === 'md' && "w-20 mt-2",
              size === 'sm' && "w-16 mt-1",
              size === 'lg' && "w-32",
              size === 'found' && "w-28 mt-4",
            )} />
          )}

          {/* Title */}
          {title && !isLoading && (
            <h3 className={cn(
              "text-neutral-800/81 wrap-break-word line-clamp-2 select-none",
              size === 'md' && "text-[1.14vh] sm:text-[11px] mt-0.5 sm:mt-2",
              size === 'sm' && "text-xs",
              size === 'lg' && "text-[1.33vh] sm:text-[15px]",
              size === 'found' && "text-[12.5px] mt-4",
            )}>
              {title}
            </h3>
          )}

          {/* Price section */}
          {(price || oldPrice) && (
            <div className="flex flex-col items-center gap-1 mt-1 sm:mt-4 w-full select-none">
              {/* Price skeleton while loading */}
              {(price || oldPrice) && isLoading ? (
                <div className="flex items-center gap-2">
                  <Skeleton className="w-12 h-4 rounded" />
                  <Skeleton className="w-16 h-4 rounded" />
                </div>
              ) : (
                <>
                  {/* Discount badge and old price */}
                  {discountPercent != 0 && discountPercent !== undefined && (
                    <div className="flex items-center gap-2">
                        <span className="text-[0.65rem] sm:text-xs font-bold text-red-50
                        bg-red-600 px-[0.3rem] py-[0.06rem] sm:px-1.5 sm:py-0.5 rounded-3xl">
                            {discountPercent.toLocaleString('fa-IR')}٪
                        </span>
                        <span className="text-neutral-400 text-[0.6rem] sm:text-xs line-through
                        flex items-center gap-0.5">
                          {formatPrice(oldPrice)}
                          <TomanIcon className="w-3 h-3 text-neutral-400" />
                        </span>
                    </div>
                  )}
                  {/* Current price */}
                  <span className={cn(
                    "font-bold text-neutral-950 flex items-center gap-0.5 text-md sm:text-xl",
                    size === 'md' && "font-bold sm:font-extrabold",
                    size === 'lg' && "font-extrabold",
                  )}>
                    {formatPrice(price)}
                    <TomanIcon className="w-3.5 h-3.5 text-black font-extrabold" />
                  </span>
                </>
              )}
            </div>
          )}
          
          {/* Custom children content */}
          {children && <div className="w-full mt-auto">{children}</div>}
        </div>
      )}

      {/* XL size: children are absolutely positioned over the image */}
      {size === 'xl' && children && (
        <div className="absolute inset-0 z-10 w-full h-full select-none">
          {children}
        </div>
      )}
    </Wrapper>
  )
}