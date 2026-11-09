"use client";

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { Home, Info, Dumbbell, User, ShoppingCart } from "lucide-react";
import { HeaderBanner } from "./header-banner";
import SearchLabel from "../ui/search-lable";
import Image from "next/image";
import { useAuthStatus } from '@/app/hooks/use-auth-status';
import { useState, useCallback, useEffect } from "react";
import { useFinder } from '../hooks/use-finder';

const Routes = [
  { path: '/', label: 'خانه', icon: Home },
  { path: '/page/about', label: 'درباره ی ما', icon: Info },
  { path: '/best-selling', label: 'پرفروش ترین ها', icon: Dumbbell },  
  { path: '/discount', label: 'تخفیفات شگفت انگیز', icon: User },
  { path: '/suggested', label: 'پیشنهاد ما', icon: User },
  { path: '/sets', label: 'اکسسوری ست', icon: User },
];

const mobileRoutes = [
  { path: '/', label: 'خانه', icon: Home },
  { path: '/page/about', label: 'درباره ی ما', icon: Info },
  { path: '/cart', label: 'سبد خرید', icon: ShoppingCart },
  { path: '/discount', label: 'تخفیفات ویژه', icon: User },
  { path: '/best-selling', label: 'پرفروش ترین ها', icon: Dumbbell },
];

export function AppNav() {
  const pathname = usePathname();
  const router = useRouter();
  const { isAuthenticated, loading } = useAuthStatus();
  const [searchQuery, setSearchQuery] = useState("");
  const [suggestions, setSuggestions] = useState<string[]>([]);
  const [isLoadingSuggestions, setIsLoadingSuggestions] = useState(false);
  const { searchSuggest } = useFinder();

  const handleSearch = useCallback(() => {
    if (searchQuery.trim()) {
      setSuggestions([]);
      router.push(`/search?q=${encodeURIComponent(searchQuery.trim())}`);
    }
  }, [searchQuery, router]);

  const handleSuggestionClick = useCallback((suggestion: string) => {
    setSearchQuery(suggestion);
    router.push(`/search?q=${encodeURIComponent(suggestion)}`);
  }, [router]);

  useEffect(() => {
    const fetchSuggestions = async () => {
      if (searchQuery.trim().length > 0) {
        setIsLoadingSuggestions(true);
        try {
          const result = await searchSuggest(searchQuery.trim());
          if (result && Array.isArray(result)) {
            setSuggestions(result);
          } else {
            setSuggestions([]);
          }
        } catch (error) {
          console.error('Error fetching suggestions:', error);
          setSuggestions([]);
        } finally {
          setIsLoadingSuggestions(false);
        }
      } else {
        setSuggestions([]);
      }
    };

    const debounceTimer = setTimeout(fetchSuggestions, 300);
    return () => clearTimeout(debounceTimer);
  }, [searchQuery, searchSuggest]);

  return (
    <>
      <header className="sticky top-0 z-50 bg-red-100/20 backdrop-blur-lg">
        <HeaderBanner />
        
        <div className="w-full">
          <div className="mx-[0.5em] sm:mx-[1em] xl:mx-[3em] px-[0.1rem] sm:px-4 py-0.5 sm:py-2 
          flex flex-wrap justify-between">
            <div className="flex items-center gap-3 text-black">
              <Link href="/products/cart" className="hover:text-rose-800 transition-colors
              hidden md:block">
                <ShoppingCart size={20}/>
              </Link>
              {!loading && !isAuthenticated && (
                  <>
                      <span className="text-gray-600 text-2xl hidden md:block">|</span>
                      <Link href="/login" className="hover:text-rose-800 transition-colors
                      font-medium hidden md:block">
                          ثبت‌نام <span className="text-[0.7rem]">|</span> ورود
                      </Link>
                  </>
              )}
            </div>
            <div className="flex items-center gap-3 md:gap-3">
              <SearchLabel
                  placeholder="...جستجو در ای تی کی"
                  className="mt-3"
                  value={searchQuery}
                  onChange={(value) => setSearchQuery(value)}
                  onSearchClick={handleSearch}
                  onEnterPress={handleSearch}
                  suggestions={suggestions}
                  onSuggestionClick={handleSuggestionClick}
                  isLoading={isLoadingSuggestions}
              />            
              <Link href="/" className="md:hidden">
                <Image 
                  src="/etiky.svg" 
                  alt="etiky" 
                  width={63} 
                  height={30} 
                  priority
                />
              </Link>
              <Link href="/" className="hidden md:block">
                <Image 
                  src="/etiky.svg" 
                  alt="etiky" 
                  width={113} 
                  height={30} 
                  priority
                />
              </Link>
            </div>
          </div>
        </div>
        
        <nav className="hidden md:block w-full border-b border-gray-800/50">
          <div className="mx-[2.5em] xl:mx-[4.4em] py-1">
            <div className="flex items-center justify-end">
              <div className="flex items-center gap-7 xl:gap-11"> 
                {Routes.map((route) => {
                  const isActive = pathname === route.path;
                  return (
                    <Link
                      key={route.path}
                      href={route.path}
                      className={`text-sm font-medium transition-colors duration-200 relative group
                                ${isActive ? 'text-gray-600' : 'text-gray-600/78 hover:text-rose-800'}`}
                    >
                      {route.label}
                      <span className={`absolute -bottom-1 left-0 h-0.5 bg-rose-800/88
                                    transition-all duration-300 ${isActive ? 'w-full' : 'w-0 group-hover:w-full'}`}></span>
                    </Link>
                  );
                })} 
              </div>
            </div>
          </div>
        </nav>
      </header>

      <div className="md:hidden fixed bottom-0 left-0 right-0 bg-white shadow-lg
      border-t border-gray-200 z-50 pb-safe">
        <div className="flex justify-around items-center py-2">
          {mobileRoutes.map((route) => {
            const Icon = route.icon;
            const isActive = pathname === route.path;
            
            return (
              <Link
                key={route.path}
                href={route.path}
                className="flex flex-col items-center gap-0.5 transition-colors duration-200"
              >
                <Icon 
                  size={24}
                  className={isActive ? 'text-gray-900' : 'text-gray-500'} 
                  strokeWidth={isActive ? 2.5 : 1.8}
                />
                <span className={`text-[0.5rem] font-medium transition-colors duration-200 ${isActive ? 'text-gray-900' : 'text-gray-500'}`}>
                  {route.label}
                </span>
              </Link>
            );
          })}
        </div>
      </div>
      
      <div className="md:hidden h-8"></div>
    </>
  );
}