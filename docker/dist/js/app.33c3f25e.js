(function(){"use strict";var e={8826:function(e,t,n){var o=n(9242),r=n(891),a=n(3396);function u(e,t,n,o,r,u){const l=(0,a.up)("AdbForm");return(0,a.wg)(),(0,a.j4)(l)}function l(e,t,n,o,r,u){const l=(0,a.up)("a-button"),i=(0,a.up)("a-input"),m=(0,a.up)("a-form-item"),s=(0,a.up)("a-space"),c=(0,a.up)("a-form"),d=(0,a.up)("a-collapse-panel"),p=(0,a.up)("a-collapse");return(0,a.wg)(),(0,a.iD)(a.HY,null,[(0,a.Wm)(l,{type:"dashed",id:"add_item",onClick:u.addItem},{default:(0,a.w5)((()=>[(0,a.Uk)("增加一组")])),_:1},8,["onClick"]),(0,a.Wm)(p,{activeKey:r.activeKey,"onUpdate:activeKey":t[0]||(t[0]=e=>r.activeKey=e)},{default:(0,a.w5)((()=>[((0,a.wg)(!0),(0,a.iD)(a.HY,null,(0,a.Ko)(r.items,((t,n)=>((0,a.wg)(),(0,a.j4)(d,{key:n,header:t.title},{default:(0,a.w5)((()=>[(0,a.Wm)(c,{model:t.commandForm,name:"basic","label-col":{span:4},"wrapper-col":{span:16},autocomplete:"off",onFinish:e.onFinish,onFinishFailed:e.onFinishFailed},{default:(0,a.w5)((()=>[(0,a.Wm)(m,{label:"客户端ip",name:"ip",rules:[{required:!0,message:"请输入客户端ip"}]},{default:(0,a.w5)((()=>[(0,a.Wm)(i,{value:t.commandForm.ip,"onUpdate:value":e=>t.commandForm.ip=e},null,8,["value","onUpdate:value"])])),_:2},1024),(0,a.Wm)(m,{label:"端口",name:"port",rules:[{required:!1,message:""}]},{default:(0,a.w5)((()=>[(0,a.Wm)(i,{value:t.commandForm.port,"onUpdate:value":e=>t.commandForm.port=e},null,8,["value","onUpdate:value"])])),_:2},1024),(0,a.Wm)(m,{label:"包名",name:"packageName",rules:[{required:!1,message:""}]},{default:(0,a.w5)((()=>[(0,a.Wm)(i,{value:t.commandForm.packageName,"onUpdate:value":e=>t.commandForm.packageName=e},null,8,["value","onUpdate:value"])])),_:2},1024),(0,a.Wm)(m,{label:"代理地址",name:"proxyAddr",rules:[{required:!1,message:""}]},{default:(0,a.w5)((()=>[(0,a.Wm)(i,{value:t.commandForm.proxyAddr,"onUpdate:value":e=>t.commandForm.proxyAddr=e},null,8,["value","onUpdate:value"])])),_:2},1024),(0,a.Wm)(m,{label:"自由命令",name:"cmd",rules:[{required:!1,message:""}]},{default:(0,a.w5)((()=>[(0,a.Wm)(i,{value:t.commandForm.cmd,"onUpdate:value":e=>t.commandForm.cmd=e},null,8,["value","onUpdate:value"])])),_:2},1024),(0,a.Wm)(m,{"wrapper-col":{offset:4,span:16}},{default:(0,a.w5)((()=>[(0,a.Wm)(s,{wrap:""},{default:(0,a.w5)((()=>[(0,a.Wm)(l,{type:"primary","html-type":"submit",onClick:e=>u.sendHttpRequest(t.commandForm,"setProxy")},{default:(0,a.w5)((()=>[(0,a.Uk)("设置代理")])),_:2},1032,["onClick"]),(0,a.Wm)(l,{type:"primary","html-type":"submit",onClick:e=>u.sendHttpRequest(t.commandForm,"delProxy")},{default:(0,a.w5)((()=>[(0,a.Uk)("清除代理")])),_:2},1032,["onClick"]),(0,a.Wm)(l,{type:"primary",danger:"","html-type":"submit",onClick:e=>u.sendHttpRequest(t.commandForm,"stop")},{default:(0,a.w5)((()=>[(0,a.Uk)("终止进程")])),_:2},1032,["onClick"]),(0,a.Wm)(l,{type:"primary",danger:"","html-type":"submit",onClick:e=>u.sendHttpRequest(t.commandForm,"clear")},{default:(0,a.w5)((()=>[(0,a.Uk)("清理缓存")])),_:2},1032,["onClick"]),(0,a.Wm)(l,{type:"primary","html-type":"submit",onClick:e=>u.sendHttpRequest(t.commandForm,"free")},{default:(0,a.w5)((()=>[(0,a.Uk)("执行输入命令")])),_:2},1032,["onClick"])])),_:2},1024)])),_:2},1024)])),_:2},1032,["model","onFinish","onFinishFailed"])])),_:2},1032,["header"])))),128))])),_:1},8,["activeKey"])],64)}n(7658);var i=n(252),m=n(4870),s=n(4161);function c(e,t,n={},o={}){return t&&(o.data=t),d(e,"post",n,o)}function d(e,t,n,o){return o.url=e,o.method=t||"get",(0,s.Z)(o).then((t=>"string"===typeof t?Promise.reject({response:{statusText:`接口${e}未返回正确格式`}}):(n.successMsg&&i.ZP.success("string"===typeof n.successMsg?n.successMsg:"操作成功"),Promise.resolve(t)))).catch((e=>{let t=n.errorMsg||e.response&&e.response.data&&e.response.data.errMessage||e.response.statusText||"未知错误";return!1!==n.errorMsg&&i.ZP.error(t),Promise.reject(e)}))}s.Z.defaults.timeout=25e3;const p={ip:"",port:"5555"};var f={data(){return{activeKey:(0,m.iH)(["0"]),items:[{title:"adb控制",inputValue:"",commandForm:Object.assign({},p)}]}},methods:{sendHttpRequest(e,t){e.op=t,c("/adb",e,{errorMsg:!0}).then((e=>{i.ZP.info({content:()=>(0,a.h)("pre",{class:"shell-output"},e.data.data.split("\n").map((e=>(0,a.h)("p",{},e))))})})).catch((e=>{console.log(e)}))},addItem(){if(this.items.length>=5)return void i.ZP.error("已达到最大数量！");const e=prompt("输入标题");e&&this.items.push({title:e,inputValue:"",commandForm:Object.assign({},p)})}}},v=n(89);const y=(0,v.Z)(f,[["render",l],["__scopeId","data-v-4079cbd6"]]);var b=y,h={name:"App",components:{AdbForm:b}};const g=(0,v.Z)(h,[["render",u]]);var w=g;n(1849);(0,o.ri)(w).use(r.ZP).mount("#app")}},t={};function n(o){var r=t[o];if(void 0!==r)return r.exports;var a=t[o]={exports:{}};return e[o].call(a.exports,a,a.exports,n),a.exports}n.m=e,function(){var e=[];n.O=function(t,o,r,a){if(!o){var u=1/0;for(s=0;s<e.length;s++){o=e[s][0],r=e[s][1],a=e[s][2];for(var l=!0,i=0;i<o.length;i++)(!1&a||u>=a)&&Object.keys(n.O).every((function(e){return n.O[e](o[i])}))?o.splice(i--,1):(l=!1,a<u&&(u=a));if(l){e.splice(s--,1);var m=r();void 0!==m&&(t=m)}}return t}a=a||0;for(var s=e.length;s>0&&e[s-1][2]>a;s--)e[s]=e[s-1];e[s]=[o,r,a]}}(),function(){n.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return n.d(t,{a:t}),t}}(),function(){n.d=function(e,t){for(var o in t)n.o(t,o)&&!n.o(e,o)&&Object.defineProperty(e,o,{enumerable:!0,get:t[o]})}}(),function(){n.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(e){if("object"===typeof window)return window}}()}(),function(){n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)}}(),function(){n.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})}}(),function(){var e={143:0};n.O.j=function(t){return 0===e[t]};var t=function(t,o){var r,a,u=o[0],l=o[1],i=o[2],m=0;if(u.some((function(t){return 0!==e[t]}))){for(r in l)n.o(l,r)&&(n.m[r]=l[r]);if(i)var s=i(n)}for(t&&t(o);m<u.length;m++)a=u[m],n.o(e,a)&&e[a]&&e[a][0](),e[a]=0;return n.O(s)},o=self["webpackChunkadb_web"]=self["webpackChunkadb_web"]||[];o.forEach(t.bind(null,0)),o.push=t.bind(null,o.push.bind(o))}();var o=n.O(void 0,[998],(function(){return n(8826)}));o=n.O(o)})();
//# sourceMappingURL=app.33c3f25e.js.map