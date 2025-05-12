  ##### TRITRUST Exchange






<div align="center">
  <br>
  <img src="https://img.shields.io/badge/hyperledger-2F3134?style=for-the-badge&logo=hyperledger&logoColor=white">
  <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white">
  <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white">
  <img src="https://img.shields.io/badge/CouchDB-EA4C89?style=for-the-badge&logo=couchdb&logoColor=white">
  <img src="https://img.shields.io/badge/Blockchain-3DDC84?style=for-the-badge&logo=blockchain-dot-com&logoColor=white">
</div>

---

## 🌟 Project Overview

<p align="center" style="font-size: 18px; color: #FF69B4;">
  <em>"A permissioned blockchain network built using Hyperledger Fabric that enables decentralized, secure, transparent, and verifiable interactions between three financial entities BankOrg, LoanproviderOrg and InsuranceOrg"</em>
</p>

<div style="background: linear-gradient(145deg, #1a1a1a, #2a2a2a); padding: 20px; border-radius: 15px; box-shadow: 0 4px 8px rgba(255,0,127,0.2);">
  
  <h3 style="color: #FF007F; border-bottom: 2px dashed #FF007F; padding-bottom: 8px;">🚨 Problem Statement</h3>
  
  Traditional Financial systems suffer from:
  <ul style="color: #CCCCCC;">
    <li>❌ Centralized control points</li>
    <li>❌ Vulnerability to data tampering</li>
    <li>❌ Lack of transparent audit trails</li>
  </ul>
  
  <h3 style="color: #00FF7F; border-bottom: 2px dashed #00FF7F; padding-bottom: 8px; margin-top: 25px;">💡 Our Solution</h3>
  
  <ul style="color: #CCCCCC;">
    <li>✅ <span style="color: #FFD700;">Decentralized</span> authority using Hyperledger Fabric</li>
    <li>✅ <span style="color: #00BFFF;">Immutable</span> record-keeping with blockchain</li>
    <li>✅ <span style="color: #FF69B4;">Real-time</span> event tracking and notifications</li>
  </ul>

</div>

---

## ✨ Key Features

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 20px; margin-top: 30px;">

<div style="background: #1a1a1a; padding: 20px; border-radius: 12px; border-left: 4px solid #FF007F;">
  <h3 style="color: #FF007F;">🔐 Role-Based Access</h3>
  <p style="color: #CCCCCC;">Granular permissions through MSP IDs:</p>
  <ul style="color: #AAAAAA;">
    <li>BankOrg: Account Creation</li>
    <li>LoanproviderOrg: Verifies and Grants Loan applications</li>
    <li>InsuranceOrg: Verifies and Adds Insurance to existing Account</li>
  </ul>
</div>
</div>


---
# System Requirements
- Docker 28.0.4
- Go 1.23


Installation

# Clone Repository
git clone https://github.com/AnusreeSreekumar/tritrust_exchange.git


# Start Network
cd minifabNetwork
./startNetwork.sh

# Deploy Chaincode
minifab ccup -n Tritrustchannel -l go -v 1.0 -d false -r false


<div align="center" style="margin-top: 40px;"> <br> <p style="color: #888;">Made By Developer Anusree Sreekumar using Hyperledger Fabric</p> <div style="display: flex; justify-content: center; gap: 15px; margin-top: 20px;">  </div> </div>


