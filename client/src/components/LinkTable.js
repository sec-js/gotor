import React from 'react';
import {
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow
} from '@material-ui/core';

export default function(props) {
    const [rows, setRows] = React.useState(props.items);

    return (
        <div>
            <TableContainer style={{width: window.outerWidth }} component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Link</TableCell>
                            <TableCell>Status</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows.map(row => (
                            <TableRow key={row.Link}>
                                <TableCell>{row.Link}</TableCell>
                                <TableCell>{row.Status.toString()}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    );
}
