// usage:
// https://moonchan.xyz/page/
// paste in the second textarea.

class R {
  constructor(s){
    this.name = s
  }
  readline(s){
    try {
      if (s.startsWith("Description"))
        this.description = s;
      else if (s.startsWith("Type"))
        this.type = s;
      else if (s.trim() === "")
        this.write(writeline)
    } catch (error) {
      console.log(s)
      console.log(error)
    }
  }
  write(fn){
    fn("// " + this.type)
    fn("// " + this.description)
    let captialized = this.name.split('_').map(w => {return w.charAt(0).toUpperCase() + w.slice(1).toLowerCase();}).join("").split(" ")[0]
    let type = typecasting(this.type.split(" ")[1])
    if (type === "*") type += typecasting(this.type.split(" ")[2])
    if (type === "[]") type += () => { w = typecasting(this.type.split(" ")[3]).replace("::", "."); return  w.charAt(0).toLowerCase() + w.slice(1);}
    if (type === "*[]") type += () => { w = typecasting(this.type.split(" ")[4]).replace("::", "."); return  w.charAt(0).toLowerCase() + w.slice(1);}
    let jsonfield = `\`json:"${this.name}"\``.replace(" OPTIONAL", ",omitempty")
    fn(`${captialized} ${type} ${jsonfield}`)
  }
}

function typecasting(type){
  if (type.startsWith("String"))
    return "string"
  if (type.startsWith("Boolean"))
    return "bool"
  if (type.startsWith("Integer"))
    return "int"
  if (type.startsWith("NULLABLE"))
    return "*"
  if (type.startsWith("Array"))
    return "[]"
  if (type.startsWith("Hash"))
    return "map[string]any"
  return type
}

let pr = null

function rl(s, i) {
  // writeline(s)
  
  if (s===""){
    if (pr !== null) {
      try {
        pr.write(writeline)
      } catch (error) {
        writeline("// " + error)
      }
    }
    pr = null
    return
  }
  
  if (pr === null) {
    pr = new R(s)
    return
  } else {
    pr.readline(s)
    // writeline(s)
    return
  } 
};

rl