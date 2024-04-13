import { useEffect, useState } from "react"
import { getReq } from "../js/utils"

function IBCPage() {
    let sNo = 1
    let USDollar = new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
    });

    const localDate = (inp) => {
        let ns = new Date(inp)
        return ns.toLocaleString()
    }

    const [status, setStatus] = useState(0)
    const [ibcFlows, setIBCFlows] = useState({})
    const [totalOutflow, setTotalOutflow] = useState(null)
    const [totalInflow, setTotalInflow] = useState(null)
    const [ibcQuotaParams, setIBCQuotaParams] = useState(null)
    const [quotaExpiresAt, setQuotaExpiresAt] = useState(null)

    const getIBCFlows = async (url) => {
        getReq(url + "/umee/uibc/v1/all_outflows").then((resp) => {
            let newIBCObj = {}
            resp.outflows.forEach((outflow) => {
                newIBCObj[outflow.denom] = {
                    "denom": outflow.denom,
                    "outflow": outflow.amount,
                    "symbol": outflow.symbol,
                    "inflow": 0
                }
            })

            getReq(url + "/umee/uibc/v1/outflows").then((resp) => {
                setTotalOutflow(resp.amount)
            })
            getReq(url + "/umee/uibc/v1/inflows").then((resp) => {
                resp.inflows.forEach((inflow) => {
                    if (newIBCObj[inflow.denom]) {
                        newIBCObj[inflow.denom]["inflow"] = inflow.amount
                    } else {
                        newIBCObj[inflow.denom] = {
                            "denom": inflow.denom,
                            "inflow": inflow.amount,
                            "symbol": inflow.symbol,
                            "outflow": 0
                        }
                    }
                })
                let sortedData = Object.entries(newIBCObj).map(([key, flow]) => {
                    return flow
                }).sort((a, b) => {
                    return parseInt(b.inflow) - parseInt(a.inflow);
                })
                setIBCFlows(sortedData)
                setTotalInflow(resp.sum)
            })
        })
    }
    useEffect(() => {
        const url = "https://umee-api.polkachu.com"
        if (ibcQuotaParams == null) {
            getReq(url + "/umee/uibc/v1/params").then((resp) => {
                setIBCQuotaParams(resp.params)
                setStatus(1)
            })
        }

        if (quotaExpiresAt == null) {
            getReq(url + "/umee/uibc/v1/quota_expires").then((resp) => {
                setQuotaExpiresAt(resp.end_time)
            })
        }

        if (Object.entries(ibcFlows).length == 0) {
            getIBCFlows(url)
        }
        const intervalCall = setInterval(() => {
            getIBCFlows(url)
        }, 8000);
        return () => {
            clearInterval(intervalCall);
        };
    }, [ibcFlows])

    return (
        <div className="container">
            <h3 className="ubuntu-bold" style={{ color: "green", marginTop: 50 }}>
                IBC Inflows & Outflow on Umee App
            </h3>

            {status == 0 ? <div className="ubuntu-light center">Getting data...</div> :
                <div>
                    {/* <p className="ubuntu-light">Total outflows from umee chain is <b>{USDollar.format(parseInt(totalOutflow))}</b></p> */}
                    {/* <p className="ubuntu-light">Total inflows to umee chain is <b>{USDollar.format(parseInt(totalInflow))}</b></p> */}
                    {/* <p className="ubuntu-light">Total outflows from umee chain is {JSON.stringify(ibcQuotaParams)}</p> */}
                    <p className="ubuntu-light">We have total ibc outflow limit <b>{USDollar.format(parseInt(ibcQuotaParams.total_quota))}</b> and each token have <b>{USDollar.format(parseInt(ibcQuotaParams.token_quota))}</b> on IBC Rate Limits and every <b>{parseInt(ibcQuotaParams.quota_duration) / 3600}</b> hours IBC Quota will reset. Currently we don't have any IBC inflow limitations but ibc inflow will effect the ibc outflows quota.</p>
                    <p className="ubuntu-light">
                        <i className="bi bi-stopwatch"></i> IBC Quota will reset at : {localDate(quotaExpiresAt)}</p>
                    {
                        Object.entries(ibcFlows).length > 0 ? <>
                            <table className="table ubuntu-light table-hover">
                                <thead>
                                    <tr>
                                        <th scope="col">#</th>
                                        <th scope="col">IBC Denom</th>
                                        <th scope="col">Symbol</th>
                                        <th scope="col" className="ubuntu-light">Inflow <b>{USDollar.format(parseInt(totalInflow))}</b></th>
                                        <th scope="col" className="ubuntu-light">Outflow <b>{USDollar.format(parseInt(totalOutflow))}</b></th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        Object.entries(ibcFlows).map(([key, flow]) => {
                                            return <tr key={key}>
                                                <td scope="row" >{sNo++}</td>
                                                <td>{flow.denom}</td>
                                                <td>{flow.symbol.toUpperCase()}</td>
                                                <td>
                                                    <i class="bi bi-arrow-left-short"></i>
                                                    {USDollar.format(parseInt(flow.inflow).toFixed(2))}
                                                </td>
                                                <td>
                                                    <i class="bi bi-arrow-right-short"></i> {USDollar.format(parseInt(flow.outflow).toFixed(2))}
                                                </td>
                                            </tr>
                                        })
                                    }
                                </tbody>
                            </table>
                        </> : <></>
                    }</div>}
        </div>
    )
}
export default IBCPage