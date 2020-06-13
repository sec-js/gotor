import React, { Component, useState } from 'react';
import { FixedSizeList } from 'react-window';
import { ListItem, ListItemText } from '@material-ui/core'; 


export default function List(props) {
    const [items, setItems] = useState(props.items);

    function renderRow({ index, style }) {
        return (
            <ListItem button style={style} key={index}>
                <ListItemText primary={`${index + 1}. ${items[index]}`}/>
            </ListItem>
        );
    }

    /*
    TODO - The list dimensions should be based off of the size of the screen rather than arbitrary numbers.
    */
    const fullRowSize = 50; 
    const reducedRowSize = 35;
    const hasManyItems = items.length > 20; 
    const rowSize =  hasManyItems ? reducedRowSize : fullRowSize;
    return (
        <div>
            <FixedSizeList height={500} width={450} itemSize={rowSize} itemCount={items.length}>
                {renderRow}
            </FixedSizeList>
        </div>
    )
}