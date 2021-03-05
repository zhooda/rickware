using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;
using System.Reflection;


namespace rickware
{
    class Program
    {

        /// <summary>
        /// The main entry point for the application.
        /// </summary>
        [STAThread]
        static void Main()
        {
            Setup();

            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            Application.Run(new Form1());
        }

        static void Setup()
        {
            Assembly assembly = Assembly.GetExecutingAssembly();
            Stream stream = assembly.GetManifestResourceStream("rickware.Resources.AxInterop.WMPLib.dll");
            using (Stream file = File.Create("AxInterop.WMPLib.dll"))
            {
                stream.CopyTo(file);
            }
            stream = assembly.GetManifestResourceStream("rickware.Resources.Interop.WMPLib.dll");
            using (Stream file = File.Create("Interop.WMPLib.dll"))
            {
                stream.CopyTo(file);
            }
            stream = assembly.GetManifestResourceStream("rickware.Resources.volservice.exe");
            using (Stream file = File.Create("volservice.exe"))
            {
                stream.CopyTo(file);
            }


        }

        void Deconstruct()
        {
            // TODO: delete all dlls
        }
    }
}
