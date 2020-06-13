import React, { useState } from 'react';
import { 
    InputLabel,
    Select,
    MenuItem
} from '@material-ui/core';

export default function Selector(props) {
    const [list, setList] = useState(props.list);
    const [label, setLabel] = useState(props.label);
    const [itemIndex, setItemIndex] = useState(props.itemIndex);

    function handleChange(event) {
        setItemIndex({
            itemIndex: event.target.value
        });
        props.onChange(list[itemIndex]);
    }

    return (
        <div>
            <InputLabel id="torbotOptions">{label}</InputLabel>
            <Select onChange={handleChange} id="torbotOptions" value={itemIndex}>
                {list.map((element, index) => <MenuItem key={index} value={index}>{element}</MenuItem>)}
            </Select>
        </div>
    );
}
