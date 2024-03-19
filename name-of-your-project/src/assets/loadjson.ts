// 哦因为纯前端跑得所以fs用不了。
// import fs from "node:fs"

export default async function loadJSON(fn: string) {
  const resp = await fetch('/assets/'+fn+'.json')
  const json = await resp.json()
  return json
}