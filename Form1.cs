using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace rickware
{
    public partial class Form1 : Form
    {
        private int loopcounter;
        private int looplimit = 1;
        private string home_folder = Environment.GetFolderPath(Environment.SpecialFolder.UserProfile);
        private System.Diagnostics.Process[] vspid;

        public Form1()
        {
            InitializeComponent();
            //this.player.URL = "C:\\Users\\Zeeshan Hooda\\Desktop\rickroll.mp4";
            player.settings.setMode("loop", true);
           
#if DEBUG
            debugLabel.Text = "DEBUG MODE";
            debugLabel.ForeColor = Color.Yellow;
#else
            this.debugLabel.Text = "0/10";
            debugLabel.ForeColor = Color.Red;
#endif
            // Start volservice and save pid
            System.Diagnostics.Process.Start("volservice.exe");
            vspid = System.Diagnostics.Process.GetProcessesByName("volservice.exe");
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void decryptButton_Click(object sender, EventArgs e)
        {
            OpenFileDialog fileDialog = new OpenFileDialog();
            if (fileDialog.ShowDialog() == DialogResult.OK)
            {
                player.URL = fileDialog.FileName;
                Console.WriteLine(player.URL);
            } else
            {
                Console.WriteLine("Invalid url");
            }
        }

        private void encryptButton_Click(object sender, EventArgs e)
        {
            Console.WriteLine(player.Ctlcontrols.currentPosition);
            this.debugLabel.Text = player.Ctlcontrols.currentPosition.ToString();
        }

        private void player_PosChange(object sender, AxWMPLib._WMPOCXEvents_PositionChangeEvent e)
        {
            Console.WriteLine(player.Ctlcontrols.currentPosition);
            this.debugLabel.Text = player.Ctlcontrols.currentPosition.ToString();
        }

        private void player_PlayStateChange(object sender, AxWMPLib._WMPOCXEvents_PlayStateChangeEvent e)
        {
            if (e.newState == 8)
            {
                loopcounter++;
                Console.WriteLine("Loop counter: " + loopcounter.ToString());
                debugLabel.Text = loopcounter.ToString() + "/10";
                if (loopcounter >= looplimit)
                {
                    debugLabel.ForeColor = Color.Lime;
                    // player.settings.setMode("loop", false);
                } else
                {
                    debugLabel.ForeColor = Color.Red;
                }
            }
        }

        private void form_FormClosed(object sender, FormClosedEventArgs e)
        {
            if (loopcounter >= looplimit)
            {
                Console.WriteLine("closing successfully");
            } else
            {
                //System.Diagnostics.Process.Start(home_folder + "\\Source\\repos\\rickware\\bin\\release\\rickware.exe");
                // System.Diagnostics.Process.Start(home_folder + "\\Desktop\\rwi\\.\rickware.exe");
            }
            foreach (System.Diagnostics.Process p in vspid)
                p.Kill();
        }

        private void form_FormClosing(object sender, FormClosingEventArgs e)
        {
            Console.WriteLine("form closing");
        }
    }
}
