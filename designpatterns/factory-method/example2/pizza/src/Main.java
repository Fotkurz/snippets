import factory.PepperoniFactory;
import factory.PizzaFactory;
import pizza.Pizza;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        PizzaStore store = new PizzaStore();

        Pizza pizza = store.prepare("pepperoni");
        pizza.prepare();
    }
}