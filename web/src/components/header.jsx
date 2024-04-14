import { Link, useLocation } from "react-router-dom"

export default function Header() {
    const getActivePath = () => {
        let path = useLocation().pathname
        return path
    }

    return (
        <>
            <nav className="root-bg navbar navbar-expand-lg">
                <div className="container ubuntu-light">
                    <a className="navbar-brand " href="#">
                        <img src="https://app.ux.xyz/favicon.png" style={{ width: 20 }}>
                        </img> Umee
                    </a>
                    <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav">
                            <li className="nav-item">
                                <Link className={getActivePath() == "/" ? "nav-link active" : "nav-link"} to="/">
                                    {getActivePath()=="/"?<i className="bi bi-house-fill"></i>:<i className="bi bi-house"></i>}  Home
                                </Link>
                            </li>
                            <li className="nav-item">
                                <Link className={getActivePath() == "/pf_stats" ? "nav-link active" : "nav-link"} to="/pf_stats">
                                {getActivePath()=="/pf_stats"?<i className="bi bi-display-fill"></i>:<i className="bi bi-display"></i>} PF Stats
                                </Link>
                            </li>
                            <li className="nav-item">
                                <Link className={getActivePath() == "/ibc" ? "nav-link active" : "nav-link"} to="/ibc">
                                    <i class="bi bi-arrow-down-up"></i> IBC
                                </Link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </>
    )
}