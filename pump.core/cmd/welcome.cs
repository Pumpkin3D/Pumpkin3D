using System;
public class Pumpkin3DHelp
{
     public static void Main()
     {
        ShowHelpMenu();
     }  

     private static void ShowHelpMenu()
     {
        Console.WriteLine("Welcome to pumpkin3D :D");
        Console.WriteLine("These are some commands to use to get started");

        Console.WriteLine("init <project_name> |Create a new Pumpkin3D project file.")
        Console.WriteLine("patch <patch_plugin_name> opt var <patch_version_number> | Download a plugin/theme/libarry from Pumpkin3D's Patch repos. You can type 'patch help' for more info!")
        Console.WriteLine("git | Use Pumpkin 3D's built in git commands.")
        Console.WriteLine("depbi | Packege your project into a binary (exe, appimage, app, deb etc.)")
        Console.WriteLine("terrain | Edit your project's 3D env settings with a UI. ")

        Console.WriteLine("If you need further support feel free to ask our amazing community!")

        Console.WriteLine("Discord | https://discord.gg/JYzYkx3fPG")
        Console.WriteLine("Reddit | https://reddit.com/r/pumpkin3d")
     }
}
