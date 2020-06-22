import React, { useState, useEffect } from 'react';
import { hot } from 'react-hot-loader';

import {
    Button,
    TextField
} from '@material-ui/core';
import { styled } from '@material-ui/styles';

import {
    Selector,
    LinkTable
} from './components';
import { getConn } from './lib/websocket';

import './App.css';

const MainTextField = styled(TextField)({
    'padding-bottom': '10%'
});

function App() {
    const [link, setLink] = useState('');
    const [links, setLinks] = useState([]);
    const [selected, setSelected] = useState('Get Links');
    const [options, setOptions] = useState(['Get Links', 'Analyze']);

    const ws = getConn();
    useEffect(() => {
        ws.onmessage = (e: MessageEvent) => {
            const message = JSON.parse(e.data);
            switch (message.type) {
                case 'GET_LINK_RESULT':
                    setLinks([...links, message.linkData]);
            }
        };
    });

    function handleSubmit() {
        switch (selected) {
            case 'Get Links':
                ws.send(JSON.stringify({
                    type: 'GET_LINKS',
                    link
                }));
                break;
            case 'Analyze':
                console.log('Analyzing links');
                break;
        }
    }

    function handleOptionChange(newOption: string) {
        setSelected(newOption);
    }

    function handleTextChange(event: React.ChangeEvent<HTMLInputElement>) {
        setLink(event.target.value);
    }

    if (links.length) {
        return (
            <div style={{
                position: 'absolute', left: '50%', top: '50%',
                transform: 'translate(-50%, -50%)'
            }} className="App">
                <LinkTable items={links}></LinkTable>
            </div>
        )
    }

    return (
        <div style={{
            position: 'absolute', left: '50%', top: '50%',
            transform: 'translate(-50%, -50%)'
        }} className="App">
            <MainTextField onChange={handleTextChange} label="Link" color="primary"/>
            <br/>
            <Selector onChange={handleOptionChange} list={options} label="Options" itemIndex={0}></Selector>
            <br/>
            <Button onClick={handleSubmit}>submit</Button>
        </div>
    );
}

export default hot(module)(App);
