const coachNames = [
  "Moon Knight", "Lunar Guardian", "Stellar Warrior", 
  "Cosmic Rider", "Orion Sentinel", "Nova Explorer",
  "Moon Knight", "Lunar Guardian", "Stellar Warrior",
  "Galaxy Hunter", "Solar Defender"
];

export const markImg = [
  { img: "/versace.webp", link: "/products/ورساچه" },
  { img: "/hablot.webp", link: "/products/هابلوت" },
  { img: "/rolex.webp", link: "/products/رولکس" },
  { img: "/omega.webp", link: "/products/اومگا" },
  { img: "/klibs.webp", link: "/products/تل سر" },
  { img: "/charm.webp", link: "/products/دست بند چرم" },
  { img: "/avis.webp", link: "/products/جاکلیدی" },
  { img: "/walar.webp", link: "/products/ساعت" },
  { img: "/felezy.webp", link: "/products/دست بند استیل" },
  { img: "/golsar.webp", link: "/products/گل سر" },
  { img: "/gooshvare.webp", link: "/products/گوشواره" },
  { img: "/kelibs.webp", link: "/products/کلیبس" },
];

export const heroImg = [
  { img: "/chekbaner.webp", link: "/products/suggested" },
  { img: "/setbanner.webp", link: "/products/sets" },
  { img: "/watchbanner.webp", link: "/products/ساعت" },
  { img: "/summerybanner.webp", link: "/products/discount" },
];
export const markCardItems = markImg.map(item => ({
  imgSrc: item.img,
  size: "sm" as const,
  link: item.link,
}));

export const heroCardItems = heroImg.map(item => ({
  imgSrc: item.img,
  size: "xl" as const,
  link: item.link,
}));

export const discountCardItems = coachNames.map(name => ({
  imgSrc: "/das.jpg",
  imgAlt: name,
  title: "ساعت بسیار زیبای برند خوب",
  variant:"magical" as const,
  price:250000,
  oldPrice:350000,
  discountPercent:50,
  size: "md" as const,
}));

export const lgCardItems = coachNames.map(name => ({
  imgSrc: "/das.jpg",
  imgAlt: name,
  title: "ساعت بسیار زیبای برند خوب",
  variant:"magical" as const,
  price:250000,
  oldPrice:350000,
  discountPercent:50,
  size: "lg" as const,
}));