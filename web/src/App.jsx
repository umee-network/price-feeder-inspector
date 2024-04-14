import Votes from './components/Votes'
import MissBeat from './components/MissBeat'
import { useEffect, useState } from 'react'
import { getReq } from './js/utils';

function App() {
	const [oracleAcceptedDenoms, setOracleAcceptedDenoms] = useState([]);
	useEffect(() => {
		const url = "https://umee-api.polkachu.com"
		if (oracleAcceptedDenoms.length == 0) {
			getReq(url + "/umee/oracle/v1/params").then(
				(oParamsResp) => {
					const d = oParamsResp.params.accept_list.map((d) =>
						d.symbol_denom.toLocaleUpperCase()
					);
					const e = Array.from(new Set(d));
					setOracleAcceptedDenoms(e);
				}
			);
		}
	}, [])
	return (
		<div className="root-app container-fluid">
			<div className="container">
				<MissBeat denoms={oracleAcceptedDenoms}></MissBeat>
				<br></br>
				<Votes></Votes>
				<br></br>
			</div>
		</div>
	)
}

export default App
