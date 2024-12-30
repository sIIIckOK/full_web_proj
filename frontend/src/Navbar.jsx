import './Navbar.css'
import { Link } from 'react-router-dom'

export default function Navbar() {
    return(
        <div className="navbar">
            <ul>
                <li><Link to="/">Siiick 24/7</Link></li>
                <li><Link to="/">Best Deals</Link></li>
                <li><Link to="/">Products</Link></li>
                <li><Link to="/">About Us</Link></li>
            </ul>

            <form action="">
                <input type="text" placeholder="  Search" name="q" />
            </form>
            <Link to="/">Cart</Link>
        </div>
    )
}

