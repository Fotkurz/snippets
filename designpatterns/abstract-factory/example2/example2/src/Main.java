import components.Button;
import components.ComponentFactory;
import components.Modal;
import components.darktheme.DarkThemeFactory;
import components.lighttheme.LightThemeFactory;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        ComponentFactory factory = new DarkThemeFactory();

        Button button = factory.createButton();
        Modal modal = factory.createModal();

        button.renderButton();
        modal.renderModal();

        factory = new LightThemeFactory();

        button = factory.createButton();
        modal = factory.createModal();

        button.renderButton();
        modal.renderModal();
    }
}