import { GuaranteeLine } from "./hero-section/guarantee-line";
import { HeroSlider } from "./hero-section/hero-slider";
import { MarkSlider } from "./hero-section/mark-slider";

export default function HeroSection() {
    return (
        <div>
            <MarkSlider/>
            <HeroSlider/>
            <GuaranteeLine/>
        </div>
    )
}