import React, { useState } from "react";

export default function MissBeat(props) {
    return (
        <div
            className="container"
            style={{
                margin: "auto",
            }}
        >
            <h3 className="ubuntu-light" style={{ color: "green",marginTop:50 }}>
                Umee Validators Price Feeder Stats
            </h3>
            <p className="ubuntu-light">There are totally <b>{props.denoms.length}</b> unique assets are accepted by Umee Oracle module</p>
            <h5 className="ubuntu-light"><b>Oracle Aceepted Assets</b></h5>
            <div className="ubuntu-light row white-color-section card-radius">
                {props.denoms.map((denom,index) => {
                    return (
                        <div key={denom} className="col-2">
                            <div className="card-title"><b>{index+1}. {denom}</b></div>
                        </div>
                    );
                })}
            </div>
        </div>
    );
}