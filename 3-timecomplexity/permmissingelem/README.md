# perm missing element



An array A consisting of N different integers is given. the array contains integers in the range [1..(n + 1)], which means that exactly one element is missing.
Your goal is to find that missing element.

Write a function:

    func solution(a []int) int

that, given an array a, returns the value of the missing element.

for example, given array a such that:
  a[0] = 2
  a[1] = 3
  a[2] = 1
  a[3] = 5

the function should return 4, as it is the missing element.

write an efficient algorithm for the following assumptions:

        n is an integer within the range [0..100,000];
        the elements of a are all distinct;
        each element of array a is an integer within the range [1..(n + 1)].

copyright 2009–2020 by codility limited. all rights reserved. unauthorized copying, publication or disclosure prohibited. 