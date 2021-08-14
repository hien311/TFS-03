function drop(arr,callback) {
    for (i in arr) {
      if (callback(arr[i]) == true ) {
        return arr.slice(i)
      }
    }
    return []
  }
  