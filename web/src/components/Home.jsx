import { Link } from "react-router-dom"


function HomePage() {
    return (
        <div className="container">
            <h3 className="ubuntu-light text-center ux-color" style={{ marginTop: 50 }}>
                Umee App Utils Store 
            </h3>
            <br></br>
            <div className="row ubuntu-light">
                <div className="col-4">
                    <div className="card">
                        <div className="card-body">
                            <Link to="/pf_stats" className="card-text text-black no-anchor">
                                <i class="bi bi-display"></i> Price Feeder Stats
                            </Link>
                        </div>
                    </div>
                </div>
                <div className="col-4">
                    <div className="card">
                        <div className="card-body">
                            <Link to="/ibc" className="card-text text-black no-anchor">
                            <i class="bi bi-arrow-down-up"></i>  IBC Inflows & Outflows
                            </Link>
                        </div>
                    </div>
                </div>
                {/* <div className="col-4">
                    <div className="card">
                        <div className="card-body">
                            This is some text within a card body.
                        </div>
                    </div>
                </div> */}
            </div>
        </div>
    )
}

export default HomePage 