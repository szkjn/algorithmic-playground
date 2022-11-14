def combineStrings(arr: list, n: int) -> list:

    res = []
    new_str = ""
    
    while len(arr) > 0:

        if new_str == "":           
            new_str = arr[0]
            arr.pop(0)
            if len(arr) == 0: break

        if len(new_str) + len(arr[0]) < n:
            new_str += " " + arr[0]
            arr.pop(0)
            if len(arr) == 0: break
            
        if (
            len(new_str) == n
            or (
                len(new_str) < n
                and len(new_str) + len(arr[0]) >= n
                )
            ):
            res.append(new_str)
            new_str = ""

        if len(arr) == 1:         
            res.append(arr[0])
            arr.pop(0)
            if len(arr) == 0: break

    if len(new_str) > 0:
        res.append(new_str)

    print(res)

combineStrings(["a", "b", "c", "d", "e", "f", "g"], 5)
combineStrings(["a", "b", "c", "d", "e", "f", "g"], 12)
combineStrings(["alpha", "beta", "gamma", "delta", "epsilon"], 20)
