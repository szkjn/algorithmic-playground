
"""
Interview question of the week by Cassidy Williams (issue #273) :

    Given a positive integer n, return all of its anti-divisors. 
    Anti-divisors are numbers that do not divide a number by the 
    largest possible margin (1 is not an anti-divisor of any number). 
    More information here : https://oeis.org/A066272/a066272a.html
  
"""

def antidivisor(n):

  res = []
  
  for k in range(2, n):
    
    if k % 2 == 0:
      if n % k == k / 2:
        res.append(k)  
    else:
      if (
        n % k == (k-1)/2 
        or n % k == (k+1)/2
      ):
        res.append(k)
  
  print(res)

antidivisor(234)
