import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import IBCPage from './components/IBC.jsx';
import NoPage from './components/NoPage.jsx';
import Layout from './components/Layout.jsx';
import HomePage from './components/Home.jsx';
import ExgRates from './components/ExgRates.jsx';

function MainApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout/>}>
          <Route index element={<HomePage />} />
          <Route path="ibc" element={ <IBCPage></IBCPage>} />
          <Route path="pf_stats" element={ <App/>} />
          <Route path="exg_rates" element={ <ExgRates/>} />
          <Route path="*" element={<NoPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}


ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <MainApp />
  </React.StrictMode>,
)
