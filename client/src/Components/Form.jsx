import React from 'react';
import Table from "./Table";
import {Paper} from '@material-ui/core'


function rowStyleMap(row,idx){
  if(row.Broken){
    return {background:'rgb(ff, 0,0)'};
  } else {
    if (idx%2==0){
      return {background:'hsl(100, 100%, 100%)'};
    } else {
      return {background:'rgb(200,200,255)'};
    }
  }
}

export default function Form({Id,Name,AssetType,BinName,ParentBinName,ParentName,ParentAssetType,SubMovie}) {
  if((AssetType=="SubClip") || (!SubMovie) || (SubMovie.length==0) ){
    return <></>
  }
  return (
      <Paper style={{padding:10}}>
        <h3 style={{color:'blue'}}>{Name}</h3>
        {SubMovie&&AssetType=="List"&&<h4>
          محتويات القائمة
        </h4>}
        {(SubMovie&&AssetType=="Clip")&&<h4>
          الكليبات المرتبطة
        </h4>}
        {SubMovie&&<Table rowStyleMap={rowStyleMap}  style={{height:"85vh"}}  rows={SubMovie} cols={[{name:"Name",display:"الاسم"},{name:"AssetType",display:"النوع"}]}></Table>}
      </Paper>
  );
}