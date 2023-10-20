
// Online Java Compiler
// Use this editor to write, compile and run your Java code online
import java.util.*;

class a {

    static Map<String, Integer> hmap = new HashMap<>();

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int max_notes = sc.nextInt();
        int a_100 = sc.nextInt();
        int b_200 = sc.nextInt();
        int c_500 = sc.nextInt();
        int d_1000 = sc.nextInt();
        int withDraw = sc.nextInt();
        int amt = a_100 * 100 + b_200 * 200 + c_500 * 500 + d_1000 * 1000;
        if (amt > withDraw) {
            System.out.println("No Way Home");
            return;
        }

        int temp = recur(withDraw, 0, max_notes, a_100, b_200, c_500, d_1000);
        if (temp == Integer.MIN_VALUE)
            System.out.println("No Way Home");
        else
            System.out.println(temp);
    }

    static int recur(int amt_left, int notes_used, int max_notes, int a, int b, int c, int d) {

        if (amt_left < 0 || notes_used >= max_notes) {
            return Integer.MIN_VALUE;
        }

        if (hmap.containsKey(a + " " + b + " " + c + " " + d))
            return hmap.get(a + " " + b + " " + c + " " + d);

        if (amt_left == 0) {
            // System.out.println(amt_left + " " + notes_used + " " + a + " " + b + " " + c + " " + d);
            hmap.put(a + " " + b + " " + c + " " + d, notes_used);
            return notes_used;
        }
        

        int notes_max = Integer.MIN_VALUE;
        if (a > 0)
            notes_max = Math.max(max_notes, recur(amt_left - 100, notes_used + 1, max_notes, a - 1, b, c, d));
        if (b > 0)
            notes_max = Math.max(max_notes, recur(amt_left - 200, notes_used + 1, max_notes, a, b - 1, c, d));
        if (c > 0)
            notes_max = Math.max(max_notes, recur(amt_left - 500, notes_used + 1, max_notes, a, b, c - 1, d));
        if (d > 0)
            notes_max = Math.max(max_notes, recur(amt_left - 1000, notes_used + 1, max_notes, a, b, c, d - 1));
        return notes_max;
        

    }
}