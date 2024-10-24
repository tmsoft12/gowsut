import win32print
import win32ui
def Printc(o):

    san = o
    etiket_metin = f"Hormatly Musteri\n*********************\nsizin nobatynyz\n          {san}\n*********************"

    yazici_adı = win32print.GetDefaultPrinter()

    hprinter = win32print.OpenPrinter(yazici_adı)
    try:
        hdc = win32ui.CreateDC()
        hdc.CreatePrinterDC(yazici_adı)
        
        hdc.StartDoc("Yazdırma İşi")
        hdc.StartPage()
        
        font = win32ui.CreateFont({
            "name": "Courier",
            "height": 50,  
            "weight": 700 
        })
        hdc.SelectObject(font)
        
      
        x, y = 50, 50  
        line_height = 50  
        
        for line in etiket_metin.split('\n'):
            hdc.TextOut(x, y, line)
            y += line_height
        
        hdc.EndPage()
        hdc.EndDoc()
        
    finally:
        win32print.ClosePrinter(hprinter)
        hdc.DeleteDC() 
Printc("Nas Atma Yakup")