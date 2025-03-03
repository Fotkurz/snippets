import burger.factory.BeefBurgerRestaurant;
import burger.factory.ChickenBurgerRestaurant;
import burger.factory.Restaurant;
import burger.factory.VeggieBurgerRestaurant;
import burger.product.Burger;

public class Main {
    public static void main(String[] args) {
        Restaurant restaurant = new BeefBurgerRestaurant();
        restaurant.order();

        restaurant = new VeggieBurgerRestaurant();
        restaurant.order();

        restaurant = new ChickenBurgerRestaurant();
        restaurant.order();
    }
}