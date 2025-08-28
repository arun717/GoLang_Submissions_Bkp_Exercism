public class Lasagna {
    // TODO: define the 'expectedMinutesInOven()' method
    public int expectedMinutesInOven() {
        return 40;
    }

    // TODO: define the 'remainingMinutesInOven()' method
    public int remainingMinutesInOven(int actual_minutes) {
        return 40 - actual_minutes;
    }

    // TODO: define the 'preparationTimeInMinutes()' method
    public int preparationTimeInMinutes(int no_of_layers) {
        return 2 * no_of_layers;
    }

    // TODO: define the 'totalTimeInMinutes()' method
    public int totalTimeInMinutes(int no_of_layers, int no_of_min) {
        return preparationTimeInMinutes(no_of_layers) + no_of_min;
    }
}
