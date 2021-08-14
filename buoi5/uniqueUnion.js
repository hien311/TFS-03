function uniqueUnion(...theArgs) {
    let a = []
    while (theArgs.length > 0) {
      a = [...a,...theArgs.shift()]
    }
    return  [... new Set(a)]
  }
  
  export default uniqueUnion