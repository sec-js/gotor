import React, { useState } from 'react';
import { 
    InputLabel,
    Select,
    MenuItem
} from '@material-ui/core';

type SelectorProps = {
    list: string[];
    label: string;
    itemIndex: number;
    onChange: (value: string) => void;
}
export default function Selector(props: SelectorProps) {
    const [list, setList] = useState(props.list);
    const [label, setLabel] = useState(props.label);
    const [itemIndex, setItemIndex] = useState(props.itemIndex);

    function handleChange(event: React.ChangeEvent<HTMLSelectElement>) {
        setItemIndex(parseInt(event.target.value));
        props.onChange(list[itemIndex]);
    }

    return (
        <div>
            <InputLabel id="torbotOptions">{label}</InputLabel>
            <Select onChange={handleChange} id="torbotOptions" value={itemIndex}>
                {list.map((item: string, index: number) => <MenuItem key={index} value={index}>{item}</MenuItem>)}
            </Select>
        </div>
    );
}
