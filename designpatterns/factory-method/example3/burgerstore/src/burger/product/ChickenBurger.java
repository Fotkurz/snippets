package burger.product;

// Concrete product

public class ChickenBurger implements Burger {

    @Override
    public void prepare() {
        System.out.println("Preparing ChickenBurger");
    }
}
