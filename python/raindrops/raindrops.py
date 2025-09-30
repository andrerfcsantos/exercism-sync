def convert(number):
    drops = ('' + ('Pling' if number % 3 == 0 else '') + ('Plang' if number % 5 == 0 else '') + ('Plong' if number % 7 == 0 else ''))
    return drops or str(number)


