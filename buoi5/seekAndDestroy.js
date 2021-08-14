function seekAndDestroy(arr, ...theArgs) {
    for (v of theArgs) {
     arr = arr.filter(x => x != v)
    }
    return arr
  }
  