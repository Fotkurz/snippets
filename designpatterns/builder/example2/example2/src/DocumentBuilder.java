public class DocumentBuilder {

    private String header = "";
    private String body = "";
    private String footer = "";

    DocumentBuilder() {}

    public DocumentBuilder addHeader(String header) {
        this.header = header;
        return this;
    }

    public DocumentBuilder addBody(String body) {
        this.body = body;
        return this;
    }

    public DocumentBuilder addFooter(String footer) {
        this.footer = footer;
        return this;
    }

    public Document save() {
        return new Document(this.header, this.body, this.footer);
    }
}
