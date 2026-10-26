"use client"

import React, { useCallback, useRef, useState } from "react";
import { Range, getTrackBackground } from "react-range";
import { TomanIcon } from "./toman";

interface PriceRangeSliderProps {
    min?: number;
    max?: number;
    step?: number;
    onRangeChange?: (values: { min: number; max: number }) => void;
    delay?: number;
}

const PriceRangeSlider: React.FC<PriceRangeSliderProps> = ({
    min = 0,
    max = 1000,
    step = 10,
    onRangeChange,
    delay = 400, // Default debounce delay
}) => {
    const [values, setValues] = useState<[number, number]>([min, max / 4]);

    // Tracks if user is actively dragging a thumb
    const [isDragging, setIsDragging] = useState(false);

    // Timeout refs for debouncing
    const timeoutRef = useRef<NodeJS.Timeout | null>(null); // For non-drag changes
    const dragTimeoutRef = useRef<NodeJS.Timeout | null>(null); // For drag changes

    /**
     * Final callback executor - sends final values to parent
    */
    const applyRangeChange = useCallback((newValues: [number, number]) => {
        if (onRangeChange) {
            onRangeChange({ min: newValues[0], max: newValues[1] });
        }
    }, [onRangeChange]);

    /**
     * Handles range changes with debouncing.
     * - If dragging: uses dragTimeoutRef
     * - If not dragging: uses timeoutRef
     * Prevents excessive callback calls during drag
    */
    const handleChange = useCallback((newValues: number[]) => {
        const newValuesTuple = newValues as [number, number];
        setValues(newValuesTuple);

        if (timeoutRef.current) {
            clearTimeout(timeoutRef.current);
        }
        if (dragTimeoutRef.current) {
            clearTimeout(dragTimeoutRef.current);
        }

        // Set new timeout based on drag state
        if (isDragging) {
            dragTimeoutRef.current = setTimeout(() => {
                applyRangeChange(newValuesTuple);
            }, delay);
        } else {
            timeoutRef.current = setTimeout(() => {
                applyRangeChange(newValuesTuple);
            }, delay);
        }
    }, [isDragging, delay, applyRangeChange]);

    React.useEffect(() => {
        return () => {
            if (timeoutRef.current) {
                clearTimeout(timeoutRef.current);
            }
            if (dragTimeoutRef.current) {
                clearTimeout(dragTimeoutRef.current);
            }
        };
    }, []);

    return (
        <div className="w-full max-w-md mx-auto p-4 bg-white/35 
        lg:bg-white backdrop-blur-sm rounded-2xl shadow-sm border 
        border-gray-100/50 transition-all hover:shadow-md">
            <div className="flex justify-between items-center mb-6">
                <span className="text-xs md:text-sm font-medium text-gray-500/">بازه قیمت</span>
                <div className="flex items-center gap-3 text-sm font-mono">
                    <span className="px-3 py-1 bg-gray-50/30 rounded-lg text-gray-700">
                        {values[0].toLocaleString('fa-IR')}
                        <TomanIcon className="w-3 h-3 text-gray-500" />
                    </span>
                    <span className="text-gray-300">—</span>
                    <span className="px-3 py-1 bg-gray-50/30 rounded-lg text-gray-700">
                        {values[1].toLocaleString('fa-IR')}
                        <TomanIcon className="w-3 h-3 text-gray-500" />
                    </span>
                </div>
            </div>

            <Range
                values={values}
                step={step}
                min={min}
                max={max}
                onChange={handleChange}
                renderTrack={({ props, children }) => (
                    <div
                        {...props}
                        style={{
                            ...props.style,
                            height: "6px",
                            width: "100%",
                            background: getTrackBackground({
                                values,
                                colors: ["#e5e7eb", "rgb(217,47,74)", "#e5e7eb"],
                                min,
                                max,
                            }),
                            borderRadius: "4px",
                        }}
                        className="cursor-pointer"
                    >
                        {children}
                    </div>
                )}
                renderThumb={({ props, isDragged }) => {
                    if (isDragged !== isDragging) {
                        setIsDragging(isDragged);
                        
                        if (!isDragged && dragTimeoutRef.current) {
                        }
                    }
                    
                    return (
                        <div
                            {...props}
                            style={{
                                ...props.style,
                                height: "20px",
                                width: "20px",
                                borderRadius: "50%",
                                background: "#ffffff",
                                border: "2px solid rgb(217,47,74)",
                                display: "flex",
                                justifyContent: "center",
                                alignItems: "center",
                                transition: "all 0.15s ease",
                            }}
                        >
                            <div
                                style={{
                                    width: "4px",
                                    height: "4px",
                                    borderRadius: "50%",
                                    background: isDragged ? "rgb(217,47,74)" : "rgb(217,47,74)",
                                }}
                            />
                        </div>
                    );
                }}
            />
        </div>
    );
};

export default PriceRangeSlider;