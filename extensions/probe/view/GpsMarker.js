var GpsMarker;(()=>{var e={598:(e,t,r)=>{"use strict";r.r(t),r.d(t,{default:()=>s});var n=r(81),a=r.n(n),o=r(645),i=r.n(o)()(a());i.push([e.id,"\n.amap-page-container[data-v-65edabdb]{height: calc(100% - 30px); width: 100%;}\n.amap-wrapper[data-v-65edabdb] {\r\n  width: 100%;\r\n  height: 100%;\n}\r\n",""]);const s=i},645:e=>{"use strict";e.exports=function(e){var t=[];return t.toString=function(){return this.map((function(t){var r="",n=void 0!==t[5];return t[4]&&(r+="@supports (".concat(t[4],") {")),t[2]&&(r+="@media ".concat(t[2]," {")),n&&(r+="@layer".concat(t[5].length>0?" ".concat(t[5]):""," {")),r+=e(t),n&&(r+="}"),t[2]&&(r+="}"),t[4]&&(r+="}"),r})).join("")},t.i=function(e,r,n,a,o){"string"==typeof e&&(e=[[null,e,void 0]]);var i={};if(n)for(var s=0;s<this.length;s++){var c=this[s][0];null!=c&&(i[c]=!0)}for(var l=0;l<e.length;l++){var d=[].concat(e[l]);n&&i[d[0]]||(void 0!==o&&(void 0===d[5]||(d[1]="@layer".concat(d[5].length>0?" ".concat(d[5]):""," {").concat(d[1],"}")),d[5]=o),r&&(d[2]?(d[1]="@media ".concat(d[2]," {").concat(d[1],"}"),d[2]=r):d[2]=r),a&&(d[4]?(d[1]="@supports (".concat(d[4],") {").concat(d[1],"}"),d[4]=a):d[4]="".concat(a)),t.push(d))}},t}},81:e=>{"use strict";e.exports=function(e){return e[1]}},265:(e,t,r)=>{var n=r(598);n.__esModule&&(n=n.default),"string"==typeof n&&(n=[[e.id,n,""]]),n.locals&&(e.exports=n.locals),(0,r(346).Z)("10561795",n,!1,{})},346:(e,t,r)=>{"use strict";function n(e,t){for(var r=[],n={},a=0;a<t.length;a++){var o=t[a],i=o[0],s={id:e+":"+a,css:o[1],media:o[2],sourceMap:o[3]};n[i]?n[i].parts.push(s):r.push(n[i]={id:i,parts:[s]})}return r}r.d(t,{Z:()=>h});var a="undefined"!=typeof document;if("undefined"!=typeof DEBUG&&DEBUG&&!a)throw new Error("vue-style-loader cannot be used in a non-browser environment. Use { target: 'node' } in your Webpack config to indicate a server-rendering environment.");var o={},i=a&&(document.head||document.getElementsByTagName("head")[0]),s=null,c=0,l=!1,d=function(){},p=null,u="data-vue-ssr-id",f="undefined"!=typeof navigator&&/msie [6-9]\b/.test(navigator.userAgent.toLowerCase());function h(e,t,r,a){l=r,p=a||{};var i=n(e,t);return v(i),function(t){for(var r=[],a=0;a<i.length;a++){var s=i[a];(c=o[s.id]).refs--,r.push(c)}for(t?v(i=n(e,t)):i=[],a=0;a<r.length;a++){var c;if(0===(c=r[a]).refs){for(var l=0;l<c.parts.length;l++)c.parts[l]();delete o[c.id]}}}}function v(e){for(var t=0;t<e.length;t++){var r=e[t],n=o[r.id];if(n){n.refs++;for(var a=0;a<n.parts.length;a++)n.parts[a](r.parts[a]);for(;a<r.parts.length;a++)n.parts.push(g(r.parts[a]));n.parts.length>r.parts.length&&(n.parts.length=r.parts.length)}else{var i=[];for(a=0;a<r.parts.length;a++)i.push(g(r.parts[a]));o[r.id]={id:r.id,refs:1,parts:i}}}}function m(){var e=document.createElement("style");return e.type="text/css",i.appendChild(e),e}function g(e){var t,r,n=document.querySelector("style["+u+'~="'+e.id+'"]');if(n){if(l)return d;n.parentNode.removeChild(n)}if(f){var a=c++;n=s||(s=m()),t=M.bind(null,n,a,!1),r=M.bind(null,n,a,!0)}else n=m(),t=C.bind(null,n),r=function(){n.parentNode.removeChild(n)};return t(e),function(n){if(n){if(n.css===e.css&&n.media===e.media&&n.sourceMap===e.sourceMap)return;t(e=n)}else r()}}var y,b=(y=[],function(e,t){return y[e]=t,y.filter(Boolean).join("\n")});function M(e,t,r,n){var a=r?"":n.css;if(e.styleSheet)e.styleSheet.cssText=b(t,a);else{var o=document.createTextNode(a),i=e.childNodes;i[t]&&e.removeChild(i[t]),i.length?e.insertBefore(o,i[t]):e.appendChild(o)}}function C(e,t){var r=t.css,n=t.media,a=t.sourceMap;if(n&&e.setAttribute("media",n),p.ssrId&&e.setAttribute(u,t.id),a&&(r+="\n/*# sourceURL="+a.sources[0]+" */",r+="\n/*# sourceMappingURL=data:application/json;base64,"+btoa(unescape(encodeURIComponent(JSON.stringify(a))))+" */"),e.styleSheet)e.styleSheet.cssText=r;else{for(;e.firstChild;)e.removeChild(e.firstChild);e.appendChild(document.createTextNode(r))}}}},t={};function r(n){var a=t[n];if(void 0!==a)return a.exports;var o=t[n]={id:n,exports:{}};return e[n](o,o.exports,r),o.exports}r.n=e=>{var t=e&&e.__esModule?()=>e.default:()=>e;return r.d(t,{a:t}),t},r.d=(e,t)=>{for(var n in t)r.o(t,n)&&!r.o(e,n)&&Object.defineProperty(e,n,{enumerable:!0,get:t[n]})},r.o=(e,t)=>Object.prototype.hasOwnProperty.call(e,t),r.r=e=>{"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})};var n={};(()=>{"use strict";r.r(n),r.d(n,{default:()=>o});var e=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"amap-page-container"},[r("div",{staticClass:"amap-wrapper",attrs:{id:"markermap"}},[r("el-amap",{attrs:{vid:"amapGpsMarker",zoom:e.zoom,center:e.center,events:e.mapEvents}},e._l(e.markers,(function(e,t){return r("el-amap-marker",{key:t,attrs:{position:e.position,title:e.title}})})),1)],1)])};e._withStripped=!0;const t={name:"GpsMarker",props:{loading:{type:Boolean,default:!0},apiData:{type:Object},title:{type:String,default:""}},data:()=>({markers:[{title:"测试",icon:"img/mark_b.png",position:[116.405467,39.907761],text:"测试地址"}],center:[116.397128,39.916527],zoom:7,mapEvents:{init(e){e.setMapStyle("amap://styles/normal")}}}),computed:{},watch:{apiData:{immediate:!0,handler(e,t){this.loading||this.initChart()}}},created(){this.initMap()},methods:{initMap(){console.log("---------------GPSMARKER1----------------"),setTimeout((()=>{this.initChart()}),1e3)},initChart(){console.log("---------------GPSMARKER2----------------");let e=this.apiData.fields,t=e[e.length-1];for(let t=0;t<e.length;t++)this.markers.push({title:e[t].title,icon:"img/mark_b.png",position:[e[t].lng,e[t].lat],text:e[t].address});this.center=[t.lng,t.lat],console.log("markers",this.markers),console.log("center",this.center)}},async mounted(){}};r(265);var a=function(e,t,r,n,a,o,i,s){var c,l="function"==typeof e?e.options:e;if(t&&(l.render=t,l.staticRenderFns=[],l._compiled=!0),l._scopeId="data-v-65edabdb",c)if(l.functional){l._injectStyles=c;var d=l.render;l.render=function(e,t){return c.call(t),d(e,t)}}else{var p=l.beforeCreate;l.beforeCreate=p?[].concat(p,c):[c]}return{exports:e,options:l}}(t,e);a.options.__file="src/GpsMarker.vue";const o=a.exports})(),GpsMarker=n})();