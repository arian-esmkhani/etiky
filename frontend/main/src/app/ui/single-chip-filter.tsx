"use client";

import React, { useState } from "react";

interface SingleChipFilterProps {
    title: string;
    options: string[];
    defaultSelected?: string;
    onSelect?: (selected: string | null) => void;
}

/**
 * A single-select chip filter component.
 * Allows users to select one option from a list of chips.
 * Clicking an already selected chip deselects it (toggle off).
 * 
 * Features:
 * - Single selection only (mutually exclusive)
 * - Toggle on/off by clicking the same chip
 * - Visual feedback with checkmark icon when selected
 * - Responsive design with glassmorphism
*/
const SingleChipFilter: React.FC<SingleChipFilterProps> = ({
    title,
    options,
    defaultSelected = "",
    onSelect,
}) => {
// Track currently selected option (null = none selected)
const [selected, setSelected] = useState<string | null>(defaultSelected || null);

/**
 * Handles chip click with toggle behavior:
 * - If chip is already selected → deselect it (set to null)
 * - If chip is not selected → select it
 * - Notifies parent via onSelect callback
*/
const handleChipClick = (option: string) => {
const newSelected = selected === option ? null : option;
setSelected(newSelected);
if (onSelect) {
    onSelect(newSelected);
}
};

return (
    <div className="w-full lg:max-w-md lg:mx-auto p-2 sm:p-4 bg-white/35 lg:bg-white rounded-xl 
    shadow-sm border border-gray-100">
        <h3 className="text-xs font-medium text-gray-500 mb-3 ml-4">{title}</h3>
        <div className="flex flex-wrap items-end gap-0.5 sm:gap-2">
        {options.map((option) => {
            const isSelected = selected === option;
            return (
            <button
                key={option}
                onClick={() => handleChipClick(option)}
                className={`px-3 py-1.5 rounded-full text-xs font-medium transition-all duration-150 border ${
                isSelected
                    ? "bg-red-500 text-white border-none"
                    : "bg-white/36 lg:bg-white text-gray-600 hover:bg-gray-200 border-none"
                }`}
            >
                {isSelected && (
                <svg className="w-3 h-3 inline-block ml-1" fill="currentColor" viewBox="0 0 20 20">
                    <path
                    fillRule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clipRule="evenodd"
                    />
                </svg>
                )}
                {option}
            </button>
            );
        })}
        </div>
    </div>
    );
};

export default SingleChipFilter;