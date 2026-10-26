'use client'

import { useState, useCallback, useEffect} from 'react'
import { NavButton } from '@/app/ui/nav-button'
import { Card, CardProps } from '@/app/ui/card'
import useEmblaCarousel from 'embla-carousel-react'
import { cn } from '@/app/lib/utils'
import Autoplay from 'embla-carousel-autoplay'

interface InfiniteSliderProps {
  items: CardProps[]
  className?: string
  showNavigation?: boolean
  slideDivClassName?: string // Additional styles for slide track
  onActiveSlideChange?: (activeSlide: CardProps, index: number) => void // Callback when slide changes
  loop?: boolean // Enable infinite looping
  autoplay?: boolean // Auto-play slides
  autoplayDelay?: number // Auto-play interval in ms
  align?: 'start' | 'center' | 'end'
  dragFree?: boolean
  showActiveCard?: boolean
  showDots?: boolean
  dotsPosition?: 'bottom' | 'top' | 'inside'
  dotClassName?: string
  activeDotClassName?: string
}

/**
 * Infinite carousel component using Embla Carousel.
 * Displays cards in a scrollable, swipeable carousel with navigation.
 * 
 * Features:
 * - Infinite looping (optional)
 * - Auto-play with configurable delay
 * - Drag/swipe support
 * - Navigation buttons with enabled/disabled states
 * - Dot indicators with custom styling
 * - Active card highlighting (scale effect)
 * - RTL support (Persian/Arabic)
 * - Responsive card sizing
 * - Callback for active slide changes
*/
export function InfiniteSlider({ 
  items, 
  className,
  showNavigation = true,
  slideDivClassName = "",
  onActiveSlideChange ,
  loop = false,
  autoplay = false,
  autoplayDelay = 5000,
  align = 'start',
  showActiveCard = false,
  dragFree = true,
  showDots = false,
  dotsPosition = 'bottom',
  dotClassName = '',
  activeDotClassName = '',
}: InfiniteSliderProps) {
  const [emblaRef, emblaApi] = useEmblaCarousel({ 
    align,
    loop,
    skipSnaps: false,
    dragFree,
    direction: 'rtl',
  }, autoplay ? [Autoplay({ 
  delay: autoplayDelay, 
  stopOnInteraction: false,  // Continue after user interaction
  stopOnMouseEnter: true, // Pause on hover
})] : [])

  // State for navigation and slide tracking
  const [prevBtnEnabled, setPrevBtnEnabled] = useState(false)
  const [nextBtnEnabled, setNextBtnEnabled] = useState(false)
  const [activeIndex, setActiveIndex] = useState(0) // Currently active (for scaling)
  const [selectedIndex, setSelectedIndex] = useState(0) // Currently selected (for dots)
  const [scrollSnaps, setScrollSnaps] = useState<number[]>([]) // Total slides count

  // Navigation handlers
  const scrollPrev = useCallback(() => {
    if (emblaApi) emblaApi.scrollPrev()
  }, [emblaApi])

  const scrollNext = useCallback(() => {
    if (emblaApi) emblaApi.scrollNext()
  }, [emblaApi])

  const scrollTo = useCallback((index: number) => {
    if (emblaApi) emblaApi.scrollTo(index)
  }, [emblaApi])

  /**
   * Updates navigation button states and active slide information.
   * Uses requestAnimationFrame for smooth updates.
  */
  const updateButtonState = useCallback(() => {
    if (!emblaApi) return
    
    // Schedule update on next frame for better performance
    requestAnimationFrame(() => {
      setPrevBtnEnabled(emblaApi.canScrollPrev())
      setNextBtnEnabled(emblaApi.canScrollNext())
      
      const newActiveIndex = emblaApi.selectedScrollSnap()
      setActiveIndex(newActiveIndex)
      
      // Notify parent about slide change
      if (onActiveSlideChange) {
        onActiveSlideChange(items[newActiveIndex], newActiveIndex)
      }
    })
  }, [emblaApi, items, onActiveSlideChange])

  /**
   * Set up Embla event listeners on mount.
   * - onSelect: Triggered when slide changes
   * - onReInit: Triggered when carousel re-initializes
  */
  useEffect(() => {
    if (!emblaApi) return
    
    // Initial update
    updateButtonState()
    setScrollSnaps(emblaApi.scrollSnapList())
    
    // Event handlers
    const onSelect = () => {
      setSelectedIndex(emblaApi.selectedScrollSnap())
      updateButtonState()
    }
    
    emblaApi.on('select', onSelect)
    emblaApi.on('reInit', onSelect)

    // Cleanup event listeners
    return () => {
      emblaApi.off('select', onSelect)
      emblaApi.off('reInit', onSelect)
    }
  }, [emblaApi, updateButtonState])

  return (
    <div className={cn('relative', className)}>
      {/* Carousel viewport */}
      <section className="overflow-hidden" ref={emblaRef} dir="rtl">
        {/* Slide track */}
        <div className={cn('flex', slideDivClassName)}>
          {items.map((item, index) => (
            <div 
              key={index} 
              className={cn(
                item.size === 'xl' && 'flex-[0_0_100%] h-full',
                item.size === 'lg' && 'px-2',
                item.size === 'md' && 'px-0.5',
                item.size === 'sm' && 'px-[0.1px] sm:px-3',
                showActiveCard && index === activeIndex && 'scale-109'
              )}
            >
              <Card {...item} />
            </div>
          ))}
        </div>
      </section>
      
      {/* Navigation buttons */}
      {showNavigation && (
        <>
          <span className="absolute top-1/2 left-4 transform -translate-y-1/2">
            <NavButton
              variant="left"
              onClick={scrollNext}
              disabled={!nextBtnEnabled}
            />
          </span>
          
          <span className="absolute top-1/2 right-4 transform -translate-y-1/2">
            <NavButton
              variant="right"
              onClick={scrollPrev}
              disabled={!prevBtnEnabled}
            />
          </span>
        </>
      )}

      {/* Dot indicators */}
      {showDots && (
        <div className={cn(
          'flex flex-row-reverse justify-center gap-2',
          dotsPosition === 'bottom' && 'mt-4',
          dotsPosition === 'top' && 'mb-4',
          dotsPosition === 'inside' && 'absolute bottom-4 left-1/2 -translate-x-1/2 z-10'
        )}>
          {scrollSnaps.map((_, index) => (
            <button
              key={index}
              onClick={() => scrollTo(index)}
              className={cn(
                'w-1.5 h-1.5 rounded-full transition-all duration-300',
                index === selectedIndex 
                  ? 'bg-red-100 w-4' 
                  : 'bg-white',
                dotClassName,
                index === selectedIndex && activeDotClassName
              )}
              aria-label={`Go to slide ${index + 1}`}
            />
          ))}
        </div>
      )}
    </div>
  )
}