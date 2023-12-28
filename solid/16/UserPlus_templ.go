// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func UserPlusIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M8.5 4.5C8.5 5.88071 7.38071 7 6 7C4.61929 7 3.5 5.88071 3.5 4.5C3.5 3.11929 4.61929 2 6 2C7.38071 2 8.5 3.11929 8.5 4.5Z\" fill=\"black\"></path> <path d=\"M9.99951 13C10.5518 13 11.0099 12.5478 10.9008 12.0064C10.44 9.72096 8.42072 8 5.99951 8C3.57829 8 1.55903 9.72096 1.09823 12.0064C0.989075 12.5478 1.44722 13 1.99951 13H9.99951Z\" fill=\"black\"></path> <path d=\"M12.5 3.5C12.9142 3.5 13.25 3.83579 13.25 4.25V5.25H14.25C14.6642 5.25 15 5.58579 15 6C15 6.41421 14.6642 6.75 14.25 6.75H13.25V7.75C13.25 8.16421 12.9142 8.5 12.5 8.5C12.0858 8.5 11.75 8.16421 11.75 7.75V6.75H10.75C10.3358 6.75 10 6.41421 10 6C10 5.58579 10.3358 5.25 10.75 5.25H11.75V4.25C11.75 3.83579 12.0858 3.5 12.5 3.5Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}