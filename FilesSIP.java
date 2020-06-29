import java.util.*; 
import java.lang.*; 
import java.io.*; 

public class FilesSIP  {
  public static void main (String[] args) {
    if (args.length != 1) {
      System.out.println("usage: java FilesSip folder");
      System.exit(1);
    }
    else {
      File f = new File(args[0]);
      if (!f.isDirectory()){
        System.out.println(args[0] + " is not a valid folder name");
        System.exit(1);
      }
    }
  }
}
