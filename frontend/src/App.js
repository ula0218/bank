import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { BrowserRouter as Router, Route, Routes, useNavigate, useParams } from 'react-router-dom';
import './App.css'; // 用於自定義樣式

function BankSelector({ onBankCodeChange }) {
    const [bankCodes, setBankCodes] = useState([]);
    const [selectedBankCode, setSelectedBankCode] = useState('');

    useEffect(() => {
        axios.get('http://localhost:80/')
            .then(response => {
                const sortedBankCodes = response.data.sort((a, b) => a.code.localeCompare(b.code));
                setBankCodes(sortedBankCodes);
            })
            .catch(error => {
                console.error('Error fetching bank codes:', error);
            });
    }, []);

    const handleBankCodeChange = (event) => {
        const code = event.target.value;
        setSelectedBankCode(code);
        onBankCodeChange(code);
    };

    return (
        <div className="bank-selector">
            <h1>銀行查詢功能</h1>
            <label htmlFor="bankCode">銀行代碼查詢:</label>
            <select id="bankCode" onChange={handleBankCodeChange} value={selectedBankCode}>
                <option value="">-- 銀行代碼 --</option>
                {bankCodes.map((bank, index) => (
                    <option key={index} value={bank.code}>{`${bank.code} ${bank.name}`}</option>
                ))}
            </select>
        </div>
    );
}

function BranchList() {
    const [branches, setBranches] = useState([]);
    const { bankCode } = useParams();

    useEffect(() => {
        if (bankCode) {
            axios.get(`http://localhost:80/${bankCode}/branches`)
                .then(response => {
                    setBranches(response.data);
                })
                .catch(error => {
                    console.error(`Error fetching branches for bank ${bankCode}:`, error);
                    setBranches([]);
                });
        }
    }, [bankCode]);

    return (
        <div className="branch-card">
            {branches.map((branch, index) => (
                <div key={index}>
                    <strong>分行代碼:</strong> {branch.branch_code}<br />
                    <strong>地址:</strong> {branch.address}<br />
                    <strong>電話:</strong> {branch.phone}<br />
                    <hr />
                </div>
            ))}
        </div>
    );
}

function App() {
    const navigate = useNavigate();

    const handleBankCodeChange = (code) => {
        navigate(`/${code}`);
    };

    return (
        <div className="App">
            <BankSelector onBankCodeChange={handleBankCodeChange} />
            <Routes>
                <Route path="/:bankCode" element={<BranchList />} />
            </Routes>
        </div>
    );
}

export default function MainApp() {
    return (
        <Router>
            <App />
        </Router>
    );
}
