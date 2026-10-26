"use client";

import { createContext, useContext, useState, ReactNode, useCallback, useRef } from "react";

type ToastVariant = "success" | "error" | "warning" | "info";

interface ToastMessage {
    id: string;
    title?: string;
    description: string;
    variant: ToastVariant;
    timeout?: number;
    isRemoving?: boolean;
}

// Context API interface for toast methods
interface ToastContextType {
    success: (description: string, title?: string) => void;
    error: (description: string, title?: string) => void;
    warning: (description: string, title?: string) => void;
    info: (description: string, title?: string) => void;
    dismiss: (id: string) => void;
    dismissAll: () => void;
}

// Create context with null initial value (must be used within Provider)
const ToastContext = createContext<ToastContextType | null>(null);

/**
 * Toast notification system provider.
 * Provides toast methods to children components via context.
 * 
 * Features:
 * - Multiple toast variants (success, error, warning, info)
 * - Auto-dismiss with configurable timeout
 * - Maximum 3 toasts visible at once (FIFO queue)
 * - Smooth enter/exit animations
 * - Manual dismiss and dismiss all
 * - Centered fixed position at top of screen
*/
export function ToastProvider({ children }: { children: ReactNode }) {
    const [toasts, setToasts] = useState<ToastMessage[]>([]);

    // Ref to store timeout IDs for each toast (for cleanup)
    const timeoutsRef = useRef<{ [key: string]: NodeJS.Timeout }>({});
    const MAX_TOASTS = 3;

    /**
     * Dismiss a specific toast by ID with exit animation.
     * 1. Sets isRemoving flag → triggers exit animation
     * 2. Removes from DOM after 300ms (animation duration)
    */
    const dismiss = useCallback((id: string) => {
        if (timeoutsRef.current[id]) {
            clearTimeout(timeoutsRef.current[id]);
            delete timeoutsRef.current[id];
        }
        setToasts((prev) => prev.map(toast => 
            toast.id === id ? { ...toast, isRemoving: true } : toast
        ));

        // Remove from DOM after animation completes
        setTimeout(() => {
            setToasts((prev) => prev.filter((toast) => toast.id !== id));
        }, 300);
    }, []);

    /**
     * Add a new toast to the queue.
     * If max toasts exceeded, removes the oldest toast.
     * Auto-dismiss timer starts if timeout > 0.
    */
    const addToast = useCallback((title: string | undefined, description: string, variant: ToastVariant, timeout: number = 4000) => {
        // Generate unique ID using timestamp
        const id = Date.now().toString();
        const newToast: ToastMessage = {
            id,
            title,
            description,
            variant,
            timeout,
        };

        // Add toast and enforce MAX_TOASTS limit (FIFO)
        setToasts((prev) => {
            // Remove oldest toast (first in array)
            const newToasts = [...prev, newToast];
            if (newToasts.length > MAX_TOASTS) {
                const oldestId = newToasts[0].id;
                if (timeoutsRef.current[oldestId]) {
                    clearTimeout(timeoutsRef.current[oldestId]);
                    delete timeoutsRef.current[oldestId];
                }
                return newToasts.slice(1);
            }
            return newToasts;
        });

        if (timeout > 0) {
            const timer = setTimeout(() => {
                dismiss(id);
            }, timeout);
            timeoutsRef.current[id] = timer;
        }
    }, [dismiss]);

    const success = useCallback((description: string, title: string = "موفق") => {
        addToast(title, description, "success", 4000);
    }, [addToast]);

    const error = useCallback((description: string, title: string = "خطا") => {
        addToast(title, description, "error", 5000);
    }, [addToast]);

    const warning = useCallback((description: string, title: string = "هشدار") => {
        addToast(title, description, "warning", 4000);
    }, [addToast]);

    const info = useCallback((description: string, title: string = "اطلاعیه") => {
        addToast(title, description, "info", 4000);
    }, [addToast]);

    const dismissAll = useCallback(() => {
        Object.values(timeoutsRef.current).forEach(clearTimeout);
        timeoutsRef.current = {};
        setToasts([]);
    }, []);

    const getVariantStyles = (variant: ToastVariant, isRemoving?: boolean) => {
        const base = "border-r-4 p-4 rounded-lg shadow-lg mb-3 transition-all duration-300";
        const animation = isRemoving 
            ? "opacity-0 translate-x-full scale-95" 
            : "animate-in slide-in-from-top-2 fade-in";
        
        switch (variant) {
            case "success":
                return `${base} ${animation} bg-green-50 border-green-500 text-green-800`;
            case "error":
                return `${base} ${animation} bg-red-50 border-red-500 text-red-800`;
            case "warning":
                return `${base} ${animation} bg-yellow-50 border-yellow-500 text-yellow-800`;
            case "info":
                return `${base} ${animation} bg-blue-50 border-blue-500 text-blue-800`;
            default:
                return `${base} ${animation} bg-gray-50 border-gray-500 text-gray-800`;
        }
    };

    const value: ToastContextType = {
        success,
        error,
        warning,
        info,
        dismiss,
        dismissAll,
    };

    return (
        <ToastContext.Provider value={value}>
            {children}
            <div className="fixed top-20 left-1/2 -translate-x-1/2 z-50 flex flex-col w-80 max-w-full pointer-events-none">
                <div className="pointer-events-auto">
                    {toasts.map((toast) => (
                        <div
                            key={toast.id}
                            className={getVariantStyles(toast.variant, toast.isRemoving)}
                            role="alert"
                        >
                            <div className="flex items-start justify-between">
                                <div>
                                    {toast.title && (
                                        <h4 className="font-semibold text-sm mb-1">{toast.title}</h4>
                                    )}
                                    <p className="text-sm">{toast.description}</p>
                                </div>
                                <button
                                    onClick={() => dismiss(toast.id)}
                                    className="ml-4 text-gray-400 hover:text-gray-600 transition-colors text-xl leading-none"
                                >
                                    ×
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </ToastContext.Provider>
    );
}

/**
 * Custom hook to use toast methods from context.
 * Must be used within ToastProvider, otherwise throws error.
*/
export function useToast() {
    const context = useContext(ToastContext);
    if (!context) {
        throw new Error("useToast must be used within a ToastProvider");
    }
    return context;
}