"use client"

import { ReactNode } from "react"
import { cn } from "../lib/utils"
import { Loader2 } from "lucide-react"

type ButtonSize = "sm" | "lg"
type ButtonTone = "primary" | "secondary"

export interface ButtonProps {
  size?: ButtonSize
  tone?: ButtonTone
  className?: string
  onClick?: () => void
  children?: ReactNode
  disabled?: boolean
  loading?: boolean
}

const base =
  "inline-flex items-center justify-center rounded-md font-medium focus:outline-none transition-colors duration-200 ease-in-out shadow-sm active:scale-95"  // Subtle press feedback

const sizes: Record<ButtonSize, string> = {
  sm: "px-2 py-1.5 text-sm",
  lg: "px-4 py-2 text-base",
}

const tones: Record<ButtonTone, string> = {
  primary:
    "bg-[rgb(217,47,74)] text-white hover:from-indigo-600 hover:to-blue-700 disabled:opacity-50 disabled:cursor-not-allowed",
  secondary:
    "bg-white text-gray-700 border border-gray-300 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed",
}

export function AppButton({
  size = "sm",
  tone = "primary",
  className = "",
  onClick,
  children,
  disabled = false,
  loading = false,
}: React.PropsWithChildren<ButtonProps>) {
  return (
    <button
      onClick={onClick}
      disabled={disabled || loading} // Loading takes precedence over disabled
      className={cn(base, sizes[size], tones[tone], className)}
    >
      {loading && (
        <Loader2 className="mr-2 h-4 w-4 animate-spin" />
      )}
      {children}
    </button>
  )
}
