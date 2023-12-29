// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AdjustmentsHorizontalIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path d=\"M18.75 12.75L20.25 12.75C20.6642 12.75 21 12.4142 21 12C21 11.5858 20.6642 11.25 20.25 11.25L18.75 11.25C18.3358 11.25 18 11.5858 18 12C18 12.4142 18.3358 12.75 18.75 12.75Z\"></path> <path d=\"M12 6C12 5.58579 12.3358 5.25 12.75 5.25L20.25 5.25002C20.6642 5.25002 21 5.5858 21 6.00002C21 6.41423 20.6642 6.75002 20.25 6.75002L12.75 6.75C12.3358 6.75 12 6.41421 12 6Z\"></path> <path d=\"M12 18C12 17.5858 12.3358 17.25 12.75 17.25L20.25 17.25C20.6642 17.25 21 17.5858 21 18C21 18.4142 20.6642 18.75 20.25 18.75L12.75 18.75C12.3358 18.75 12 18.4142 12 18Z\"></path> <path d=\"M3.75001 6.75001L5.25001 6.75C5.66422 6.75 6 6.41421 6 5.99999C6 5.58578 5.66421 5.25 5.24999 5.25L3.74999 5.25001C3.33578 5.25002 3 5.58581 3 6.00002C3 6.41424 3.33579 6.75002 3.75001 6.75001Z\"></path> <path d=\"M5.25001 18.75L3.75001 18.75C3.33579 18.75 3 18.4142 3 18C3 17.5858 3.33578 17.25 3.74999 17.25L5.24999 17.25C5.66421 17.25 6 17.5858 6 18C6 18.4142 5.66422 18.75 5.25001 18.75Z\"></path> <path d=\"M3 12C3 11.5858 3.33579 11.25 3.75 11.25H11.25C11.6642 11.25 12 11.5858 12 12C12 12.4142 11.6642 12.75 11.25 12.75H3.75C3.33579 12.75 3 12.4142 3 12Z\"></path> <path d=\"M9 3.75C7.75736 3.75 6.75 4.75736 6.75 6C6.75 7.24264 7.75736 8.25 9 8.25C10.2426 8.25 11.25 7.24264 11.25 6C11.25 4.75736 10.2426 3.75 9 3.75Z\"></path> <path d=\"M12.75 12C12.75 10.7574 13.7574 9.75 15 9.75C16.2426 9.75 17.25 10.7574 17.25 12C17.25 13.2426 16.2426 14.25 15 14.25C13.7574 14.25 12.75 13.2426 12.75 12Z\"></path> <path d=\"M9 15.75C7.75736 15.75 6.75 16.7574 6.75 18C6.75 19.2426 7.75736 20.25 9 20.25C10.2426 20.25 11.25 19.2426 11.25 18C11.25 16.7574 10.2426 15.75 9 15.75Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
