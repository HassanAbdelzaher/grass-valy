import React from 'react';
import Title from './Title';
export default function Table({rows,cols,title,onRowClick,style,rowStyleMap}) {
    let rr;
    rr=rows
    if(!rr ||rr.length==0){
        return (<div>NO DATA</div>)
    }
    let cc =[]
    if(cols){
        cc=cols
    } else {
        Object.keys(rows[0]).forEach((key)=>{
            cc.push({field:key,display:key})
        })
    }
    let st={...{marginLeft:5,overflowY:'auto'},...(style||{})}
    return (
        <div style={st}>
            <Title>{title}</Title>
            <table className="mdl-data-table mdl-js-data-table mdl-shadow--2dp">
                <thead>
                <tr>
                    {
                        (cc||[]).map((c,cid)=>{
                            return <td  key={cid}>{c.display||c.title||c}</td>
                        })
                    }
                </tr>
                </thead>
                <tbody>
                {(rows||[]).map((r,id)=>{
                    let rowStyle={}
                    if(typeof rowStyleMap=="function"){
                        rowStyle=rowStyleMap(r,id)
                    }
                    return (<tr style={rowStyle} key={id} onClick={()=>{
                        if(onRowClick){
                            onRowClick(r,id)
                        }
                    }}>
                        {
                            (cc||[]).map((c,cid)=>{
                                let f=c.field||c.column||c.name||c
                                let val=r[f]
                                if(typeof c.map=="function"){
                                    val=c.map(val,r)
                                }
                                let tStyle={}
                                if(typeof c.styleMap=="function"){
                                    tStyle=c.styleMap(r[f],r)
                                }
                                return <td  key={cid+""+id}>{val}</td>
                            })
                        }
                    </tr>)
                })}
                </tbody>
            </table>
        </div>
)}