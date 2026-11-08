import { useCallback, useState } from 'react';
import { userService } from '../service/user-service';
import { AuthResponse, LoginResponse } from '../types/data-type';
import { ApiError } from '../types/api-type';

export const useAuth = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const signup = useCallback(
        async (number: string, password: string): Promise<AuthResponse> => {
            setLoading(true);
            setError(null);
            try {
                return await userService.signup(number, password);
            } catch (err) {
                const error = err as ApiError;
                let errorMessage = '';
                switch (error.type) {
                    case 'CONFLICT_ERROR':
                        errorMessage = 'این شماره قبلاً ثبت نام کرده است';
                        break;
                    case 'VALIDATION_ERROR':
                        errorMessage = 'اطلاعات وارد شده معتبر نیست';
                        break;
                    case 'INTERNAL_SERVER_ERROR':
                        errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                        break;
                    default:
                        errorMessage = 'خطا در ثبت نام';
                }
                setError(errorMessage);
                throw new Error(errorMessage); 
            } finally {
                setLoading(false);
            }
        },
        []
    );

    const adminSignup = useCallback(
        async (number: string, password: string): Promise<AuthResponse> => {
            setLoading(true);
            setError(null);
            try {
                return await userService.adminSignup(number, password);
            } catch (err) {
                const error = err as ApiError;
                let errorMessage = '';
                switch (error.type) {
                    case 'CONFLICT_ERROR':
                        errorMessage = 'این شماره قبلاً ثبت نام کرده است';
                        break;
                    case 'VALIDATION_ERROR':
                        errorMessage = 'اطلاعات وارد شده معتبر نیست';
                        break;
                    case 'INTERNAL_SERVER_ERROR':
                        errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                        break;
                    default:
                        errorMessage = 'خطا در ایجاد ادمین';
                }
                setError(errorMessage);
                throw new Error(errorMessage); 
            } finally {
                setLoading(false);
            }
        },
        []
    );

    const login = useCallback(
        async (number: string, password: string): Promise<LoginResponse> => {
            setLoading(true);
            setError(null);
            try {
                return await userService.login(number, password);
            } catch (err) {
                const error = err as ApiError;
                let errorMessage = '';
                switch (error.type) {
                    case 'CONFLICT_ERROR':
                        errorMessage = 'کاربری با این شماره یافت نشد';
                        break;
                    case 'INVALID_PASSWORD':
                        errorMessage = 'رمز عبور اشتباه است';
                        break;
                    case 'VALIDATION_ERROR':
                        errorMessage = 'اطلاعات وارد شده معتبر نیست';
                        break;
                    case 'INTERNAL_SERVER_ERROR':
                        errorMessage = 'خطای سرور، لطفاً دقایقی دیگر تلاش کنید';
                        break;
                    default:
                        errorMessage = 'خطا در ورود';
                }
                setError(errorMessage);
                throw new Error(errorMessage); 
            } finally {
                setLoading(false);
            }
        },
        []
    );

    return {
        loading,
        error,
        signup,
        adminSignup,
        login,
    };
};