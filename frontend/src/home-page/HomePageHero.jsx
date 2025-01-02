import heroImg from '../assets/hero-img.png'

export default function HomePageHero() {
    return (
        <div className="hero">
            <div className="hero-text">
                <h1>Siiick 24/7</h1>
                <p>For a pleasant shopping experience</p>
                <button>Shop Now</button>
            </div>
            <div className="hero-img">
                <img src={heroImg} />
            </div>
        </div>
    )
}

