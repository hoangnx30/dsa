import java.util.Arrays;

public class MaximumMatchingOfPlayerWithTrainer {

    static int matchPlayersAndTrainers(int[] players, int[] trainers) {
        int result = 0;

        Arrays.sort(players);
        Arrays.sort(trainers);

        int i = 0;
        int j = 0;

        while (i < players.length && j < trainers.length) {
            if (players[i] <= trainers[j]) {
                i++;
                j++;
                result++;
            } else {
                j++;
            }

        }

        return result;
    }

    public static void main(String[] args) {
        int[] players = new int[]{4, 7, 9};
        int[] trainers = new int[]{8, 2, 5, 8};

        int result = matchPlayersAndTrainers(players, trainers);

        System.out.printf("Players: %s, Trainer: %s => Result: %d", Arrays.toString(players), Arrays.toString(trainers), result);
    }


}
