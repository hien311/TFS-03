function drop(arr,callback) {
    for (let i in arr) {
      if (callback(arr[i]) == true ) {
        return arr.slice(i)
      }
    }
    return []
  }
export {drop}