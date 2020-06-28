import React, { useState, useEffect } from 'react';
import {
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow
} from '@material-ui/core';

type Link = { name: string, status: boolean };

export default function(props: { items: Link[] }) {
    const [rows, setRows] = useState(props.items);
    useEffect(() => setRows(props.items));

    return (
        <div>
            <TableContainer style={{width: window.innerWidth }} component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Link</TableCell>
                            <TableCell>Status</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows.map((row: Link, index: number) => (
                            <TableRow key={index}>
                                <TableCell>{row.name}</TableCell>
                                <TableCell>{row.status.toString()}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    );
}
