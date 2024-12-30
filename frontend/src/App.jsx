import './App.css'

import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

import Navbar from './Navbar'
import HomePage from './HomePage'
import Footer from './Footer.jsx'  

function App() {
    return (
        <>
            <Router>
                <Navbar />
                <Routes>
                    <Route path='/' element={<HomePage />}></Route>
                </Routes>
                <Footer />
            </Router>
        </>
    )
}

export default App
