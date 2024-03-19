// this is used in pure web front.

import { ReactNode } from "react"

const SPAN = "SPAN"
const A = "A"

export default function render(rawHTML: string) {
  const dom = (new DOMParser()).parseFromString(rawHTML, 'text/html');

  console.log(dom.body.firstChild)

  if (dom.body.firstChild == null) return (<></>)
  return domToReact(dom.body.firstChild)
}

function domToReact(child: ChildNode): ReactNode;

function domToReact(child: ChildNode|ChildNode[]) {
  if (child === null) return (<>null</>)

  switch (child.nodeName) {
    case SPAN:
      
      break;  
    default:
      return (<>default</>)
  }
}

function domArrayToReact(children: Iterable<ChildNode|null>) {

}
