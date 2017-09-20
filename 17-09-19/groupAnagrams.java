
import java.util.*;

public class groupAna {
    public List<List<String>> groupAnagrams(String[] strs){
        if(strs == null || strs.length == 0) return new ArrayList<List<String>>();
        Map<String,List<String>> map = new HashMap<String,List<String>>();
        for (String s:strs){
            char[] ca = s.toCharArray();
            Arrays.sort(ca);
            String KeyStr = String.valueOf(ca);
            if(!map.containsKey(KeyStr)) map.put(KeyStr,new ArrayList<String>());
            map.get(KeyStr).add(s);
         }
        return new ArrayList<List<String>>(map.values());
    }
}
