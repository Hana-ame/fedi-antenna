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
  getName(){
    let name = this.name.match(/\w+/g).join("_")
    return name
  }
  write(){
    let name = this.getName()

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
        name = name.match(/\w+/g).join("")
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
        writeline("// " + this.describe)
        writeline(`${name} := c.Query("${name}")`)
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
    this.otherparams = []
    this.dataformparams = []
  }
  // api
  readline(s, i){
    // writeline("//////"+s)
    // meta datas
    if (i === 0){
      this.name = s
    }else if (i === 1){
      this.methodPath = s
    }else if (i === 2){
      this.descripe = s
    }else if (i === 4){
      this.returns = s
    }

    // writeline("//////"+s)
    // Request or Response
    switch (s) {
      case "Request":
        this.now = s
        return
      case "Response":
        // finally
        this.write()

        this.now = s
        try{
          if (this.paramr == null) break;
          this.paramr.write(writeline)
          this.paramr = null
        }catch(error){
          writeline(error.stack)
        }

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
          if (this.paramr.paramtype === "Form data parameters") this.dataformparams.push(this.paramr)
          else this.otherparams.push(this.paramr)
          this.paramr = null
        }
        break;
    }
    

    this.i++;
  }
  write(){        
    writeline("package controller")

    this.goFuncName = this.name.split(" ").join("_")
    writeline("// " + this.methodPath)
    writeline(`func ${this.goFuncName}(c *gin.Context) {`)
    
    for(const p of this.otherparams){
      p.write()
    }

    if (this.dataformparams.length > 0) {
      writeline(`var data *model.${this.goFuncName}`)
      writeline(`c.Bind(&data)`)
    }
    
	  writeline(`o, err := handler.${this.goFuncName}(`)
    for(const p of this.otherparams){
      writeline(`${p.getName()},`)
    }
    if (this.dataformparams.length > 0) {
      writeline(`data,`)
    }
    writeline(`)`)
    writeline(`if err != nil {`)
		writeline(`  c.JSON(http.StatusInternalServerError, err)`)
		writeline(`  return`)
    writeline(`}`)

    writeline(`c.JSON(http.StatusOK, o)`)
    writeline(`return`)
    writeline(`}`)

    if (this.dataformparams.length !== 0) {

      writeline("package model")
      writeline(`// form data parameters`)
      writeline(`type ${this.goFuncName} struct{`)
      
      for(const p of this.dataformparams){
        p.write()
      }

      writeline(`}`)
    }
    let arr = this.methodPath.split(" ")
    writeline( `r.${arr[0]}("${arr[1]}",controller.${this.goFuncName})`)
  }
}

let apir = new APIR()

function rl(s, i) {
  // writeline("//////"+s)
   
  if (s===""){
    return
  }
  
  // api reader entrance
  if (apir === null) {
    apir = new APIR()
    return
  } else {
    
    apir.readline(s, i)
    return
  } 
  
};

rl