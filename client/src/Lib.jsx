import React from 'react';
import Table from './Components/Table';
import From from './Components/Form';

export default class Lib extends React.Component {
  constructor(props){
    super(props)
    this.state={
      bin:[],
      movies:[],
      loading:false
    }
    this.handleClick=this.handleClick.bind(this)
    this.handleDetailsClick=this.handleDetailsClick.bind(this)
  }

  handleClick(r){
    let host=window.location.hostname
    let url=`http://${host}:8080/api/bin/${r.BinId||r.binId||r.id}`
    this.setState({loading:true})
    fetch(url).then((response)=>{
      let feed=response.json()
      feed.then((data)=>{
        this.setState({loading:false,movies:data})
      }).catch((ex)=>{
        this.setState({error:ex,loading:false})
        this.showToast({message:ex})
      })
    }).catch((err)=>{
      this.setState({error:err,loading:false})
      this.showToast({message:err})
    })
  }

  handleDetailsClick(r){
    let host=window.location.hostname
    let url=`http://${host}:8080/api/movie/${r.Id||r.id}`
    this.setState({loading:true})
    fetch(url).then((response)=>{
      let feed=response.json()
      feed.then((data)=>{
        this.setState({loading:false,movie:data})
      }).catch((ex)=>{
        this.setState({error:ex,loading:false})
        this.showToast({message:ex})
      })
    }).catch((err)=>{
      this.setState({error:err,loading:false})
      this.showToast({message:err})
    })
  }

  showToast(data){
    setTimeout(()=>{
      try
      {
        var snackbarContainer = document.querySelector('#demo-toast-example');
        snackbarContainer.MaterialSnackbar.showSnackbar(data);
      } catch (e) {
        alert(data)
      }
    },100)
  }

  componentDidMount() {
    let host=window.location.hostname
    let url=`http://${host}:8080/api/bin`
    setTimeout(()=>{
      this.setState({loading:true})
      fetch(url).then((response)=>{
        let feed=response.json()
        feed.then((data)=>{
          this.setState({loading:false,bin:data})
        }).catch((ex)=>{
          this.setState({error:ex,loading:false})
          this.showToast({message:ex})
        })
      }).catch((err)=>{
        this.setState({error:err,loading:false})
        this.showToast({message:err})
      })
    },100)
  }

  styleMap(val,row){
    if(row.Broken){
      return undefined
    }
    return     (val=="List"?{background:'hsl(39, 100%, 50%)'}:(val=="SubClip"?{background:'hsl(248, 53%, 58%)'}:{background:'hsl(147, 50%, 47%)'}))
  }
  rowStyleMap(row,idx){
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
  render(){

    return (
        <React.Fragment>
          {this.state.loading&&<div id="p2" className="mdl-progress mdl-js-progress mdl-progress__indeterminate"></div>}
          <div style={{display:'flex',direction:'column',justifyContent: 'stretch',alignItems:'stretch'}}>
            <Table style={{height:"85vh"}} onRowClick={this.handleClick} cols={[{display:"الاسم",field:"BinName"}]} rows={this.state.bin}></Table>
            <Table  onRowClick={this.handleDetailsClick} rowStyleMap={this.rowStyleMap}  style={{height:"85vh"}}  rows={this.state.movies} cols={[{name:"Name",display:"الاسم"},{name:"AssetType",display:"النوع",styleMap:this.styleMap}]}></Table>
            <From   {...this.state.movie}></From>
          </div>
        </React.Fragment>
    );
  }
}