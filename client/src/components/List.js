import React, { Component } from 'react';
import { FixedSizeList } from 'react-window';
import { ListItem, ListItemText } from '@material-ui/core'; 

class List extends Component {
    constructor(props) {
        super(props);
        this.state = {
            items: props.items
        }
        this.renderRow = this.renderRow.bind(this);
    }

    renderRow({ index, style }) {
        return (
            <ListItem button style={style} key={index}>
                <ListItemText primary={`${index + 1}. ${this.state.items[index]}`}/>
            </ListItem>
        );
    }

    render() {
        /*
        TODO - The list dimensions should be based off of the size of the screen rather than arbitrary numbers.
        */
        const fullRowSize = 50; 
        const reducedRowSize = 35;
        const hasManyItems = this.state.items.length > 20; 
        const rowSize =  hasManyItems ? reducedRowSize : fullRowSize;
        return (
            <div>
                <FixedSizeList height={500} width={450} itemSize={rowSize} itemCount={this.state.items.length}>
                    {this.renderRow}
                </FixedSizeList>
            </div>
        )
    }
}

export default List;