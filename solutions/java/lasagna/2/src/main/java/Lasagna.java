public class Lasagna {    
    public int expectedMinutesInOven() {
        return 40;
    }

    public int remainingMinutesInOven(int actual_minutes) {
        return expectedMinutesInOven() - actual_minutes;
    }

    public int preparationTimeInMinutes(int no_of_layers) {
        return 2 * no_of_layers;
    }

    public int totalTimeInMinutes(int no_of_layers, int no_of_min) {
        return preparationTimeInMinutes(no_of_layers) + no_of_min;
    }
}
