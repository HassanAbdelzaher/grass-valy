import React from 'react';
import Table from "./Table";


function styleMap(val,row){
  if(row.Broken){
    return undefined
  }
  return     (val=="List"?{background:'hsl(39, 100%, 50%)'}:(val=="SubClip"?{background:'hsl(248, 53%, 58%)'}:{background:'hsl(147, 50%, 47%)'}))
}

function rowStyleMap(row,idx){
  if(row.Broken){
    if (idx%2==0){
      return {background:'hsl(0, 100%, 50%)'};
    } else {
      return {background:'hsl(0, 50%, 50%)'};
    }
  } else {
    if (idx%2==0){
      return {background:'hsl(100, 100%, 100%)'};
    } else {
      return {background:'rgb(200,200,255)'};
    }
  }
}

export default function Form({Id,Name,AssetType,BinName,ParentBinName,ParentName,ParentAssetType,SubMovie}) {
  return (
    <div class="demo-card-wide mdl-card mdl-shadow--2dp">
  <div class="mdl-card__title">
    <h2 class="mdl-card__title-text">{Name||'Un Titled'}</h2>
  </div>
  <div class="mdl-card__supporting-text">
    <div>{AssetType}</div>
    <div>{BinName}</div>
    <div>{Id}</div>
  </div>
  <div class="mdl-card__actions mdl-card--border">
    <a class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">
      Parent
    </a>
    <div>
      {ParentName}
    </div>
    <div>
      {ParentAssetType}
    </div>
    <div>
      {ParentBinName}
    </div>
  </div>
  <div class="mdl-card__menu">
    <button class="mdl-button mdl-button--icon mdl-js-button mdl-js-ripple-effect">
      <i class="material-icons">{ParentName}</i>
    </button>
  </div>
      <div>
        {SubMovie&&<Table rowStyleMap={rowStyleMap}  style={{height:"85vh"}}  rows={SubMovie} cols={[{name:"Name",display:"الاسم"},{name:"AssetType",display:"النوع",styleMap:styleMap}]}></Table>}
      </div>
</div>
  );
}