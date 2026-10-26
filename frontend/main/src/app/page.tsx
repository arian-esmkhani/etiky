import { BestSellingSlider } from "./components/best-selling-slider";
import { CheapestSlider } from "./components/cheapest-slider";
import { DiscountSlider } from "./components/discount-slider";
import HeroSection from "./components/hero-section";
import { HexSection } from "./components/hex-section";
import { SetSlider } from "./components/set-slider";
import { SuggestedSlider } from "./components/suggested-slider";
import { TrendSlider } from "./components/trend-slider";

export default function Home() {
  return (
    <div>
      <HeroSection/>
      <DiscountSlider/>
      <HexSection/>
      <SuggestedSlider/>
      <BestSellingSlider/>
      <SetSlider/>
      <TrendSlider/>
      <CheapestSlider/>
    </div>
  )
}