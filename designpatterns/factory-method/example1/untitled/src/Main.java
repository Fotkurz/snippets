import button.Button;
import factory.Dialog;
import factory.LinuxDialog;
import factory.WindowsDialog;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        Dialog factory = new WindowsDialog();

        Button button = factory.createButton();

        button.render();
        button.onClick();

        factory = new LinuxDialog();

        button = factory.createButton();

        button.render();
        button.onClick();
    }
}