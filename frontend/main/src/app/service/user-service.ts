import { LoginDto, AuthResponse, LoginResponse } from '../types/data-type';

export const userService = {

    async signup(number: string, password: string): Promise<AuthResponse> {
        const url = new URL(`/api/auth/signup`, window.location.origin);
        const body: LoginDto = { number, password };

        const response = await fetch(url.toString(), {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(body),
        });

        checkUserResponse(response);

        return response.json();
    },

    async adminSignup(number: string, password: string): Promise<AuthResponse> {
        const url = new URL(`/api/auth/admin/signup`, window.location.origin);
        const body: LoginDto = { number, password };

        const response = await fetch(url.toString(), {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(body),
        });

        checkUserResponse(response);

        return response.json();
    },

    async login(number: string, password: string): Promise<LoginResponse> {
        const url = new URL(`/api/auth/login`, window.location.origin);
        const body: LoginDto = { number, password };

        const response = await fetch(url.toString(), {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(body),
            credentials: 'include',
        });

        checkUserResponse(response);

        return response.json();
    },
};

// Helper function to handle user-related API errors
function checkUserResponse(response: Response): void {
    if (response.status === 204) {
        throw {
            type: 'NO_CONTENT'
        };
    }

    if (!response.ok) {
        switch (response.status) {
            case 400:
                throw {
                    type: 'VALIDATION_ERROR'
                };
            case 401:
                throw {
                    type: 'INVALID_PASSWORD'
                };
            case 409:
                throw {
                    type: 'CONFLICT_ERROR'
                };
            case 500:
                throw {
                    type: 'INTERNAL_SERVER_ERROR'
                };
            default:
                throw {
                    type: 'UNKNOWN_ERROR'
                };
        }
    }
}