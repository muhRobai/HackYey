# jumhek 20200812

# create an object
class Item:
    def __init__(self, position, value):
        self.position = position
        self.value = value

# define position all value
def tapeArr(payload):
    items = []
    count = 1
    for i in payload:
        x = Item(count, i)
        items.append(x)
        count = count + 1

    return items

def sort(ar):
    n = len(ar)
    # Traverse through all array elements
    for i in range(n):
        # Last i elements are already in correct position
        for j in range(0, n-i-1):
            # Swap if the element found is greater than the next element
            if ar[j] > ar[j+1]:
                ar[j], ar[j+1] = ar[j+1], ar[j]
    
    return ar

def convertToInt(payload):
    item = []
    for i in payload:
        item.append(int(i))
    
    return item

# execute 
def execItem(payload):
    item = payload.split( )
    intArr = convertToInt(item)
    sorts = sort(intArr)
    position = tapeArr(item)
    start = 1
    value = 0
    total = 0
    for i in sorts :
        for j in position :
            if i == int(j.value) :
                if start > int(j.position) :
                    value = start - j.position
                else :
                    value = j.position - start
                start = j.position
                total = total + (value * 10)
    
    return total
    

# items = execItem("4 8 2 1 5 7 9")
# items = execItem("40 50 20 21 22 23 24 25")
# items = execItem("1 2 3 4 5 6 7 8 9")
items = execItem("100 95 90 80 40 10 8")

print(items)