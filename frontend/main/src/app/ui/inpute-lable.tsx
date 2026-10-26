"use client"

import { cn } from "../lib/utils";

export interface InputProps {
  className?: string
  placeholder?: string
  value?: string;
  onChange?: (value: string) => void;
  type?: string;
  name?: string;
}

export function AppInput({
  className,
  placeholder = ' ...',
  value,
  onChange,
  type = 'text',
  name
}: InputProps) {

  return (
    <div className={cn("w-50 sm:w-65", className)}>
        <input
        type={type}
        name={name}
        value={value}
        onChange={(e) => onChange?.(e.target.value)}
        placeholder={placeholder}        
        className="p-1 px-2 bg-amber-200/43 rounded-3xl w-full border-2 
            border-amber-200 text-sm text-gray-800 
            focus:outline-none focus:border-amber-50/80
            dark:bg-gray-600/50 dark:border-gray-900/80 dark:focus:bg-gray-600/90
            dark:focus:border-gray-900
            transition-all duration-200"/>
    </div>
  );
}