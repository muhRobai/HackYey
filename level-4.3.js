myFunction=(item)=>{
  let isOk = true
  let enkrip = ''
  let teks = /^[a-zA-Z ]+$/ 
  
  // for check only alfabeth
  if (teks.test(item) === false || item.length > 1000) {
    return console.log('masukan teks sesuai kriteria')
  }

  //splite item to array
  let items = item.split(" ")
  
  //get item 
  for (let i = 0; i < items.length; i++) {
    if (items[i].length < 3) {
      isOk = false
    }
    //enkrip items
    for (let j = items[i].length - 1 ; j >= 0; j--) {
        enkrip += items[i][j]
        if (j === 0) {
          enkrip += ' '
        }
    }
  }

  if (!isOk) {
    console.log('masukan teks sesuai kriteria')
    return 
  }
  console.log(enkrip)
}

myFunction('iadab itsap ulalreb')
myFunction('italem irad irigayaj')
myFunction('nalub kusutret gnalali')

