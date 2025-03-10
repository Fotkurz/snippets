package components.lighttheme;

import components.Button;
import components.ComponentFactory;
import components.Modal;

public class LightThemeFactory extends ComponentFactory {
    @Override
    public Button createButton() {
        return new LightButton();
    }

    @Override
    public Modal createModal() {
        return new LightModal();
    }
}
