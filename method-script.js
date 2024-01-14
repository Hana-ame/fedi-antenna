// usage:
// https://moonchan.xyz/page/
// paste in the second textarea.

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

class PARAMR {
  constructor(paramtype, s){
    this.paramtype = paramtype
    this.name = s
  }
  readline(s){
    this.describe = s
  }
  write(){
    let name = this.name.match(/\w+/g).join("_")

    let type = this.describe.split(".")[0]
    let gotype = ""
    if (type.startsWith("REQUIRED ")) {
      type = type.substring("REQUIRED ".length)
    }
    if (type.startsWith("Array of ")) {
      type = type.substring("Array of ".length)
      gotype = "[]"
    }    
    if (type.startsWith("String")){
      gotype += "string"
    }if (type.startsWith("Boolean")){
      gotype += "bool"
    }if (type.startsWith("Integer")){
      gotype += "int"
    }if (type.startsWith("Hash")){
      gotype += "map[string]any"
    }

    switch (this.paramtype) {
      case "Headers":
        writeline("// " + this.describe)
        writeline(`${name} := c.GetHeader("${this.name}")`)
        break;
      case "Form data parameters":
        let goVarName = this.name.split('_').map(w => {return w.charAt(0).toUpperCase() + w.slice(1).toLowerCase();}).join("").split(" ")[0]
        // let gotype = gotype
        let formVarName = this.name        
        writeline("// " + this.describe)
        writeline(`${goVarName} ${gotype} \`json:"${formVarName}"\``)
        break;
      case "Path parameters":
        writeline("// " + this.describe)
        writeline(`${name} := c.Param("${name}")`)
        break;
      case "Query parameters":
        break;
      default:
        break;
    }    
  }
}


class APIR {
  constructor(){
    this.i = 0
    this.now = ""
    this.paramr = null
  }
  readline(s, i){
    if (i === 0){
      this.name = s
    }else if (i === 1){
      this.methodPath = s
    }else if (i === 2){
      this.descripe = s
    }else if (i === 4){
      this.returns = s
      
      this.goFuncName = this.name.split(" ").join("_")
      writeline("// " + this.methodPath)
      writeline(`func ${this.goFuncName}(c *gin.Context) {`)
    }

    // Request or Response
    switch (s) {
      case "Request":
        this.now = s
        return
      case "Response":
        this.now = s
        try{
          if (this.paramr == null) return
          this.paramr.write(writeline)
          this.paramr = null
        }catch(error){
          writeline(error.stack)
        }
        writeline(`}`)
        
        return;
      default:
        break;
    }
    if (this.now !== "Request") {
      return
    }
    // which parameter
    switch (s) {
      case "Headers":
        this.paramnow = s
        break;
      case "Form data parameters":
        this.paramnow = s
        writeline(`}`)
        writeline(`// form data parameters`)
        writeline(`type ${this.goFuncName} struct{`)
        break;
      case "Path parameters":
        this.paramnow = s
        break;
      case "Query parameters":
        this.paramnow = s
        return;    
      default:
        if (this.paramr == null) {
          this.paramr = new PARAMR(this.paramnow, s)
        } else {
          this.paramr.readline(s)
          this.paramr.write()
          this.paramr = null
        }
        break;
    }
    

    this.i++;
  }
}

let apir = new APIR()

function rl(s, i) {
  // writeline("//////"+s)
   
  if (s===""){
    return
  }
  
  if (apir === null) {
    apir = new APIR()
    return
  } else {
    
    apir.readline(s, i)
    return
  } 
};

rl