import{a as o,r as e}from"./app-62185dda.js";function m({message:r,className:t=""}){return r?o("p",{className:"text-sm text-red-600 dark:text-red-400 "+t,children:r}):null}const k=e.forwardRef(function({type:t="text",name:d,id:s,value:i,className:u,autoComplete:c,required:f,isFocused:x,handleChange:g},n){const a=n||e.useRef();return e.useEffect(()=>{x&&a.current.focus()},[]),o("div",{className:"flex flex-col items-start",children:o("input",{type:t,name:d,id:s,value:i,className:"border-gray-300 dark:border-gray-700 dark:bg-gray-900 dark:text-gray-300 focus:border-indigo-500 dark:focus:border-indigo-600 focus:ring-indigo-500 dark:focus:ring-indigo-600 rounded-md shadow-sm "+u,ref:a,autoComplete:c,required:f,onChange:l=>g(l)})})});export{m as I,k as T};
