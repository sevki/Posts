using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using OPCAutomation;
namespace OPC
{
    public partial class Form1 : Form
    {
        OPCServer ObjOPCServer;
        OPCGroups ObjOPCGroups;
        OPCGroup ObjOPCGroup;
        string OPCServerName;
        public Form1()
        {
            try
            {
                InitializeComponent();
                OPCServerName = "{Here comes your OPC server’s name}";
                ObjOPCServer = new OPCServer();
                ObjOPCServer.Connect(OPCServerName, "";);
                ObjOPCGroups = ObjOPCServer.OPCGroups;
                ObjOPCGroup = ObjOPCGroups.Add("OPCGroup1");
                ObjOPCGroup.DataChange += new DIOPCGroupEvent_DataChangeEventHandler(ObjOPCGroup_DataChange);
                ObjOPCGroup.OPCItems.AddItem("{tag name or address (like {plc name on server}!%mw0)}", 1);
                ObjOPCGroup.UpdateRate = 10;
                ObjOPCGroup.IsActive = true;
                ObjOPCGroup.IsSubscribed = true;
            }
            catch (Exception e){
                MessageBox.Show(e.ToString());
            }
        }

        private void button1_Click(object sender, EventArgs e)
        {
        }

        private void ObjOPCGroup_DataChange(int TransactionID, int NumItems, ref Array ClientHandles, ref Array ItemValues, ref Array Qualities, ref Array TimeStamps)
        {
            for (int i = 1; i <= NumItems; i++)
            {
                if ((Convert.ToInt32(ClientHandles.GetValue(i)) == 1))
                {
                    textBox1.Text =  ItemValues.GetValue(i).ToString();
                }
            }
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            OPCServerClass GlobalOPCServer = new OPCServerClass();
            Array ServerList = (Array)GlobalOPCServer.GetOPCServers("");
            for (int i = 1; i <= ServerList.Length; i++)
            {
                comboBox1.Items.Add(ServerList.GetValue(i).ToString());
            }
        }
    }
}
