/**
 * Created by KSM on 2017/5/8.
 */

public class Container {
    public int maxArea(int[] height) {
        int maxarea = 0,left = 0,right = height.length-1;
        while(left < right){
            maxarea = Math.max(maxarea,Math.min(height[left],height[right]) * (right - left));
            if(height[left] >= height[right])
                --right;
            else
                ++left;
        }
        return maxarea;
    }

}
