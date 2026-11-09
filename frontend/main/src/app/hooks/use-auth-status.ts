import { useEffect, useState } from 'react';

export const useAuthStatus = () => {
    const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const response = await fetch('/api/auth/check', {
                    credentials: 'include',
                    headers: {
                        'Cache-Control': 'no-cache',
                    },
                });
                
                setIsAuthenticated(response.ok);
            } catch (error) {
                console.error('Auth check failed:', error);
                setIsAuthenticated(false);
            } finally {
                setLoading(false);
            }
        };
        
        checkAuth();
    }, []);

    return { isAuthenticated, loading };
};