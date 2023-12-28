// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CheckBadgeIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15 8C15 8.98224 14.5279 9.85423 13.7983 10.4015C13.9273 11.3044 13.6446 12.2549 12.95 12.9494C12.2553 13.6441 11.3048 13.9269 10.4019 13.7978C9.85465 14.5277 8.98248 15 8 15C7.01773 15 6.14572 14.5279 5.59848 13.7983C4.69556 13.9273 3.74514 13.6445 3.05055 12.9499C2.35595 12.2553 2.07316 11.3048 2.2022 10.4019C1.47228 9.85464 1 8.98248 1 8C1 7.01773 1.47208 6.14573 2.20173 5.59848C2.07273 4.69557 2.35552 3.74513 3.0501 3.05055C3.74471 2.35595 4.69519 2.07316 5.59813 2.2022C6.14535 1.47228 7.01752 1 8 1C8.98224 1 9.85423 1.47205 10.4015 2.20167C11.3044 2.07265 12.2549 2.35544 12.9494 3.05003C13.6441 3.74465 13.9269 4.69516 13.7978 5.59813C14.5277 6.14535 15 7.01752 15 8ZM11.7086 5.15657C12.0364 5.40984 12.0967 5.88086 11.8435 6.20862L7.59347 11.7086C7.45747 11.8846 7.25037 11.9912 7.0281 11.9995C6.80583 12.0078 6.59133 11.9171 6.44254 11.7518L4.19254 9.25176C3.91544 8.94388 3.9404 8.46966 4.24828 8.19257C4.55617 7.91547 5.03038 7.94043 5.30748 8.24831L6.95616 10.0802L10.6565 5.29145C10.9098 4.96369 11.3808 4.9033 11.7086 5.15657Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}