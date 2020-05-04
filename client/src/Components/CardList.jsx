import React from 'react';
import Title from './Title';
import { Box, Paper, Card, CardHeader, CardContent, CardActions } from '@material-ui/core'
import { Check } from '@material-ui/icons'
export default function Table({ rows, cols, title, onRowClick, style, rowStyleMap }) {
    let rr;
    rr = rows
    if (!rr || rr.length == 0) {
        return (<div>NO DATA</div>)
    }
    let cc = []
    if (cols) {
        cc = cols
    } else {
        Object.keys(rows[0]).forEach((key) => {
            cc.push({ field: key, display: key })
        })
    }
    let st = { ...{ margin: 5, padding: 10, overflowY: 'auto' }, ...(style || {}) }
    return (
        <Box style={{ ...st, background: 'gray',minWidth:'300px' }} >
            <Title>{title}</Title>
            {(rows || []).map((r, id) => {
                let rowStyle = { margin: 5, height: '60px' }
                let bg=r.Broken?"#aa,55,55":'#fff9f9';
                if (typeof rowStyleMap == "function") {
                    rowStyle = { ...rowStyle, ...(rowStyleMap(r, id)) }
                }
                return (<Card style={{ margin: 5,backgroundColor:bg,padding:10}} onClick={() => {
                    if (onRowClick) {
                        onRowClick(r, id)
                    }
                }}>

                    <CardContent>
                        <Paper style={{ padding: 5,margin:5 }}>
                            <div style={{padding:5}}>{r.Name || "no name"}</div>
                            {r.BinName && <div> {r.BinName}</div>}
                        </Paper>
                        {r.ParentName&&<Paper style={{ padding: 5,margin:5,backgroundColor:'lightgray' }}>
                        <div>Parent</div>
                            {r.ParentName && <div style={{padding:5}}>{r.ParentName}</div>}
                            {r.ParentBinName && <div style={{padding:5}}>Bin= {r.ParentBinName}</div>}
                        </Paper>}
                    </CardContent>
                    <CardActions>
                        {r.AssetType}
                    </CardActions>

                </Card>)
            })}
        </Box>
    )
}