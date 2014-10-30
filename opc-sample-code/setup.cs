// Declare a Server first.
String OPCServerName = "{Here comes your OPC s name}";
ObjOPCServer = new OPCServer();
ObjOPCServer.Connect(OPCServerName, "";);
// Then a Object group because to add an item changed event.
ObjOPCGroups = ObjOPCServer.OPCGroups;
ObjOPCGroup = ObjOPCGroups.Add("OPCGroup1");
// The item changed event.
ObjOPCGroup.DataChange += new
DIOPCGroupEvent_DataChangeEventHandler(ObjOPCGroup_DataChange);
// Now the Iteem
ObjOPCGroup.OPCItems.AddItem("{tag name or address (like {plc name on server}!%mw0)}", 1);
// Update rate and other miscellaneous
ObjOPCGroup.UpdateRate = 10;
ObjOPCGroup.IsActive = true;
ObjOPCGroup.IsSubscribed = true;
