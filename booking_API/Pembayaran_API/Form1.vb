Imports System.Net
Imports System.IO
Imports Newtonsoft.Json


Public Class Form1

    Sub get_booking()
        Dim url As String = "http://localhost:1323/tbl_booking"
        Dim request As HttpWebRequest = CType(WebRequest.Create(url), HttpWebRequest)
        request.Method = "GET"

        Dim response As HttpWebResponse = CType(request.GetResponse(), HttpWebResponse)
        Dim reader As New StreamReader(response.GetResponseStream())
        Dim responseText As String = reader.ReadToEnd()

        ' Konversi respons JSON ke dalam DataTable
        Dim table As DataTable = JsonConvert.DeserializeObject(Of DataTable)(responseText)

        ' Tampilkan data pada DataGridView
        DataGridView1.DataSource = table
    End Sub

    Private Sub DataGridView1_CellContentClick(ByVal sender As System.Object, ByVal e As System.Windows.Forms.DataGridViewCellEventArgs) Handles DataGridView1.CellContentClick

    End Sub

    
    Private Sub Form1_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
        Dim url As String = "http://localhost:1323/tbl_booking"
        Dim request As HttpWebRequest = CType(WebRequest.Create(url), HttpWebRequest)
        request.Method = "GET"

        Dim response As HttpWebResponse = CType(request.GetResponse(), HttpWebResponse)
        Dim reader As New StreamReader(response.GetResponseStream())
        Dim responseText As String = reader.ReadToEnd()

        ' Konversi respons JSON ke dalam DataTable
        Dim table As DataTable = JsonConvert.DeserializeObject(Of DataTable)(responseText)

        ' Tampilkan data pada DataGridView
        DataGridView1.DataSource = table
    End Sub
End Class


