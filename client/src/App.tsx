import React, { useState, useEffect } from 'react';
import { hot } from 'react-hot-loader';
import axios from 'axios';

import {
    Button,
    TextField,
    CircularProgress
} from '@material-ui/core';
import { styled } from '@material-ui/styles';

import {
    Selector,
    LinkTable
} from './components';

import './App.css';

const MainTextField = styled(TextField)({
    'padding-bottom': '10%'
});

const isURL = (url: string): boolean => {
    try {
        const u = new URL(url);
        if (!u) return false;
        if (!u.hostname) return false;
        if (!u.protocol) return false;
    } catch (err) {
        return false;
    }

    return true;
}

function App() {
    const [link, setLink] = useState('');
    const [links, setLinks] = useState([]);
    const [selected, setSelected] = useState('Get Links');
    const [options, setOptions] = useState(['Get Links', 'Analyze']);
    const [startLoad, setStartLoad] = useState(false);
    const [hasError, setHasError] = useState(false);
    const [errorText, setErrorText] =  useState('Error Found.');



    const setNewError = (message: string) => {
        setHasError(true);
        setErrorText(message);
    };

    async function handleSubmit() {
        if (!isURL(link)) {
            setNewError('URL is invalid.');
            return;
        }
        setStartLoad(true);
        switch (selected) {
            case 'Get Links':
                    try {
                        const response = await axios.get(`http://localhost:3050?link=${link}`);
                        if (Array.isArray(response.data.length) && !response.data.length) throw new Error('No links found.');
                        setLinks(response.data);
                    } catch (err) {
                        setNewError(err.message)
                        return;
                    } finally {
                        setStartLoad(false);
                    }
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
        if (hasError) setHasError(false);
        setLink(event.target.value);
    }

    function handleOnKeyDown(event: React.KeyboardEvent<HTMLInputElement>) {
        switch (event.key) {
            case 'Enter':
                return handleSubmit();
        }
    }

    switch (true) {
        case startLoad:
            return <CircularProgress style={{
                    position: 'absolute', left: '50%', top: '50%',
                    transform: 'translate(-50%, -50%)'
                }} className="App"/>;
        case links.length !== 0:
            return (
                <div style={{
                    position: 'absolute', left: '50%', top: '50%',
                    transform: 'translate(-50%, -50%)'
                }} className="App">
                    <LinkTable items={links}></LinkTable>
                </div>
            )
        default:
            return (
                <div style={{
                    position: 'absolute', left: '50%', top: '50%',
                    transform: 'translate(-50%, -50%)'
                }} className="App">
                    {hasError ?
                        <MainTextField error onKeyDown={handleOnKeyDown} onChange={handleTextChange} helperText={errorText} label="Link" color="primary"/> :
                        <MainTextField onKeyDown={handleOnKeyDown} onChange={handleTextChange} label="Link" color="primary"/>}
                    <br/>
                    <Selector onChange={handleOptionChange} list={options} label="Options" itemIndex={0}></Selector>
                    <br/>
                    <Button onClick={handleSubmit}>submit</Button>
                </div>
            );
    }
}

export default hot(module)(App);
