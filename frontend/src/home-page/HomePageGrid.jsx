import './HomePageGrid.css'

import chickenImg from "./assets/chicken.jpg"

export default function HomePageGrid({items}) {
    // TODO(shots): use fetch() to fetch items from 
    // the server on api/get-items with "amount" in json body

    // TODO(shots): change this to a grid instead of divs
    return(
     <div className="items">
            <div className="pro-container">
                {items.map((elem, index) => (
                    <div className="pro" key={index}>
                        {console.log(elem)}
                        <img src={chickenImg} />
                        <div className="des">
                            <span>FreshCuts</span>
                            <h5>{elem.title}</h5>
                            <h4>Rs. {elem.price}</h4>
                        </div>
                        <a href="/"><i className="bi bi-cart"></i></a>
                    </div>
                ))}
            </div>
        </div>
    )
    // <section id="deals">
    //     <div className="deals">
    //         <h1>DEALS OF THE DAY</h1>
    //         <div className="pro-container">
    //             <div className="pro">
    //                 <img src={chickenImg} />
    //                 <div className="des">
    //                     <span>FreshCuts</span>
    //                     <h5>Chicken Breast</h5>
    //                     <h4>Rs. 179</h4>
    //                 </div>
    //                 <a href="#"><i className="bi bi-cart"></i></a>
    //             </div>
    //             <div className="pro">
    //                 <img src={chickenImg} />
    //                 <div className="des">
    //                     <span>FreshCuts</span>
    //                     <h5>Chicken Breast</h5>
    //                     <h4>Rs. 179</h4>
    //                 </div>
    //                 <a href="#"><i className="bi bi-cart"></i></a>
    //             </div>
    //             <div className="pro">
    //                 <img src={chickenImg} />
    //                 <div className="des">
    //                     <span>FreshCuts</span>
    //                     <h5>Chicken Breast</h5>
    //                     <h4>Rs. 179</h4>
    //                 </div>
    //                 <a href="#"><i className="bi bi-cart"></i></a>
    //             </div>
    //             <div className="pro">
    //                 <img src={chickenImg} />
    //                 <div className="des">
    //                     <span>FreshCuts</span>
    //                     <h5>Chicken Breast</h5>
    //                     <h4>Rs. 179</h4>
    //                 </div>
    //                 <a href="#"><i className="bi bi-cart"></i></a>
    //             </div>
    //             <div className="pro">
    //                 <img src={chickenImg} />
    //                 <div className="des">
    //                     <span>FreshCuts</span>
    //                     <h5>Chicken Breast</h5>
    //                     <h4>Rs. 179</h4>
    //                 </div>
    //                 <a href="#"><i className="bi bi-cart"></i></a>
    //             </div>
    //             <div className="pro">
    //                 <img src={chickenImg} />
    //                 <div className="des">
    //                     <span>FreshCuts</span>
    //                     <h5>Chicken Breast</h5>
    //                     <h4>Rs. 179</h4>
    //                 </div>
    //                 <a href="#"><i className="bi bi-cart"></i></a>
    //             </div>
    //             <div className="pro">
    //                 <img src={chickenImg} />
    //                 <div className="des">
    //                     <span>FreshCuts</span>
    //                     <h5>Chicken Breast</h5>
    //                     <h4>Rs. 179</h4>
    //                 </div>
    //                 <a href="#"><i className="bi bi-cart"></i></a>
    //             </div>
    //             <div className="pro">
    //                 <img src={chickenImg} />
    //                 <div className="des">
    //                     <span>FreshCuts</span>
    //                     <h5>Chicken Breast</h5>
    //                     <h4>Rs. 179</h4>
    //                 </div>
    //                 <a href="#"><i className="bi bi-cart"></i></a>
    //             </div>
    //         </div>
    //     </div>
    // </section>
}

