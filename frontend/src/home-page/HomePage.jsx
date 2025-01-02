import { useEffect, useState } from 'react'
import './HomePage.css'

import HomePageHero from './components/HomePageHero.jsx'
import HomePageGrid from './HomePageGrid.jsx'

export default function HomePage() {
    const [items, setItems] = useState([])
    useEffect(() => {
        fetch("api/get-items")
            .then(res => {
                if (!res.ok) throw new Error('Network error')
                return res.json()
            })
            .then(js => {
                setItems(js)
            })
            .catch(err => console.log("[ERROR]", err))
    }, [])

    return(
        <>
            <HomePageHero />
            <HomePageGrid items={items}/>
        </>
    )
}
