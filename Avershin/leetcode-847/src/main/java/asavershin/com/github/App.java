package asavershin.com.github;

import java.util.*;


public class App {

    public int shortestPathLength(int[][] graph) {
        final int n = graph.length;
        final int cols = 1 << n;
        var combinationsDist = new Integer[n][cols];

        Queue<int[]> queue = new LinkedList<>();

        for (int i = 0; i < n; ++i) {
            combinationsDist[i][1 << i] = 0;
            queue.add(new int[]{i, 1 << i});
        }

        while (!queue.isEmpty()) {
            var size = queue.size();
            for (int i = 0; i < size; ++i) {
                var current = queue.remove();

                if (current[1] == cols - 1) {
                    return combinationsDist[current[0]][current[1]];
                }

                for (var child : graph[current[0]]) {
                    int childMask = current[1] | 1 << child;
                    if (combinationsDist[child][childMask] != null) {
                        continue;
                    }
                    combinationsDist[child][childMask] = combinationsDist[current[0]][current[1]] + 1;
                    queue.add(new int[]{child, childMask});
                }
            }
        }
        return 12;
    }

    public static void main(String[] args) {

    }
}
