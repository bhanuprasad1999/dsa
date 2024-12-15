"""
Given an array of numbers sorted in ascending order and a target sum, find a pair in the array whose sum is equal to the given target.

Write a function to return the indices of the two numbers (i.e. the pair) such that they add up to the given target. If no such pair exists return [-1, -1].

Example 1:

Input: [1, 2, 3, 4, 6], target=6
Output: [1, 3]
Explanation: The numbers at index 1 and 3 add up to 6: 2+4=6
"""
import os


def find_pair_nums_for_target_sum(nums, target):
    left, right = (0, len(nums)-1)
    while left < right :
        curr_sum = nums[left] + nums[right]
        if curr_sum== target:
            return (nums[left], nums[right])
        
        elif curr_sum < target:
            left += 1
        else:
            right -= 1
    return (-1, -1)

if __name__ == '__main__':
    nums = [1, 2, 3, 4, 6]
    target = 6
    target_nums = find_pair_nums_for_target_sum(nums, target)
    print(target_nums) #the final two numbers that target the sum.
    