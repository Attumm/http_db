(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{330:function(e,t,a){e.exports=a(538)},472:function(e,t,a){},474:function(e,t,a){},535:function(e,t,a){},538:function(e,t,a){"use strict";a.r(t);a(331),a(336),a(338),a(341),a(347),a(365),a(367);var n=a(1),r=a.n(n),s=a(7),i=a.n(s),o=a(541),l=a(59),c=a(60),u=a(63),p=a(61),h=a(64),d=a(78),g=r.a.createContext(),m=function(e,t){switch(t.type){case"SET_SEARCH_RESULTS":return Object(d.a)({},e,{results:t.payload});case"SHOW_TABLE_LOADING_OVERLAY":return Object(d.a)({},e,{tableResultsAreLoading:t.payload});case"SHOW_RESULTS_TABLE":return Object(d.a)({},e,{showResultsTable:t.payload});default:return e}},f=function(e){function t(){var e,a;Object(l.a)(this,t);for(var n=arguments.length,r=new Array(n),s=0;s<n;s++)r[s]=arguments[s];return(a=Object(u.a)(this,(e=Object(p.a)(t)).call.apply(e,[this].concat(r)))).state={results:[],tableResultsAreLoading:!1,showResultsTable:!1,dispatch:function(e){return a.setState(function(t){return m(t,e)})}},a}return Object(h.a)(t,e),Object(c.a)(t,[{key:"render",value:function(){var e=Object(d.a)({},this.state),t=this.props.children;return r.a.createElement(g.Provider,{value:e},t)}}]),t}(n.Component),b=(g.Consumer,a(403),a(327)),v=(a(318),a(161)),S=(a(411),a(229)),w=(a(412),a(20)),O=(a(415),a(326)),y=a(225),R=a.n(y),E=a(320),C=a(121),L=(a(420),a(228)),A=(a(319),a(99)),T=a(321),j=a(122),D=a.n(j),H=window.__env.searchApi;D.a.interceptors.request.use(function(e){return e},function(e){return Promise.reject(e)});var _=function(e,t,a){return D.a.get("".concat(H).concat(e),a)},x=a(322),W=a(6),N=a.n(W),k=(a(472),function(e){function t(){var e,a;Object(l.a)(this,t);for(var n=arguments.length,s=new Array(n),i=0;i<n;i++)s[i]=arguments[i];return(a=Object(u.a)(this,(e=Object(p.a)(t)).call.apply(e,[this].concat(s)))).state={gridApi:null,gridParams:null},a.gridWrapperRef=r.a.createRef(),a.overlayLoadingTemplate='<div class="spin-wrapper">\n    <div class="ant-spin ant-spin-lg ant-spin-spinning">\n    <span class="ant-spin-dot ant-spin-dot-spin"><i class="ant-spin-dot-item"></i><i class="ant-spin-dot-item"></i><i class="ant-spin-dot-item"></i><i class="ant-spin-dot-item"></i></span>\n    </div>\n    </div>',a.overlayNoRowsTemplate='<div class="no-results-wrapper">\n    <div class="message">\n    <i class="anticon anticon-exclamation-circle">\n    <svg viewBox="64 64 896 896" class="" data-icon="exclamation-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M464 688a48 48 0 1 0 96 0 48 48 0 1 0-96 0zM488 576h48c4.4 0 8-3.6 8-8V296c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8z"></path></svg>\n    </i> \n    <span>No rows to show</span>\n    </div>\n    </div>',a.onGridReady=function(e){var t=a.props.autoHeight;a.setState({gridApi:e.api,gridParams:e,gridColumnApi:e.columnApi}),t&&e.api.resetRowHeights()},a.showRequiredOverlay=function(e){var t=a.state.gridApi,n=a.props.rowData;t&&(e?t.showLoadingOverlay():e||n.length?t.hideOverlay():t.showNoRowsOverlay())},a.createColumnDefsWithOptions=function(){var e=a.props,t=e.rowData,n=e.columns;if(!t.length)return[];var r=[],s=!0,i=!1,o=void 0;try{for(var l,c=n[Symbol.iterator]();!(s=(l=c.next()).done);s=!0){var u=l.value;u.isShow&&r.push(Object(d.a)({},a.defaultColumnDefs,{headerName:u.label,field:u.name}))}}catch(p){i=!0,o=p}finally{try{s||null==c.return||c.return()}finally{if(i)throw o}}return r.map(function(e){return e})},a}return Object(h.a)(t,e),Object(c.a)(t,[{key:"componentDidUpdate",value:function(e){}},{key:"componentWillUnmount",value:function(){var e=this.state.gridApi;e&&e.destroy&&e.destroy()}},{key:"render",value:function(){var e=this.state.gridApi,t=this.props,a=t.tableHeight,n=t.enableColResize,s=t.rowData,i=t.enableSorting,o=t.enableFilter,l=t.floatingFilter,c=t.suppressMenu,u=t.headerHeight,p=t.floatingFiltersHeight,h=t.rowHeight,d=t.pagination,g=t.context,m=t.autoHeight,f=t.tableResultsAreLoading;this.showRequiredOverlay(f);var b=this.createColumnDefsWithOptions(),v=N()("ag-theme-balham",{"auto-height":m});if(this.gridWrapperRef.current&&e){var S=this.gridWrapperRef.current.clientWidth;320*b.length<=S&&e.sizeColumnsToFit()}return r.a.createElement("div",{ref:this.gridWrapperRef,className:v,style:{height:"".concat(a,"px"),visibility:e?"visible":"hidden"}},r.a.createElement(x.AgGridReact,{enableColResize:n,columnDefs:b,rowData:s,enableSorting:i,enableFilter:o,floatingFilter:l,suppressMenu:c,headerHeight:u,floatingFiltersHeight:p,rowHeight:h,pagination:d,suppressRowClickSelection:!0,suppressCellSelection:!0,rowSelection:"single",overlayLoadingTemplate:this.overlayLoadingTemplate,overlayNoRowsTemplate:this.overlayNoRowsTemplate,getRowHeight:this.getRowHeight,onGridReady:this.onGridReady,context:g,suppressColumnMoveAnimation:!0,animateRows:!0,defaultColDef:{width:320}}))}}]),t}(n.Component));k.defaultProps={pagination:!0,enableFilter:!0,enableSorting:!0,enableColResize:!0,floatingFilter:!0,suppressMenu:!0,autoHeight:!1,isCmaTable:!1,isReportsCmaTable:!1,tableResultsAreLoading:!1,headerHeight:40,floatingFiltersHeight:40,rowHeight:38,context:{},hiddenColumns:[],filterData:{}};var z=k,F=(a(474),A.a.Option),G=L.a.OptGroup,V=function(e){function t(){var e,a;Object(l.a)(this,t);for(var n=arguments.length,s=new Array(n),i=0;i<n;i++)s[i]=arguments[i];return(a=Object(u.a)(this,(e=Object(p.a)(t)).call.apply(e,[this].concat(s)))).signal=D.a.CancelToken.source(),a.defaultColumnDefs={suppressMenu:!0},a.state={tableHeight:500,hiddenColumns:[],serachOptions:[{Title:"Default",value:"default"}],serachOption:"default",searchResults:[],searchFields:[],searchView:"table",showResultsTable:!1,isLoading:!1,searchValue:"",context:{componentParent:Object(C.a)(Object(C.a)(a))},autoCompleteDataSource:{resentSearch:[],suggessions:[]}},a.tableWrapperRef=r.a.createRef(),a.debouncedOnChange=Object(T.debounce)(function(e){a.getSearchResults(e)},500),a.onChange=function(e,t){var n=a.state.autoCompleteDataSource;a.setState({autoCompleteDataSource:{resentSearch:n.resentSearch,suggessions:t.length?["loading"]:[]}}),a.setState({searchValue:t},function(){a.debouncedOnChange(e)})},a.getSearchResults=function(){var e=Object(E.a)(R.a.mark(function e(t){var n,r,s,i,o,l;return R.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:if(n=a.state,r=n.searchValue,s=n.serachOption,i=a.state.autoCompleteDataSource,a.handleWindowResize(),o={},a.lastSession=o,t({type:"SHOW_TABLE_LOADING_OVERLAY",payload:!0}),r){e.next=10;break}return t({type:"SET_SEARCH_RESULTS",payload:[]}),t({type:"SHOW_TABLE_LOADING_OVERLAY",payload:!1}),e.abrupt("return");case 10:return a.setState({autoCompleteDataSource:{resentSearch:i.resentSearch,suggessions:[]}}),a.setState({showResultsTable:!0,isLoading:!0}),e.prev=12,e.next=15,_(r,s,{cancelToken:a.signal.token});case 15:if(l=e.sent,a.lastSession===o){e.next=18;break}return e.abrupt("return");case 18:a.onGetSearchSuccess(t,l),e.next=25;break;case 21:e.prev=21,e.t0=e.catch(12),console.log(e.t0),a.onGetSearchError(t);case 25:case"end":return e.stop()}},e,null,[[12,21]])}));return function(t){return e.apply(this,arguments)}}(),a.onGetSearchSuccess=function(e,t){a.parseTableData(t.data?t.data:[]),a.setState({isLoading:!1}),e({type:"SHOW_TABLE_LOADING_OVERLAY",payload:!1})},a.parseTableData=function(e){var t=e.meta;a.setState({searchResults:e.data,searchFields:e.meta.fields,searchView:t.view,showResultsTable:!0})},a.onGetSearchError=function(e){O.a.error("Unfortunately there was an error getting the results"),a.setState({isLoading:!1}),e({type:"SHOW_TABLE_LOADING_OVERLAY",payload:!1})},a.setTableDimensions=function(){var e=a.tableWrapperRef.current.getBoundingClientRect().top,t=window.innerHeight-e-24;a.setState({tableHeight:t})},a.handleWindowResize=function(){a.setTableDimensions()},a.searchOptionChange=function(e){a.setState({serachOption:e})},a.renderOption=function(){var e=a.state.autoCompleteDataSource;return Object.keys(e).map(function(t){return r.a.createElement(G,{key:t,label:"suggessions"===t?"Suggestions":"Recent Searches"},e[t].map(function(e){return r.a.createElement(F,{key:e,text:e},"loading"===e?r.a.createElement(w.a,{type:"loading"}):r.a.createElement("div",{className:"global-search-item"},r.a.createElement("span",{className:"global-search-item-desc"},e),"resentSearch"===t?r.a.createElement("span",{className:"search-history-remove"},"Remove"):""))}))})},a}return Object(h.a)(t,e),Object(c.a)(t,[{key:"componentWillMount",value:function(){var e=this.context.dispatch;window.addEventListener("resize",this.handleWindowResize),e({type:"SHOW_TABLE_LOADING_OVERLAY",payload:!1}),e({type:"SHOW_RESULTS_TABLE",payload:!1}),this.setState({autoCompleteDataSource:{resentSearch:[],suggessions:[]}})}},{key:"componentWillUnmount",value:function(){this.signal.cancel("Canceled"),window.removeEventListener("resize",this.handleWindowResize)}},{key:"render",value:function(){var e=this,t=this.context.dispatch,a=this.state,s=a.serachOptions,i=a.serachOption,o=a.searchResults,l=a.searchFields,c=a.tableHeight,u=a.hiddenColumns,p=a.searchView,h=a.showResultsTable,d=a.context,g=a.isLoading;return r.a.createElement(n.Fragment,null,r.a.createElement("div",{className:"search-input-wrapper initial-search-done"},r.a.createElement(b.a,{gutter:12},r.a.createElement(S.a,{span:6},r.a.createElement(A.a,{placeholder:"Select Option",className:"search-select-options",onSelect:this.searchOptionChange,value:i},s.map(function(e){return r.a.createElement(F,{value:e.value,key:e.value},e.Title)}))),r.a.createElement(S.a,{span:18},r.a.createElement("div",{className:"global-search-wrapper"},r.a.createElement(L.a,{className:"search-input",size:"large",style:{width:"100%"},dropdownStyle:{width:"500px"},onSearch:this.handleSearch,optionLabelProp:"text",onChange:function(a){return e.onChange(t,a)},ref:function(t){e.searchInput=t}},r.a.createElement(v.a,{className:"search-input",placeholder:"Search...",prefix:r.a.createElement(w.a,{type:"search",style:{color:"rgba(0,0,0,.25)"}}),size:"large"})))))),r.a.createElement("div",{className:"search-table-wrapper"},r.a.createElement("div",{ref:this.tableWrapperRef,className:"search-results"},"table"===p&&h&&r.a.createElement(z,{ref:this.dataTableRef,columns:l,rowData:o,tableHeight:c,hiddenColumns:u,context:d,tableResultsAreLoading:g,enableFilter:!1,enableSorting:!1,floatingFilter:!1,suppressMenu:!1}))))}}]),t}(n.Component);V.contextType=g;var B=V,M=(a(535),function(e){function t(){return Object(l.a)(this,t),Object(u.a)(this,Object(p.a)(t).apply(this,arguments))}return Object(h.a)(t,e),Object(c.a)(t,[{key:"render",value:function(){return r.a.createElement(f,null,r.a.createElement(B,null))}}]),t}(r.a.Component));i.a.render(r.a.createElement(o.a,null,r.a.createElement(M,null)),document.getElementById("root"))}},[[330,2,1]]]);
//# sourceMappingURL=main.eab68b13.chunk.js.map