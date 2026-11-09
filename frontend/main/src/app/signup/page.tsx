'use client'

import Link from "next/link";
import { AppButton } from "../ui/button";
import { AppInput } from "../ui/inpute-lable";
import { useAuth } from "@/app/hooks/use-auth";
import { useState } from "react";
import Image from "next/image";
import { LoginDto } from "@/app/types/data-type";
import { useToast } from "../ui/toast";
import { useRouter } from "next/navigation";

export default function Signup() {
    const toast = useToast();
    const router = useRouter();
    const [formData, setFormData] = useState<LoginDto>({
        number: '',
        password: '',
    });
    const [confirmPassword, setConfirmPassword] = useState('');
    const [passwordError, setPasswordError] = useState('');

    const { signup, loading: hookLoading, error: hookError} = useAuth();

    const handleChange = (value: string, field: keyof LoginDto): void => {
        setFormData(prev => ({
            ...prev,
            [field]: value
        }));
        if (field === 'password') {
            setPasswordError('');
        }
    };

    const handleConfirmChange = (value: string): void => {
        setConfirmPassword(value);
        setPasswordError('');
    };

    const handleSubmit = async (): Promise<void> => {
        if (!formData.number || !formData.password) {
            toast.warning('لطفاً همه فیلدها را پر کنید');
            return;
        }
        if (formData.password !== confirmPassword) {
            setPasswordError('رمز عبور و تکرار آن مطابقت ندارند');
            return;
        }
        try {
            const result = await signup(formData.number, formData.password);
            if (result) {
                toast.success('با موفقیت ثبت نام شدید');
                router.push('/login');
            }
        } catch {
            toast.error(hookError || "")
        }
    };
    
    return(
        <div className="mx-auto w-[36vw] h-144
        flex justify-center  rounded-4xl mt-20 shadow-xl/30
        bg-[rgba(240,221,224,0.19)]">
            <div className="p-3 flex flex-col">
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
                    <div className="pt-9">
                        <span className="text-sm/5 font-serif font-light p-1">تکرار رمز عبور</span>
                        <AppInput
                        onChange={handleConfirmChange}
                        value={confirmPassword}
                        type="password"/>
                        {passwordError && <p className="text-red-500 text-xs mt-1">{passwordError}</p>}
                    </div>
                </div>
                <AppButton className="mx-1" size="lg" onClick={handleSubmit} disabled={hookLoading}>
                    {hookLoading ? 'در حال ثبت نام...' : 'ثبت نام'}
                </AppButton>                
                <span className="pt-8 pl-3">حساب کاربری دارید؟<Link href="/login" className="text-cyan-900">  ورود</Link></span>
            </div>
        </div>
    )
}