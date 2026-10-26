import { cn } from "@/app/lib/utils"
import { ChevronLeft, ChevronRight } from 'lucide-react'

type NavButtonVariant = 'left' | 'right'

export interface NavButtonProps {
  variant: NavButtonVariant
  className?: string
  onClick?: () => void
  disabled?: boolean
}

const base = 'inline-flex items-center justify-center w-9 h-9 bg-white/80 rounded-full border border-neutral-200/60 text-neutral-400 hover:text-neutral-800 active:translate-x-0.5 disabled:opacity-0 disabled:pointer-events-none transition-all duration-100'

export function NavButton({
  variant,
  className,
  onClick,
  disabled = false,
}: NavButtonProps) {
  // Maps variant to corresponding icon direction
  const Icon = variant === 'right' ? ChevronRight : ChevronLeft
  const iconSize = 18

  return (
    <button
      className={cn(base, className)}
      onClick={onClick}
      disabled={disabled}
      aria-label={variant === 'right' ? 'Next' : 'Previous'} // Screen reader support
    >
      <Icon size={iconSize} />
    </button>
  )
}