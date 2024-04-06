import { useEffect, useState } from "react";
import { getReq } from "../js/utils";

const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14]

export default function Votes() {
	const [vStatus, setvStatus] = useState(0);
	const [validators, setValidators] = useState([]);
	const [missCounters, setMissCounters] = useState({});
	let sNo = 1
	const url = "https://umee-api.polkachu.com"
	const getPFMissCounters = async (url) => {
		getReq(url + "/umee/oracle/v1/miss_counters").then((resp) => {
			const newMissCounters = JSON.parse(JSON.stringify(missCounters))
			let sortedMissCounters = resp.miss_counters.sort((a, b) => {
				return parseInt(b.miss_counter) - parseInt(a.miss_counter);
			});
			sortedMissCounters.forEach((missCounter) => {
				let mc = newMissCounters[missCounter.validator]
				if (newMissCounters.hasOwnProperty(missCounter.validator) && mc.hasOwnProperty('miss_count')) {
					if (mc['misses'].length >= 14) {
						mc['misses'] = []
					}
					if (mc['miss_count'] < missCounter.miss_counter) {
						mc['misses'].push({ "count": missCounter.miss_counter, "status": true })
						mc['miss_status'] = true
					} else {
						mc['misses'].push({ "count": missCounter.miss_counter, "status": false })
						mc['miss_status'] = false
					}
					mc['miss_count'] = missCounter.miss_counter
					newMissCounters[missCounter.validator] = mc
				} else {
					let newObj = {
						'miss_count': missCounter.miss_counter,
						'misses': [],
						'miss_status': false,
					}
					newMissCounters[missCounter.validator] = newObj
				}
				setMissCounters(newMissCounters)
			})
		})
	}

	const getValidatorName = (valoperAddr) => {
		for (let i = 0; i < validators.length; i++) {
			if (validators[i].operator_address == valoperAddr) {
				return validators[i].description.moniker
			}
		}
	}

	const getValidator = (valoperAddr) => {
		for (let i = 0; i < validators.length; i++) {
			if (validators[i].operator_address == valoperAddr) {
				return validators[i]
			}
		}
		return null
	}

	const textMiss = (missObj) => {
		if (missObj.status) {
			return "😔 miss counter : " + missObj.count
		} else {
			return "✅ miss counter : " + missObj.count
		}
	}

	useEffect(() => {
		// const url = "https://canon-4.api.network.umee.cc"
		if (validators.length == 0) {
			getReq(
				url + "/cosmos/staking/v1beta1/validators?status=BOND_STATUS_BONDED"
			).then((resp) => {
				setValidators(resp.validators);
				setvStatus(1)
			});
		}
		if (Object.entries(missCounters).length == 0) {
			getPFMissCounters(url)
		}
		const intervalCall = setInterval(() => {
			getPFMissCounters(url)
		}, 5000);
		return () => {
			clearInterval(intervalCall);
		};
	}, [missCounters]);

	return (
		<>
			{vStatus == 0 ? (
				<p style={{ textAlign: "center" }}>
					<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-clockwise" viewBox="0 0 16 16">
						<path fill-rule="evenodd" d="M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2z" />
						<path d="M8 4.466V.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L8.41 4.658A.25.25 0 0 1 8 4.466" />
					</svg> Getting validators data... </p>
			) : (
				<div className="container">
					<br></br>
					{/* <div className="row">
					<div className="col">
							<div className="card white-color-section">
								<div className="card-body">
									<h5 className="card-title">Active Validators</h5>
									<p className="card-text">{validators.length}</p>
								</div>
							</div>
						</div>
						<div className="col">
							<div className="card white-color-section">
								<div className="card-body">
									<h5 className="card-title">Oracle Assets</h5>
									<p className="card-text">{oracleAcceptedDenoms.length}</p>
								</div>
							</div>
						</div>
					</div>
					<br></br> */}
					<h5 className="ubuntu-light"><b>Oracle votes submitted by validators price-feeder</b></h5>
					<div className="row white-color-section card-radius">
						{Object.entries(missCounters).map(([key, missCounter]) => {
							return (getValidator(key) ?
								<div className="col-4" key={key}>
									<div class="row">
										<div class="col-8"><label className="truncate text-start">
											<span
												className="ubuntu-light ml-1 text-black  text-truncate dark:text-white">
												{sNo++}. {getValidator(key).description.moniker.substr(0, 26)}
											</span>
										</label></div>
										<div className="col-4">
											<a className="no-anchor" target="_blank" href={url+"/umee/oracle/v1/validators/"+key+"/miss"}>
												<span className="ubuntu-bold text-end"
												style={{
													"color": missCounter.miss_status ? "red" : "green",
												}}>
												{missCounter.miss_count}
											</span></a>
										</div>
									</div>
									<div className="row">
										{missCounter.misses.map((miss) => {
											return (
												<>
													<div className="col-1" key={miss.count}
														style={{
															height: '40px',
															width: '2px',
															backgroundColor: miss.status ? 'red' : 'green'
														}}
														data-toggle="tooltip" data-placement="top" title={textMiss(miss)}
													>
													</div>
													<div className=" col-1" style={{
														height: '30px',
														width: '2px',
														backgroundColor: 'white',
														paddingLeft: '0px',
														paddingRight: '0px',
														padding: '0px 0px 0px 0px'
													}}></div>
												</>
											)
										})}
										{Array.apply(0, Array(14 - missCounter.misses.length)).map((val, i) => {
											return (
												<>
													<div className="col-1" key={i}
														style={{
															height: '40px',
															width: '2px',
															backgroundColor: 'grey'
														}}>
													</div>
													<div className=" col-1" style={{
														height: '30px',
														width: '2px',
														backgroundColor: 'white',
														paddingLeft: '0px',
														paddingRight: '0px',
														padding: '0px 0px 0px 0px'
													}}></div>
												</>
											)
										})}
									</div>
								</div> : <></>
							)
						})}
					</div>

					{/*<div className=" list-group">*/}
					{/*	{validators.map((validator) => {*/}
					{/*		return (*/}
					{/*			<div*/}
					{/*				key={validator.operator_address}*/}
					{/*				className=" list-group-item list-group-item-action "*/}
					{/*				aria-current=" true"*/}
					{/*			>*/}
					{/*				<div className=" d-flex w-100 justify-content-between">*/}
					{/*					<p className=" mb-1">*/}
					{/*						{validator.description.moniker} -{" "}*/}
					{/*						{validator.operator_address}*/}
					{/*					</p>*/}
					{/*					<small>*/}
					{/*						{missCounters[validator.operator_address] ? (*/}
					{/*							<span*/}
					{/*								className=" badge bg-warning rounded-pill"*/}
					{/*								style={{color: " black"}}*/}
					{/*							>*/}
					{/*      {missCounters[validator.operator_address].miss_count}*/}
					{/*    </span>*/}
					{/*						) : (*/}
					{/*							<></>*/}
					{/*						)}*/}
					{/*					</small>*/}
					{/*				</div>*/}
					{/*				{missCounters[validator.operator_address] ? (*/}
					{/*					missCounters[validator.operator_address].misses.map(*/}
					{/*						(miss, i) => {*/}
					{/*							if (miss) {*/}
					{/*								return (*/}
					{/*									<span key={i}>*/}
					{/*          <i className=" bi bi-x"></i>{" "}*/}
					{/*        </span>*/}
					{/*								);*/}
					{/*							} else {*/}
					{/*								return (*/}
					{/*									<span key={i}>*/}
					{/*          <i className=" bi bi-check-circle"></i>{" "}*/}
					{/*        </span>*/}
					{/*								);*/}
					{/*							}*/}
					{/*						}*/}
					{/*					)*/}
					{/*				) : (*/}
					{/*					<></>*/}
					{/*				)}*/}
					{/*				<div>*/}
					{/*					{missCounters[validator.operator_address] ? (*/}
					{/*						<p style={{color: " red"}}>*/}
					{/*							{missCounters[*/}
					{/*								validator.operator_address*/}
					{/*								].missing_denoms.join("
									,")}*/}
					{/*						</p>*/}
					{/*					) : (*/}
					{/*						<></>*/}
					{/*					)}*/}
					{/*				</div>*/}
					{/*			</div>*/}
					{/*		);*/}
					{/*	})}*/}
					{/*</div>*/}
				</div>
			)
			}
		</>
	)
		;
}
