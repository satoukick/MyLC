import java.util.ArrayList;
import java.util.Scanner;

class Solution {
    public int[] searchRange(int[] nums, int target) {
        int range[] ={nums.length,-1};
        searchRange(nums,target, 0 ,nums.length-1, range);
        if(range[0] > range[1]) range[0] = -1;
        return range;
    }

    public void searchRange(int[] nums, int target, int left,int right,int range[]){
        if(left > right) return;
        int mid = (left + right) / 2;
        if(nums[mid] == target) {
            if (mid < range[0]) {
                range[0] = mid;
                searchRange(nums, target, left, mid - 1, range);
            }
            if (mid > range[1]) {
                range[1] = mid;
                searchRange(nums, target, mid + 1, right, range);
            }
        }else if(nums[mid] > target)
            searchRange(nums,target,left,mid-1,range);
        else
            searchRange(nums,target,mid+1,right,range);

    }
}

public class searchRange {
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        String input = sc.nextLine();
        int target = sc.nextInt();
        String array[] = input.split(" +");
        int num[] = new int[array.length];
        for(int i = 0 ;i< array.length;i++) {
            num[i] = Integer.parseInt(array[i]);
        }
        Solution solu = new Solution();
        int res[] = solu.searchRange(num,target);
        for(int r:res)
            System.out.print(r + " ");

    }
}
