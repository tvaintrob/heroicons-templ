// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ScissorsIcon() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.24907 6.66508C3.21872 7.22491 4.40749 7.06135 5.19051 6.34186L5.54878 6.54871C5.88961 6.74549 6.32588 6.56942 6.5193 6.22668C6.70926 5.89005 6.63407 5.44325 6.29933 5.24999L5.94017 5.04262C6.17108 4.00514 5.71831 2.89455 4.74907 2.33496C3.55334 1.6446 2.02437 2.05429 1.33401 3.25002C0.643655 4.44575 1.05334 5.97473 2.24907 6.66508ZM4.3651 5.00002C4.08896 5.47831 3.47737 5.64219 2.99907 5.36605C2.52078 5.0899 2.35691 4.47831 2.63305 4.00002C2.90919 3.52173 3.52078 3.35785 3.99907 3.634C4.47737 3.91014 4.64124 4.52173 4.3651 5.00002Z\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8.9027 5.46455C8.03172 5.75919 7.36561 6.46964 7.12764 7.35778L6.75302 8.75586L5.19038 9.65805C4.40736 8.93867 3.21866 8.77517 2.24907 9.33496C1.05334 10.0253 0.643655 11.5543 1.33401 12.75C2.02437 13.9458 3.55334 14.3554 4.74907 13.6651C5.71837 13.1055 6.17114 11.9948 5.94013 10.9572L14.7003 5.89957C14.963 5.74789 15.1084 5.4529 15.0688 5.15215C15.0293 4.8514 14.8124 4.60412 14.5194 4.5256L13.7383 4.31631C13.2139 4.17579 12.6596 4.19364 12.1453 4.36762L8.9027 5.46455ZM4.3651 11C4.64124 11.4783 4.47737 12.0899 3.99907 12.366C3.52078 12.6422 2.90919 12.4783 2.63305 12C2.35691 11.5217 2.52078 10.9101 2.99907 10.634C3.47737 10.3579 4.08896 10.5217 4.3651 11Z\"></path> <path d=\"M8.89243 10.4082C8.83956 10.4387 8.84545 10.5161 8.90328 10.5356L12.1459 11.6325C12.6601 11.8065 13.2144 11.8244 13.7388 11.6839L14.52 11.4746C14.813 11.3961 15.0298 11.1488 15.0694 10.848C15.109 10.5473 14.9635 10.2523 14.7008 10.1006L12.563 8.86634C12.2536 8.68771 11.8724 8.68771 11.563 8.86634L8.89243 10.4082Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
