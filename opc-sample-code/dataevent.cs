private void ObjOPCGroup_DataChange(int TransactionID, int NumItems, ref Array ClientHandles, ref Array ItemValues, ref Array Qualities, ref Array TimeStamps)
{
    for (int i = 1; i &lt;= NumItems; i++){
        if ((Convert.ToInt32(ClientHandles.GetValue(i)) == 1)) {
            textBox1.Text = ItemValues.GetValue(i).ToString();
        }
    }
}
