import axios from 'axios';

import React, { Component } from 'react';
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

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            url: '',
            links: {},
            selected: 'Get Links', 
            options: ['Get Links', 'Analyze']
        };
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleTextChange = this.handleTextChange.bind(this);
        this.handleOptionChange = this.handleOptionChange.bind(this);
    }

    handleSubmit() {
        switch (this.state.selected) {
            case 'Get Links':
                getLinks(this.state.url)
                    .then(({ data: links }) => this.setState({links}))
                    .catch(err => console.error(err));
                break;
            case 'Analyze':
                console.log('Analyzing links');
                break;
        }
    }

    handleOptionChange(newOption) {
        this.setState({
            selected: newOption
        });
    }

    handleTextChange(event) {
        this.setState({
            url: event.target.value
        });
    }

    render() {
        const urls = Object.keys(this.state.links);
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
                <MainTextField onChange={this.handleTextChange} label="URL" color="primary"/>
                <br/>
                <Selector onChange={this.handleOptionChange} list={this.state.options} label="Options" itemIndex={0}></Selector>
                <br/>
                <Button onClick={this.handleSubmit}>submit</Button>
            </div>
        );
    }
}

export default hot(module)(App);