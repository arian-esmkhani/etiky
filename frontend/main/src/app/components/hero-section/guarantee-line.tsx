import Link from 'next/link'

export function GuaranteeLine() {
    return (
        <div className="flex flex-wrap justify-center items-center gap-4 sm:gap-8
        md:gap-12  py-4 px-2 sm:px-4">
            <Link href="/" className="transform transition-transform hover:scale-105 active:scale-95">
                <svg 
                    xmlns="http://www.w3.org/2000/svg" 
                    viewBox="0 0 200 60" 
                    className="w-[18vh] sm:w-40 md:w-[180px] lg:w-[200px] h-auto"
                >
                    <rect width="200" height="60" rx="8" fill="#d8d8d393" />
                    
                    <g transform="translate(18, 12)">
                        <circle cx="15" cy="18" r="14" stroke="#171717" strokeWidth="2.5" fill="none"/>
                        <text x="9" y="24" fontFamily="Arial" fontWeight="bold" fontSize="18" fill="#171717">$</text>
                    </g>
                    
                    <text x="55" y="28" fontFamily="Arial, sans-serif" fontWeight="bold" fontSize="17" fill="#171717" letterSpacing="1">
                        ضمانت قیمت
                    </text>
                    <text x="55" y="44" fontFamily="Arial, sans-serif" fontWeight="normal" fontSize="11" fill="#171717" opacity="0.7">
                        بهترین قیمت بازار
                    </text>
                </svg>
            </Link>

            <Link href="/" className="transform transition-transform hover:scale-105 active:scale-95">
                <svg 
                    xmlns="http://www.w3.org/2000/svg" 
                    viewBox="0 0 200 60" 
                    className="w-[18vh] sm:w-40 md:w-[180px] lg:w-[200px] h-auto"
                >
                    <rect width="200" height="60" rx="8" fill="#d8d8d393" />
                    <g transform="translate(15, 12)">
                        <path d="M15 0 L28 5 L28 18 C28 28 15 32 15 32 C15 32 2 28 2 18 L2 5 Z" 
                            stroke="#171717" strokeWidth="2.5" fill="none" strokeLinejoin="round"/>
                        <polyline points="10,16 14,20 21,12" stroke="#171717" strokeWidth="2.5" fill="none" 
                                strokeLinecap="round" strokeLinejoin="round"/>
                    </g>
                    <text x="55" y="28" fontFamily="Arial, sans-serif" fontWeight="bold" fontSize="18" fill="#171717" letterSpacing="1">
                        سلامت کالا
                    </text>
                    <text x="55" y="44" fontFamily="Arial, sans-serif" fontWeight="normal" fontSize="11" fill="#171717" opacity="0.7">
                        تضمین سلامت و کیفیت کالا
                    </text>
                </svg>
            </Link>

            <Link href="/" className="transform transition-transform hover:scale-105 active:scale-95">
                <svg 
                    xmlns="http://www.w3.org/2000/svg" 
                    viewBox="0 0 200 60" 
                    className="w-[18vh] sm:w-40 md:w-[180px] lg:w-[200px] h-auto"
                >
                    <rect width="200" height="60" rx="8" fill="#d8d8d393" />
                    <g transform="translate(15, 15)">
                        <path d="M5 20 L5 5 L25 5 L30 12 L30 20 L25 20" 
                            stroke="#171717" strokeWidth="2.5" fill="none" strokeLinecap="round" strokeLinejoin="round"/>
                        <rect x="8" y="17" width="14" height="6" rx="2" fill="#171717" opacity="0.15"/>
                        <circle cx="11" cy="26" r="4" stroke="#171717" strokeWidth="2" fill="none"/>
                        <circle cx="24" cy="26" r="4" stroke="#171717" strokeWidth="2" fill="none"/>
                    </g>
                    <text x="55" y="28" fontFamily="Arial, sans-serif" fontWeight="bold" fontSize="18" fill="#171717" letterSpacing="1">
                        FREE
                    </text>
                    <text x="55" y="44" fontFamily="Arial, sans-serif" fontWeight="normal" fontSize="11" fill="#171717" opacity="0.7">
                    ارسال رایگان
                    </text>
                </svg>
            </Link>
        </div>
    )
}