public class Document {

    private final String header;
    private final String body;
    private final String footer;

    Document(String header, String body, String footer ){
        this.header = "Title " + header;
        this.body = body;
        this.footer = footer;
    }

    public void print() {
        String separator = "";

        for (int i = 0; i < this.header.length(); i++) {
            separator = separator.concat("-");
        }

        System.out.println(this.header);
        System.out.println(separator+"\n");
        System.out.println(this.body+"\n");
        System.out.println(this.footer);
    }
}
