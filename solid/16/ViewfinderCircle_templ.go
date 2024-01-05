// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ViewfinderCircleIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 16 16\" class=\"w-4 h-4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M3.75 2C2.7835 2 2 2.7835 2 3.75V5.25C2 5.66421 2.33579 6 2.75 6C3.16421 6 3.5 5.66421 3.5 5.25V3.75C3.5 3.61193 3.61193 3.5 3.75 3.5H5.25C5.66421 3.5 6 3.16421 6 2.75C6 2.33579 5.66421 2 5.25 2H3.75Z\"></path><path d=\"M10.75 2C10.3358 2 10 2.33579 10 2.75C10 3.16421 10.3358 3.5 10.75 3.5H12.25C12.3881 3.5 12.5 3.61193 12.5 3.75V5.25C12.5 5.66421 12.8358 6 13.25 6C13.6642 6 14 5.66421 14 5.25V3.75C14 2.7835 13.2165 2 12.25 2H10.75Z\"></path><path d=\"M3.5 10.75C3.5 10.3358 3.16421 10 2.75 10C2.33579 10 2 10.3358 2 10.75V12.25C2 13.2165 2.7835 14 3.75 14H5.25C5.66421 14 6 13.6642 6 13.25C6 12.8358 5.66421 12.5 5.25 12.5H3.75C3.61193 12.5 3.5 12.3881 3.5 12.25V10.75Z\"></path><path d=\"M14 10.75C14 10.3358 13.6642 10 13.25 10C12.8358 10 12.5 10.3358 12.5 10.75V12.25C12.5 12.3881 12.3881 12.5 12.25 12.5H10.75C10.3358 12.5 10 12.8358 10 13.25C10 13.6642 10.3358 14 10.75 14H12.25C13.2165 14 14 13.2165 14 12.25V10.75Z\"></path><path d=\"M8 10C9.10457 10 10 9.10457 10 8C10 6.89543 9.10457 6 8 6C6.89543 6 6 6.89543 6 8C6 9.10457 6.89543 10 8 10Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
