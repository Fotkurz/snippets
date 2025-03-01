package factory;

import button.Button;
import button.WindowsButton;

public class WindowsDialog extends Dialog {
    @Override
    public void render() {

    }

    @Override
    public Button createButton() {
        return new WindowsButton();
    }
}
