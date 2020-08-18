# jumhek 14 agustus 2020
# Beban untuk keledai

def convertToInt(payload):
    item = []
    for i in payload:
        item.append(int(i))
    
    return item

def main(payload):
    item = payload.split( )
    intItem = convertToInt(item)

    # check min and max length item
    if len(intItem) < 2 or len(intItem) > 100:
        return "unable-lenght"
    
    # for count item on array
    count = 0
    for i in intItem :
        count += i
    
    # to get result if count item / 2
    resp = count / 2

    # if result mod 2 is 0 must return 0
    if resp%2 == 0 :
        return int(resp % 2)

    # get average value 
    number = 0
    while number < resp:
        number += 10
    
    lest = (number - resp) * 2
    return abs(int(lest))

item = main("10 10 20 15 5 5 10 20 20 20 15 10")
# item = main("40 80 10 10 10")
# item = main("5 10")
# item = main("2 2 4 4 6 6 8 8 2 2")
print(item)