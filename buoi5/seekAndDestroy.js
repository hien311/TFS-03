function seekAndDestroy(arr, ...theArgs) {
    for (let v of theArgs) {
     arr = arr.filter(x => x != v)
    }
    return arr
  }
  
export default seekAndDestroy