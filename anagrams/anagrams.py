def isAnagram(s: str, t: str) -> bool:
    if len(s) == len(t):
        s_map = {}
        t_map = {}
        for l in s:
            if l in s_map:
                s_map[l] += 1
            else:
                s_map[l] = 0
        for l in t:
            if l in t_map:
                t_map[l] += 1
            else:
                t_map[l] = 0

        for l in s_map:
            if l not in t_map or s_map[l] != t_map[l]:
                return False
    else:
        return False
    return True
