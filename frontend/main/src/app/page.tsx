import { DiscountSlider } from "./components/discount-slider";
import HeroSection from "./components/hero-section";
import { HexSection } from "./components/hex-section";

export default function Home() {
  return (
    <div>
      <HeroSection/>
      <DiscountSlider/>
      <HexSection/>
    </div>
  )
}