// usage:
// node open_entity.mjs FilterKeyword(ENTITY_NAME)
// node open_entity.mjs folder.ENTITY_NAME

import fs from 'fs'
import path from 'path'
import { exec } from 'child_process'

const ENTITY_NAME = process.argv[2];

const Templet = `
package entities

type ENTITY_NAME struct {

}
`

const arr = ENTITY_NAME.split(".")
const fn = arr.pop()
const filePath = arr.map(s => s.toLowerCase()).join("/")

let content = Templet.replace("ENTITY_NAME", fn)
if (filePath !== "") {
  content = content.replace("entities", filePath)

  if (!fs.existsSync(filePath)) {
    fs.mkdirSync(filePath, { recursive: true });
  }
}
fs.writeFileSync(path.join(filePath, `${fn}.go`), content, 'utf-8') 


exec('code '+path.join(filePath, `${fn}.go`), (err, stdout, stderr) => {
  if (err) {
    // node couldn't execute the command
    return;
  }

  // the *entire* stdout and stderr (buffered)
  console.log(`stdout: ${stdout}`);
  console.log(`stderr: ${stderr}`);
});