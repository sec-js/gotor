import axios from 'axios';

import React, { Component, useState } from 'react';
import { hot } from 'react-hot-loader';

import {
    Button,
    TextField
} from '@material-ui/core';
import { styled } from '@material-ui/styles';

import {
    List,
    Selector
}  from './components';

import './App.css';

const MainTextField = styled(TextField)({
    'padding-bottom': '10%'
});

const getLinks = (url) => {
    const urlParam = encodeURIComponent(url);
    return axios.get(`http://localhost:3050?url=${urlParam}`);
};

function App(props) {
    const [url, setUrl] = useState('');
    const [links, setLinks] = useState({});
    const [selected, setSelected] = useState('Get Links');
    const [options, setOptions] = useState(['Get Links', 'Analyze']);

    function handleSubmit() {
        switch (selected) {
            case 'Get Links':
                getLinks(url)
                    .then(({ data: links }) => setLinks(links))
                    .catch(err => console.error(err));
                break;
            case 'Analyze':
                console.log('Analyzing links');
                break;
        }
    }

    function handleOptionChange(newOption) {
        setSelected(newOption);
    }

    function handleTextChange(event) {
        setUrl(event.target.value);
    }

    const urls = Object.keys(links);
    if (urls.length) {
        return (
            <div style={{
                position: 'absolute', left: '50%', top: '50%',
                transform: 'translate(-50%, -50%)'
            }} className="App">
                <List items={urls}></List>
            </div>
        )
    }

    return (
        <div style={{
            position: 'absolute', left: '50%', top: '50%',
            transform: 'translate(-50%, -50%)'
        }} className="App">
            <MainTextField onChange={handleTextChange} label="URL" color="primary"/>
            <br/>
            <Selector onChange={handleOptionChange} list={options} label="Options" itemIndex={0}></Selector>
            <br/>
            <Button onClick={handleSubmit}>submit</Button>
        </div>
    );
}

export default hot(module)(App);