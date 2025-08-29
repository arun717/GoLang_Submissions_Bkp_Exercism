public class CarsAssemble {
    public int CARS_PER_HOUR = 221;
    public double productionRatePerHour(int speed) {
        int cars = speed * CARS_PER_HOUR;
        if (speed <= 4) {
            return (double)cars;
        } else if (speed <= 8) {
            return cars * 0.9;
        } else if (speed == 9) {
            return cars * 0.8;
        } else {
            return cars * 0.77;
        }
    }

    public int workingItemsPerMinute(int speed) {
        return (int)productionRatePerHour(speed)/60;
    }
}
