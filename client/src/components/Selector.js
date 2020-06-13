import React, { Component } from 'react';
import { 
    InputLabel,
    Select,
    MenuItem
} from '@material-ui/core';

class Selector extends Component {
    constructor(props) {
        super(props);
        this.state = {
            list: props.list,
            label: props.label,
            itemIndex: props.itemIndex
        };
        this.handleChange = this.handleChange.bind(this);
    }

    handleChange(event) {
        this.setState({
            itemIndex: event.target.value
        });
        this.props.onChange(this.state.list[event.target.value]);
    }

    render() {
        return (
            <div>
                <InputLabel id="torbotOptions">{this.state.label}</InputLabel>
                <Select onChange={this.handleChange} id="torbotOptions" value={this.state.itemIndex}>
                    {this.state.list.map((element, index) => <MenuItem key={index} value={index}>{element}</MenuItem>)}
                </Select>
            </div>
        );
    }
}

export default Selector;