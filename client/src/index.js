import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { initConn } from './lib/websocket.ts';

initConn('ws://localhost:3050').then(() => ReactDOM.render(<App/> , document.getElementById("root")));
