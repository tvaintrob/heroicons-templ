// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BuildingStorefrontIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M4.5 7C5.18136 7 5.79906 6.72742 6.25 6.28536C6.70094 6.72742 7.31864 7 8 7C8.68136 7 9.29906 6.72742 9.75 6.28536C10.2009 6.72742 10.8186 7 11.5 7C12.8807 7 14 5.88071 14 4.5C14 3.11929 12.8807 2 11.5 2H4.5C3.11929 2 2 3.11929 2 4.5C2 5.88071 3.11929 7 4.5 7Z\"></path> <path d=\"M6.25 8.09741C5.72129 8.35499 5.12707 8.49995 4.5 8.49995C3.96955 8.49995 3.46322 8.39669 3 8.20919V12.4999H2.75C2.33579 12.4999 2 12.8357 2 13.2499C2 13.6642 2.33579 13.9999 2.75 13.9999H3.25C3.29708 13.9999 3.34314 13.9956 3.38782 13.9873C3.42388 13.9956 3.46143 13.9999 3.5 13.9999H6C6.27614 13.9999 6.5 13.7761 6.5 13.4999V10.4999C6.5 10.2238 6.72386 9.99995 7 9.99995H9C9.27614 9.99995 9.5 10.2238 9.5 10.4999V13.4999C9.5 13.7761 9.72386 13.9999 10 13.9999H12.5C12.5386 13.9999 12.5761 13.9956 12.6122 13.9873C12.6569 13.9956 12.7029 13.9999 12.75 13.9999H13.25C13.6642 13.9999 14 13.6642 14 13.2499C14 12.8357 13.6642 12.4999 13.25 12.4999H13V8.20919C12.5368 8.39669 12.0304 8.49995 11.5 8.49995C10.8729 8.49995 10.2787 8.35499 9.75 8.09741C9.22129 8.35499 8.62708 8.49995 8 8.49995C7.37292 8.49995 6.77871 8.35499 6.25 8.09741Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
