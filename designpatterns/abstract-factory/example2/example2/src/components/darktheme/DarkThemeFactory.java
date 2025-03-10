package components.darktheme;

import components.Button;
import components.ComponentFactory;
import components.Modal;
import components.lighttheme.LightModal;

public class DarkThemeFactory  extends ComponentFactory {
    @Override
    public Button createButton() {
        return new DarkButton();
    }

    @Override
    public Modal createModal() {
        return new DarkModal();
    }
}
