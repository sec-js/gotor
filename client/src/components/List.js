import React, { useState } from 'react';
import { FixedSizeList } from 'react-window';
import { ListItem, ListItemText } from '@material-ui/core'; 

/*
TODO - Needs to be fine tuned.
*/
const adjustHeight = height => height > window.innerHeight ? window.innerHeight : height;
const calculateDimensions = items => ({
    rowHeight: items.length > 20 ? 35 : 50,
    totalHeight: adjustHeight(items.length * 50),
    totalWidth: window.innerWidth * .8
});

export default function List(props) {
    const [items, setItems] = useState(props.items);

    function renderRow({ index, style }) {
        return (
            <ListItem button style={style} key={index}>
                <ListItemText primary={`${index + 1}. ${items[index]}`}/>
            </ListItem>
        );
    }

    const { totalHeight, totalWidth, rowHeight } = calculateDimensions(items);
    return (
        <div>
            <FixedSizeList height={totalHeight} width={totalWidth} itemSize={rowHeight} itemCount={items.length}>
                {renderRow}
            </FixedSizeList>
        </div>
    )
}