'use client'

import { useState } from "react";
import { AppButton } from "../ui/button";
import { AppInput } from "../ui/inpute-lable";
import { LoginDto } from "../types/data-type";
import Image from "next/image";
import { useAuth } from "../hooks/use-auth";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { useToast } from "../ui/toast";

export default function Login() {
    const toast = useToast();
    const router = useRouter();
    const searchParams = useSearchParams();
    const [formData, setFormData] = useState<LoginDto>({
        number: '',
        password: '',
    });

    const { login, loading: hookLoading, error: hookError } = useAuth();

    const from = searchParams.get('from') || '/';

    const handleChange = (value: string, field: keyof LoginDto): void => {
        setFormData(prev => ({
            ...prev,
            [field]: value
        }));
    };

    const handleSubmit = async (): Promise<void> => {
        if (!formData.number || !formData.password) {
            toast.warning('لطفاً همه فیلدها را پر کنید');
            return;
        }
        try {
            const result = await login(formData.number, formData.password);
            if (result) {
                toast.success('خوش آمدید!');
                router.push(from || '/');
            }
        } catch {
            toast.error(hookError || "مشکل در ورود");
        }
    };

    return(
        <div className="mx-auto w-[36vw] h-144
        flex justify-center p-3 rounded-4xl mt-20 shadow-xl/30
        bg-[rgba(240,221,224,0.19)]">
            <div className="p-14 flex flex-col">
                <div className="ml-12 sm:ml-14 text-black/90">
                    <Image 
                        src="/etiky.svg" 
                        alt="etiky" 
                        width={113} 
                        height={30} 
                        priority
                    />
                </div>
                <div className="py-16">
                    <div className="pt-6">
                        <span className="text-sm/5 font-serif font-light p-1">شماره موبایل</span>
                        <AppInput
                        onChange={(value) => handleChange(value, 'number')}
                        value={formData.number}/>
                    </div>
                    <div className="pt-9">
                        <span className="text-sm/5 font-serif font-light p-1">رمز عبور</span>
                        <AppInput
                        onChange={(value) => handleChange(value, 'password')}
                        value={formData.password}
                        type="password"/>
                    </div>
                </div>
                <AppButton className="mx-1" size="lg" onClick={handleSubmit} disabled={hookLoading}>
                    {hookLoading ? 'در حال ورود...' : 'ورود'}
                </AppButton>                
                <span className="pt-8 pl-3">حساب کاربری ندارید؟<Link href="/signup" className="text-cyan-900">  ثبت نام</Link></span>
            </div>
        </div>
    )
}