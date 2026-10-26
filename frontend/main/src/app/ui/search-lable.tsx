"use client"

import { Search } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import { cn } from "../lib/utils";

export interface SearchProps {
  className?: string
  placeholder?: string
  delay?: number; // Debounce delay in ms
  value?: string; // Controlled value from parent
  onChange?: (value: string) => void;
  onSearchClick?: () => void; // Custom search button handler
  onEnterPress?: () => void; // Custom Enter key handler
  suggestions?: string[]; // Auto-complete suggestions list
  onSuggestionClick?: (suggestion: string) => void;
  isLoading?: boolean; // Loading state for suggestions
}

/**
 * Advanced search component with auto-complete and keyboard navigation.
 * Features:
 * - Debounced onChange
 * - Keyboard navigation (ArrowUp/Down, Enter, Escape)
 * - Click outside to close
 * - Auto-scroll to selected suggestion
 * - Loading state support
*/
export default function SearchLabel({
  className,
  placeholder = 'search ...',
  value = "",
  onChange,
  delay = 100,
  onSearchClick,
  onEnterPress,
  suggestions = [],
  onSuggestionClick,
  isLoading = false
}: SearchProps) {
  const [query, setQuery] = useState(value);
  const [isOpen, setIsOpen] = useState(false); // Dropdown visibility state

  // Currently selected suggestion index (-1 = none)
  const [selectedIndex, setSelectedIndex] = useState(-1);
  
  // Refs for DOM manipulation
  const inputRef = useRef<HTMLInputElement>(null);
  const wrapperRef = useRef<HTMLDivElement>(null);
  const listRef = useRef<HTMLUListElement>(null);

  useEffect(() => {
    setQuery(value);
  }, [value]);

  /**
   * Debounced onChange handler.
   * Waits for user to stop typing before triggering onChange.
  */
  useEffect(() => {
    if (onChange) {
      const timeoutId = setTimeout(() => {
        onChange(query);
      }, delay);
      
      return () => clearTimeout(timeoutId);
    }
  }, [query, onChange, delay]);

  /**
   * Close suggestions when clicking outside the component.
   * Standard dropdown pattern for better UX.
  */
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (wrapperRef.current && !wrapperRef.current.contains(event.target as Node)) {
        setIsOpen(false);
        setSelectedIndex(-1);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  /**
   * Auto-scroll to the selected suggestion when navigating with keyboard.
   * Ensures the selected item is always visible.
  */
  useEffect(() => {
    if (selectedIndex >= 0 && listRef.current) {
      const selectedElement = listRef.current.children[selectedIndex] as HTMLElement;
      if (selectedElement) {
        selectedElement.scrollIntoView({
          block: 'nearest',
          behavior: 'smooth'
        });
      }
    }
  }, [selectedIndex]);

  // Input change handler - updates query and shows suggestions
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = e.target.value;
    setQuery(newValue);
    setIsOpen(true);
    setSelectedIndex(-1);
  };

  // Clear input and reset all states
  const handleClear = () => {
    setQuery("");
    setIsOpen(false);
    setSelectedIndex(-1);
    if (onChange) onChange("");
    inputRef.current?.focus();
  };

  // Handle search icon click with priority: onSearchClick > onChange
  const handleIconClick = () => {
    if (onSearchClick) {
      onSearchClick();
    } else if (query.trim() && onChange) {
      onChange(query.trim());
    }
    inputRef.current?.focus();
  };

  // Handle suggestion selection from click
  const handleSuggestionClick = (suggestion: string) => {
    setQuery(suggestion);
    setIsOpen(false);
    setSelectedIndex(-1);
    if (onSuggestionClick) {
      onSuggestionClick(suggestion);
    }
    if (onChange) {
      onChange(suggestion);
    }
    inputRef.current?.focus();
  };

  /**
   * Keyboard navigation handler.
   * Supports: ArrowDown, ArrowUp, Enter, Escape
  */
  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    const hasSuggestions = suggestions.length > 0 && isOpen;

    switch (e.key) {
      case 'ArrowDown':
        e.preventDefault();
        if (hasSuggestions) {
          setSelectedIndex(prev => 
            prev < suggestions.length - 1 ? prev + 1 : prev
          );
        }
        break;

      case 'ArrowUp':
        e.preventDefault();
        if (hasSuggestions) {
          setSelectedIndex(prev => prev > 0 ? prev - 1 : -1);
        }
        break;

      case 'Enter':
        e.preventDefault();
        if (selectedIndex >= 0 && suggestions[selectedIndex]) {
          handleSuggestionClick(suggestions[selectedIndex]);
        } else if (onEnterPress) {
          setIsOpen(false);
          setSelectedIndex(-1);
          onEnterPress();
        } else if (onSearchClick) {
          setIsOpen(false);
          setSelectedIndex(-1);
          onSearchClick();
        } else if (onChange && query.trim()) {
          setIsOpen(false);
          setSelectedIndex(-1);
          onChange(query.trim());
        }
        break;

      case 'Escape':
        e.preventDefault();
        setIsOpen(false);
        setSelectedIndex(-1);
        inputRef.current?.blur();
        break;

      default:
        break;
    }
  };

  // Show suggestions only if: open, has query, and suggestions exist
  const showSuggestions = isOpen && query.trim().length > 0 && suggestions.length > 0;

  return (
    <div ref={wrapperRef} className="relative">
      <div className={cn('relative w-[60vw] sm:w-[70vw] md:w-98 xl:w-102', [className])}>
        <Search 
          className="absolute right-3 top-2.5 sm:top-3.5 text-gray-500 w-4.5 h-4 sm:h-4.5 cursor-pointer"
          onClick={handleIconClick} 
        />
        <input
          type="text"
          value={query}
          ref={inputRef}
          onChange={handleChange}
          placeholder={placeholder}
          onKeyDown={handleKeyDown}
          onFocus={() => {
            if (query.trim().length > 0 && suggestions.length > 0) {
              setIsOpen(true);
            }
          }}
          className="w-full pr-10 pl-4 py-[0.4rem] sm:py-[0.7rem] rounded-3xl bg-gray-300/40 border
                  border-gray-600/40 text-sm 
                  text-black/85 placeholder-gray-500 transition-all duration-300 text-right
                  focus:outline-none focus:ring-0 hover:bg-gray-300/50"
        />
        {query && (
          <button
            onClick={handleClear}
            className="absolute left-4 top-3 text-gray-400 hover:text-gray-600"
            type="button"
          >
            ✕
          </button>
        )}
      </div>

      {showSuggestions && (
        <div className="absolute top-full left-0 right-0 mt-1 bg-gray-300/90 rounded-2xl shadow-lg border border-gray-200 max-h-60 overflow-y-auto z-50">
          {isLoading ? (
            <div className="px-4 py-2 text-gray-500 text-sm text-center">در حال بارگذاری...</div>
          ) : (
            <ul ref={listRef} className="py-1">
              {suggestions.map((suggestion, index) => (
                <li
                  key={index}
                  onClick={() => handleSuggestionClick(suggestion)}
                  onMouseEnter={() => setSelectedIndex(index)}
                  className={`px-4 py-2 cursor-pointer text-sm text-right transition-colors duration-150
                    ${selectedIndex === index 
                      ? 'bg-rose-50/30 text-rose-700' 
                      : 'hover:bg-gray-100'
                    }`}
                >
                  {suggestion}
                </li>
              ))}
            </ul>
          )}
        </div>
      )}
    </div>
  );
}