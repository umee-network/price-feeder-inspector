import { useEffect, useState } from "react";

const getReq = async (url) => {
  return await fetch((url = url))
    .then((resp) => resp.json())
    .then((resp) => {
      return resp;
    })
    .catch((err) => {
      console.log(err);
    });
};

const mcObj = JSON.parse(
  JSON.stringify({
    missing: false,
    missing_denoms: [],
    miss_count: 0,
    misses: [],
  })
);

export default function Votes() {
  const [vStatus, setvStatus] = useState(0);
  const [validators, setValidators] = useState([]);
  const [oracleAcceptedDenoms, setOracleAcceptedDenoms] = useState([]);
  const [missCounters, setMissCounters] = useState({});

  const getMSValue = async (valaddr) => {
    await getReq(
      "https://canon-4.api.network.umee.cc/umee/oracle/v1/validators/" +
        valaddr +
        "/miss"
    ).then((e) => {
      if (missCounters.hasOwnProperty(valaddr)) {
        let mc = missCounters[valaddr];
        if (mc["misses"].length > 30) {
          mc["misses"] = [];
        }
        if (e.miss_counter > mc["miss_count"]) {
          mc["miss_count"] = e.miss_counter;
          mc["missing"] = true;
          mc["misses"].push(true);
        } else {
          mc["missing"] = false;
          mc["misses"].push(false);
        }
        missCounters[valaddr] = mc;
        setMissCounters(missCounters);
      }
    });
  };

  const aggregateVotes = () => {
    getReq(
      "https://canon-4.api.network.umee.cc/umee/oracle/v1/validators/aggregate_votes"
    ).then((resp) => {
      const votes = resp.aggregate_votes;
      if (resp.aggregate_votes.length == 0) {
        setTimeout(() => {
          aggregateVotes();
        }, 2500);
      }
      votes.forEach((vote) => {
        if (missCounters[vote.voter]) {
          let mc = missCounters[vote.voter];
          mc["missing_denoms"] = [];
          console.log(
            oracleAcceptedDenoms.length,
            vote.exchange_rate_tuples.length
          );
          if (oracleAcceptedDenoms.length != vote.exchange_rate_tuples.length) {
            const votedDenoms = vote.exchange_rate_tuples.map((e) => e.denom);
            let d = [];
            oracleAcceptedDenoms.forEach((denom) => {
              if (!votedDenoms.includes(denom)) {
                console.log("missing denom ", denom);
                d.push(denom);
              }
            });
            mc["missing"] = true;
            mc["missing_denoms"] = d;
            missCounters[vote.voter] = mc;
            setMissCounters(missCounters);
            getMSValue(vote.voter);
            console.log("missing_denoms ", vote.voter, mc);
          } else {
            mc["missing"] = false;
          }
          missCounters[vote.voter] = mc;
          setMissCounters(missCounters);
          // getMSValue(vote.voter)
        } else {
          missCounters[vote.voter] = {
            missing: false,
            missing_denoms: [],
            miss_count: 0,
            misses: [],
          };
          setMissCounters(missCounters);
        }
      });
    });
  };

  useEffect(() => {
    getReq(
      "https://canon-4.api.network.umee.cc/cosmos/staking/v1beta1/validators?status=BOND_STATUS_BONDED"
    ).then((resp) => {
      setValidators(resp.validators);
    });
    getReq("https://canon-4.api.network.umee.cc/umee/oracle/v1/params").then(
      (oParamsResp) => {
        const d = oParamsResp.params.accept_list.map((d) =>
          d.symbol_denom.toLocaleUpperCase()
        );
        const e = Array.from(new Set(d));
        setOracleAcceptedDenoms(e);
        setvStatus(1);
        aggregateVotes();
      }
    );
    const intervalCall = setInterval(() => {
      aggregateVotes();
    }, 10000);
    return () => {
      clearInterval(intervalCall);
    };
  }, []);

  return (
    <>
      {vStatus == 0 ? (
        <p style={{ textAlign: "center" }}> Getting validators... </p>
      ) : (
        <div>
          <div className="row">
            <div className="col">
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">Total Bonded Validator</h5>
                  <p className="card-text">{validators.length}</p>
                  {/* <a href="#" className="btn btn-primary">Go somewhere</a> */}
                </div>
              </div>
            </div>
            <div className="col">
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">Oracle Assets</h5>
                  <p className="card-text">{oracleAcceptedDenoms.length}</p>
                  {/* <a href="#" className="btn btn-primary">Go somewhere</a> */}
                </div>
              </div>
            </div>
          </div>
          <br></br>

          <h5>Oracle Assets</h5>
          <div className="row">
            {oracleAcceptedDenoms.map((denom) => {
              return (
                <div key={denom} className="col-1">
                  <div className="card-title">{denom}</div>
                </div>
              );
            })}
          </div>
          <h5>Oracle votes</h5>
          <div className="list-group">
            {validators.map((validator) => {
              return (
                <div
                  key={validator.operator_address}
                  className="list-group-item list-group-item-action "
                  aria-current="true"
                >
                  <div className="d-flex w-100 justify-content-between">
                    <p className="mb-1">
                      {validator.description.moniker} -{" "}
                      {validator.operator_address}
                    </p>
                    <small>
                      {missCounters[validator.operator_address] ? (
                        <span
                          className="badge bg-warning rounded-pill"
                          style={{ color: "black" }}
                        >
                          {missCounters[validator.operator_address].miss_count}
                        </span>
                      ) : (
                        <></>
                      )}
                    </small>
                  </div>
                  {missCounters[validator.operator_address] ? (
                    missCounters[validator.operator_address].misses.map(
                      (miss, i) => {
                        if (miss) {
                          return (
                            <span key={i}>
                              <i className="bi bi-x"></i>{" "}
                            </span>
                          );
                        } else {
                          return (
                            <span key={i}>
                              <i className="bi bi-check-circle"></i>{" "}
                            </span>
                          );
                        }
                      }
                    )
                  ) : (
                    <></>
                  )}
                  <div>
                    {missCounters[validator.operator_address] ? (
                      <p style={{ color: "red" }}>
                        {missCounters[
                          validator.operator_address
                        ].missing_denoms.join(",")}
                      </p>
                    ) : (
                      <></>
                    )}
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      )}
    </>
  );
}
