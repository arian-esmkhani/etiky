"use client";

import { motion } from "framer-motion";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

const muscleGroups = [
  { id: 1, title: "زیورالات مو", img: "/hireasestory.webp" },
  { id: 2, title: "ست", img: "/setcategoory.webp" },
  { id: 3, title: "ساعت", img: "/watchcategory.webp" },
  { id: 4, title: "گردن‌بند", img: "/gardanbandcategoory.webp" },
  { id: 5, title: "انگشتر", img: "/angoshtarcategoory.webp" },
  { id: 6, title: "گوشواره", img: "/gooshcategory.webp" },
  { id: 7, title: "دست‌بند", img: "/dastbandkategoory.webp" },
  { id: 8, title: "آویز", img: "/aviscategoory.webp" },
];

export function HexSection() {
  const [hoveredId, setHoveredId] = useState<number | null>(null);
  const router = useRouter();
  const [isMobile, setIsMobile] = useState(false);

  useEffect(() => {
    const checkMobile = () => {
      setIsMobile(window.innerWidth < 768);
    };
    checkMobile();
    window.addEventListener('resize', checkMobile);
    return () => window.removeEventListener('resize', checkMobile);
  }, []);

  return (
    <section className="pt-8 sm:py-8">
      <div className="container mx-auto px-4">
        <div className="text-center mb-7 border-b-2 border-red-500/10 relative">
          <div className="absolute -bottom-0.5 left-0 right-0 h-0.5 bg-linear-to-r 
          from-transparent via-[rgb(217,47,74)] to-transparent" />
          <h2 className="text-lg sm:text-xl font-bold text-gray-800">
            خرید بر اساس دسته‌بندی
          </h2>
          <p className="text-xs md:text-base text-gray-500 mt-2 mb-2">
            مرور محصولات در دسته‌بندی‌های متنوع 
          </p>
        </div>

        <div className="grid grid-cols-4 gap-1 sm:gap-4">
          {muscleGroups.map((group, index) => {
            const isHovered = hoveredId === group.id;
            return (
              <motion.div
                key={group.id}
                className="flex flex-col items-center cursor-pointer"
                initial={{ opacity: 0, scale: 0.8, y: 30 }}
                whileInView={{ opacity: 1, scale: 1, y: 0 }}
                transition={{ duration: 0.6, delay: index * 0.14 }}
                viewport={{ 
                  once: true,
                  amount: isMobile ? 0.9 : 0.1,
                  margin: isMobile ? "-200px 0px" : "0px"
                }}
                onMouseEnter={() => setHoveredId(group.id)}
                onMouseLeave={() => setHoveredId(null)}
                onClick={() => {
                  if (group.title === "ست") {
                    router.push("/products/sets");
                  } else {
                    router.push(`/products/${encodeURIComponent(group.title)}`);
                  }
                }}
              >
                <motion.div
                  className="relative w-23 h-23 sm:w-36 sm:h-36 md:w-44 md:h-44"
                  animate={{ scale: isHovered ? 1.08 : 1 }}
                  transition={{ duration: 0.3 }}
                >
                  <div
                    className="w-full h-full bg-cover bg-center"
                    style={{
                      backgroundImage: `url(${group.img})`,
                      clipPath: "polygon(50% 0%, 93% 25%, 93% 75%, 50% 100%, 7% 75%, 7% 25%)",
                    }}
                  >
                    <motion.div
                      className="w-full h-full bg-black/40"
                      animate={{ opacity: isHovered ? 0.4 : 0 }}
                      transition={{ duration: 0.3 }}
                    />
                  </div>

                  <motion.div
                    className="absolute inset-0 pointer-events-none"
                    style={{
                      clipPath: "polygon(50% 0%, 93% 25%, 93% 75%, 50% 100%, 7% 75%, 7% 25%)",
                    }}
                    animate={{
                      border: isHovered ? "2px solid rgb(217,47,74)" : "2px solid transparent",
                      boxShadow: isHovered ? "0 0 30px 5px rgba(217,47,74,0.4)" : "none",
                    }}
                    transition={{ duration: 0.3 }}
                  />
                </motion.div>

                <motion.p
                  className="mt-5 text-sm sm:text-base md:text-md text-gray-800"
                  animate={{
                    color: isHovered ? "rgb(217,47,74)" : "#374151",
                    scale: isHovered ? 1.1 : 1,
                  }}
                  transition={{ duration: 0.3 }}
                >
                  {group.title}
                </motion.p>
              </motion.div>
            );
          })}
        </div>
      </div>
    </section>
  );
}