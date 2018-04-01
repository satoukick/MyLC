package fibonacci;

public class Fibonacci {
    public int Fibonacci(int n) {
         if(n == 1  || n == 0)
             return n;
         return Fibonacci(n-1) + Fibonacci(n - 2);
    }

    public static void main (String[] args) {
        Fibonacci s = new Fibonacci();
        int result = s.Fibonacci(1);
        System.out.println(result);
    }
}

