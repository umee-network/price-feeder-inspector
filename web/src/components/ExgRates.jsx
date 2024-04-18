import { useEffect, useState } from "react"
import { USDollar, getReq, localDate } from "../js/utils"

function ExgRates() {
    const [status, setStatus] = useState(0)
    const [exgRates, setExgRates] = useState(null)
    const [currentTime, setCurrentTime] = useState(0)

    const getExgRates = (url) => {
        getReq(url + "/umee/oracle/v1/denoms/exg_rates_timestamp").then((resp) => {
            setStatus(1)
            let c = new Date()
            setCurrentTime(c.getTime())
            setExgRates(resp.exg_rates)
        })
    }

    useEffect(() => {
        const url = "https://umee-api.polkachu.com"
        if (exgRates == null) {
            getExgRates(url)
        }
        const intervalCall = setInterval(() => {
            getExgRates(url)
            // calling every 24 seconds (like mostly 3 blocks)
        }, 24000);
        return () => {
            clearInterval(intervalCall);
        };
    }, [exgRates])

    return (
        <div className="container">
            <h3 className="ubuntu-light" style={{ color: "green", marginTop: 50 }}>
                Oracle Denom Prices & Price Updated time
            </h3>

            <p className="ubuntu-light">Oracle module stores the price timestamp of each asset, if the asset price is not submitted within <b>3</b> minutes, then leverage module will not do any transactions.</p>

            <br></br>
            {status == 0 ? <div className="ubuntu-light center">Getting data...</div> :
                <div>
                    {
                        exgRates != null && exgRates.length > 0 ? <>
                            <table className="table ubuntu-light table-hover">
                                <thead>
                                    <tr>
                                        <th scope="col">#</th>
                                        <th scope="col">Denom</th>
                                        <th scope="col"><i class="bi bi-currency-dollar"></i> Price</th>
                                        <th scope="col" className="ubuntu-light"><i class="bi bi-clock"></i> Submitted At</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {

                                        exgRates.map((exgRate, index) => {
                                            let d = new Date(exgRate.timestamp)
                                            let c = d.getTime() - currentTime
                                            if (currentTime > d.getTime()) {
                                                c = currentTime - d.getTime()
                                            }
                                            return <tr className={c > 180000 ? "table-danger" : "table-success"} key={index}
                                            >
                                                <td scope="row" >{index + 1}</td>
                                                <td>{exgRate.denom.toUpperCase()}</td>
                                                <td>{USDollar.format(exgRate.rate)}</td>
                                                <td>
                                                    {localDate(exgRate.timestamp)}
                                                </td>
                                            </tr>
                                        })
                                    }
                                </tbody>
                            </table>
                        </> : <></>
                    }</div>
            }
        </div>
    )
}

export default ExgRates