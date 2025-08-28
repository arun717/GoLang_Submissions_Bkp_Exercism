public class Lasagna {
    
    private static final int EXPECTED_MINUTES_IN_OVEN = 40;
    private static final int PREPARATION_TIME_PER_LAYER = 2;
    
    public int expectedMinutesInOven() {
        return EXPECTED_MINUTES_IN_OVEN;
    }

    public int remainingMinutesInOven(int actual_minutes) {
        return expectedMinutesInOven() - actual_minutes;
    }

    public int preparationTimeInMinutes(int no_of_layers) {
        return PREPARATION_TIME_PER_LAYER * no_of_layers;
    }

    public int totalTimeInMinutes(int no_of_layers, int no_of_min) {
        return preparationTimeInMinutes(no_of_layers) + no_of_min;
    }
    
}